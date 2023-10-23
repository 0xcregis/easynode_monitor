package ether

import "testing"

func TestLatestBlock(t *testing.T) {
	n, d, err := LatestBlock("https://ethereum.publicnode.com", "")
	if err != nil {
		t.Error(err)
	} else {
		t.Logf("number:%v,duration:%v", n, d)
	}
}
