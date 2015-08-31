package reverse

import (
	"github.com/facebookgo/ensure"
	"testing"
)

func TestInt(t *testing.T) {
	{
		array := []int{1, 3, 4, 8, 6}
		Ints(array)
		ensure.DeepEqual(t, array, []int{6, 8, 4, 3, 1})
	}
	{
		array := []int{1, 3, 4, 8, 6, 8}
		Ints(array)
		ensure.DeepEqual(t, array, []int{8, 6, 8, 4, 3, 1})
	}
	{
		array := []int{1}
		Ints(array)
		ensure.DeepEqual(t, array, []int{1})
	}
	{
		array := []int{1, 2}
		Ints(array)
		ensure.DeepEqual(t, array, []int{2, 1})
	}
	{
		array := []int{}
		Ints(array)
		ensure.DeepEqual(t, array, []int{})
	}
	{
		var array []int
		Ints(array)
		var Nil []int
		ensure.DeepEqual(t, array, Nil)
	}
}
