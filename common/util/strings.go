package util

import (
	"errors"
	"math/big"
	"strconv"
	"strings"
)

func HexToInt(hex string) (string, error) {
	if len(hex) < 1 {
		return hex, errors.New("params is null when HexToInt is called")
	}
	if !strings.HasPrefix(hex, "0x") {
		hex = "0x" + hex
	}

	i, b := new(big.Int).SetString(hex, 0)
	if b {
		return i.String(), nil
	} else {
		return hex, errors.New("parse error when HexToInt is called")
	}

	//i, err := strconv.ParseInt(hex, 0, 64)
	//if err != nil {
	//	return hex, err
	//}
	//return fmt.Sprintf("%v", i), nil
}

func HexToInt2(hex string) (int64, error) {
	if len(hex) < 1 {
		return 0, errors.New("params is null")
	}
	if !strings.HasPrefix(hex, "0x") {
		return 0, errors.New("input string must be hex string")
	}
	i, err := strconv.ParseInt(hex, 0, 64)
	if err != nil {
		return 0, err
	}
	return i, nil
}
