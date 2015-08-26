package hof

import (
	"github.com/facebookgo/ensure"
	"testing"
)

func TestMap(t *testing.T) {
	xs := []int{1, 2, 3, 4, -1, 0}
	ys := Map(xs, func(n int) int { return n * 2 }).([]int)
	ensure.DeepEqual(t, ys, []int{2, 4, 6, 8, -2, 0})
}

func TestFilter(t *testing.T) {
	xs := []int{2, 4, 5, 1, 2, 6, 4}
	ys := Filter(xs, func(n int) bool {
		return n > 3
	}).([]int)
	ensure.DeepEqual(t, ys, []int{4, 5, 6, 4})
}

func TestFoldl(t *testing.T) {
	{
		xs := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		sum := Foldl(xs, 0, func(sum, cur int) int { return sum + cur }).(int)
		ensure.DeepEqual(t, sum, 55)
	}
	{
		xs := []int{1, 2, 3}
		rst := Foldl(xs, 10, func(rst, cur int) int { return rst - cur }).(int)
		ensure.DeepEqual(t, rst, 4)
	}
}

func TestAll(t *testing.T) {
	{
		xs := []int{2, 4, 5, 6}
		isEven := func(n int) bool {
			return n%2 == 0
		}
		ensure.False(t, All(xs, isEven))
	}
	{
		xs := []byte("abcdef")
		isAlpha := func(b byte) bool {
			return b <= 'z' && b >= 'a'
		}
		ensure.True(t, All(xs, isAlpha))
	}
}

func TestExists(t *testing.T) {
	{
		xs := []int{2, 4, 5, 6}
		isEven := func(n int) bool {
			return n%2 == 0
		}
		ensure.True(t, Exists(xs, isEven))
	}
	{
		xs := []byte("abcdef")
		nonAlpha := func(b byte) bool {
			return !(b <= 'z' && b >= 'a')
		}
		ensure.False(t, Exists(xs, nonAlpha))
	}
}

func TestIn(t *testing.T) {
	{
		xs := []int{2, 4, 5, 6}
		ensure.True(t, In(4, xs))
		ensure.False(t, In(3, xs))
	}
	{
		a1 := []int{1, 2, 3}
		a2 := []int{1, 2}
		a3 := []int{1}
		a4 := []int{}
		xs := [][]int{a1, a2, a3, a4}

		var Nil []int = nil

		ensure.True(t, In([]int{1, 2}, xs))
		ensure.False(t, In([]int{2, 1}, xs))
		ensure.True(t, In(make([]int, 0), xs))
		ensure.False(t, In(Nil, xs))
	}
	{
		a1 := map[int]int{0: 0, 1: 1, 2: 2, 3: 3, 4: 4, 5: 5, 6: 6, 7: 7, 8: 8, 9: 9}
		a2 := map[int]int{0: 0, 1: 1, 2: 2, 3: 3, 4: 4, 5: 5, 6: 6, 7: 7, 8: 8}
		a3 := map[int]int{0: 0, 1: 1, 2: 2, 3: 3, 4: 4, 5: 5, 6: 6, 7: 7}

		a4 := map[int]int{}
		var a5 map[int]int = nil

		a6 := map[int]int{9: 9, 8: 8, 7: 7, 6: 6, 5: 5, 4: 4, 3: 3, 2: 2, 1: 1, 0: 0}
		a7 := map[int]int{1: 1, 2: 2, 3: 3, 4: 4, 5: 5, 6: 6, 7: 7, 8: 8, 9: 9}

		as1 := []map[int]int{a1, a2, a3, a4}
		as2 := []map[int]int{a1, a2, a3, a5}

		ensure.True(t, In(a6, as1))
		ensure.False(t, In(a7, as1))

		var Nil map[int]int = nil
		ensure.True(t, In(make(map[int]int), as1))
		ensure.False(t, In(Nil, as1))

		ensure.False(t, In(make(map[int]int), as2))
		ensure.True(t, In(Nil, as2))
	}
}
