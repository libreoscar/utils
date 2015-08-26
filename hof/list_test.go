package hof

import (
	"fmt"
	"reflect"
	"testing"
)

var _ = fmt.Println

func TestMap(t *testing.T) {
	func() {
		xs := []int{1, 2, 3, 4}
		ys := Map(xs, func(n int) int { return n * 2 }).([]int)
		if !reflect.DeepEqual(ys, []int{2, 4, 6, 8}) {
			t.Error("Map failed with ", ys)
		}
	}()
}

func TestFilter(t *testing.T) {
	func() {
		xs := []int{2, 4, 5, 6}
		ys := Filter(xs, func(n int) bool {
			return n > 3
		}).([]int)
		if !reflect.DeepEqual(ys, []int{4, 5, 6}) {
			t.Error("filter failed with ", ys)
		}
	}()
}

func TestFoldl(t *testing.T) {
	func() {
		xs := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		sum := Foldl(xs, 0, func(cur, prev int) int { return cur + prev }).(int)
		if sum != 55 {
			t.Error("Should be 55 but got ", sum)
		}
	}()
}

func TestAll(t *testing.T) {
	func() {
		xs := []int{2, 4, 5, 6}
		isEven := func(n int) bool {
			return n%2 == 0
		}
		res := All(xs, isEven)
		if res == true {
			t.Error("Should be false")
		}
	}()

	func() {
		xs := []byte("abcdef")
		isAlpha := func(b byte) bool {
			return b <= 'z' && b >= 'a'
		}
		res := All(xs, isAlpha)
		if res == false {
			t.Error("Should be true")
		}
	}()
}

func TestExists(t *testing.T) {
	func() {
		xs := []int{2, 4, 5, 6}
		isEven := func(n int) bool {
			return n%2 == 0
		}
		res := Exists(xs, isEven)
		if res == false {
			t.Error("Should be true")
		}
	}()

	func() {
		xs := []byte("abcdef")
		notAlpha := func(b byte) bool {
			return !(b <= 'z' && b >= 'a')
		}
		res := Exists(xs, notAlpha)
		if res == true {
			t.Error("Should be false")
		}
	}()
}

func TestIn(t *testing.T) {
	func() {
		xs := []int{2, 4, 5, 6}

		res := In(4, xs)
		if res == false {
			t.Error("Should be true")
		}

		res = In(1, xs)
		if res == true {
			t.Error("Should be false")
		}
	}()
}
