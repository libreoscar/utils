package hof

import (
	"github.com/facebookgo/ensure"
	"testing"
)

var tMap = map[string]int{
	"a":        100,
	"b":        99,
	"rex":      97,
	"btcchina": 90,
}

func TestKeys(t *testing.T) {
	keys := Keys(tMap).([]string)
	ensure.SameElements(t, keys, []string{"btcchina", "rex", "a", "b"})
}

func TestValues(t *testing.T) {
	vals := Values(tMap).([]int)
	ensure.SameElements(t, vals, []int{97, 99, 90, 100})
}
