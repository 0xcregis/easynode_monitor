package chain

import (
	"fmt"
	"strings"

	"github.com/0xcregis/easynode_monitor/common"
	"github.com/0xcregis/easynode_monitor/service/chain/ether"
	"github.com/tidwall/gjson"
)

func GetLatestBlock(chinCode string, fullNode *common.FullNode) (number string, during int64, err error) {
	//http host+port
	/**
	  nodeUri like this:
	  [
	   "https://ethereum.publicnode.com"
	   ]
	*/

	var host string
	if strings.HasPrefix(fullNode.NodeUri, "{") || strings.HasPrefix(fullNode.NodeUri, "[") {
		if gjson.Parse(fullNode.NodeUri).IsArray() {
			host = gjson.Parse(fullNode.NodeUri).Array()[0].String()
		} else {
			return "", 0, fmt.Errorf("node uri is error")
		}
	} else {
		host = fullNode.NodeUri
	}

	token := fullNode.NodeJwt
	if chinCode == "200" {
		return ether.LatestBlock(host, token)
	}
	return "", 0, nil
}
