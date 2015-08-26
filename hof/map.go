package hof

import (
	"github.com/BurntSushi/ty"
	"reflect"
)

func Keys(m interface{}) interface{} {
	chk := ty.Check(new(func(map[ty.A]ty.B) []ty.A), m)
	vm := chk.Args[0]
	tkeys := chk.Returns[0]
	vkeys := reflect.MakeSlice(tkeys, vm.Len(), vm.Len())
	for i, vkey := range vm.MapKeys() {
		vkeys.Index(i).Set(vkey)
	}
	return vkeys.Interface()
}

func Values(m interface{}) interface{} {
	chk := ty.Check(new(func(map[ty.A]ty.B) []ty.B), m)
	vm := chk.Args[0]
	tvals := chk.Returns[0]
	vvals := reflect.MakeSlice(tvals, vm.Len(), vm.Len())
	for i, vkey := range vm.MapKeys() {
		vvals.Index(i).Set(vm.MapIndex(vkey))
	}
	return vvals.Interface()
}
