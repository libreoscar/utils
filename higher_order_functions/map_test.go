package hof

import (
	"fmt"
	"github.com/libreoscar/evm/utils"
	"testing"
)

var tMap = map[string]int{
	"a":        100,
	"b":        99,
	"rex":      97,
	"btcchina": 90,
}

func TestKeys(t *testing.T) {
	keys := utils.Keys(tMap).([]string)
	fmt.Println(keys)
}

func TestValues(t *testing.T) {
	vals := utils.Values(tMap).([]int)
	fmt.Println(vals)
}
