package ether

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/0xcregis/easynode_monitor/common/util"
	"github.com/tidwall/gjson"
)

func LatestBlock(host, token string) (number string, duration int64, err error) {
	req := `{
				 "id": 1,
				 "jsonrpc": "2.0",
				 "method": "eth_blockNumber"
			}
			`
	start := time.Now()
	resp, err := SendReq(req, host, token)
	if err != nil {
		return "", 0, err
	}

	duration = time.Now().Sub(start).Microseconds()

	//{"jsonrpc":"2.0","id":1,"result":"0x1189fb8"}
	r := gjson.Parse(resp).Get("result").String()

	n, err := util.HexToInt(r)
	if err != nil {
		return "", 0, err
	}

	return n, duration, nil
}

func SendReq(reqBody string, host, token string) (resp string, err error) {
	reqBody = strings.Replace(reqBody, "\t", "", -1)
	reqBody = strings.Replace(reqBody, "\n", "", -1)

	if len(token) > 1 {
		host = fmt.Sprintf("%v/%v", host, token)
	}
	payload := strings.NewReader(reqBody)

	req, err := http.NewRequest("POST", host, payload)
	if err != nil {
		return "", err
	}

	req.Header.Add("accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("cache-control", "no-cache")
	//req.Header.Add("Postman-Token", "181e4572-a9db-453a-b7d4-17974f785de0")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)

	if err != nil {
		return "", err
	}

	if gjson.ParseBytes(body).Get("error").Exists() {
		return "", errors.New(string(body))
	}

	return string(body), nil
}
