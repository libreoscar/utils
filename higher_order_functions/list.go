package hof

import (
	"github.com/BurntSushi/ty"
	"reflect"
)

func Map(xs, f interface{}) interface{} {
	chk := ty.Check(
		new(func(func(ty.A) ty.B, []ty.A) []ty.B),
		f, xs)
	vf, vxs, tys := chk.Args[0], chk.Args[1], chk.Returns[0]

	xsLen := vxs.Len()
	vys := reflect.MakeSlice(tys, xsLen, xsLen)
	for i := 0; i < xsLen; i++ {
		vy := vf.Call([]reflect.Value{vxs.Index(i)})[0]
		vys.Index(i).Set(vy)
	}
	return vys.Interface()
}

func Filter(xs, p interface{}) interface{} {
	chk := ty.Check(
		new(func(func(ty.A) bool, []ty.A) []ty.A),
		p, xs)
	vp, vxs, tys := chk.Args[0], chk.Args[1], chk.Returns[0]

	xsLen := vxs.Len()
	vys := reflect.MakeSlice(tys, 0, xsLen)
	for i := 0; i < xsLen; i++ {
		vx := vxs.Index(i)
		vy := vp.Call([]reflect.Value{vx})[0]
		if vy.Bool() {
			vys = reflect.Append(vys, vx)
		}
	}
	return vys.Interface()
}

func Foldl(xs, init, f interface{}) interface{} {
	chk := ty.Check(
		new(func(func(ty.A, ty.B) ty.B, ty.B, []ty.A) ty.B),
		f, init, xs)
	vf, vinit, vxs, tb := chk.Args[0], chk.Args[1], chk.Args[2], chk.Returns[0]

	xsLen := vxs.Len()
	vb := reflect.New(tb).Elem()
	vb.Set(vinit)

	for i := 0; i < xsLen; i++ {
		// TODO (Rex): make sure order is right
		vb.Set(vf.Call([]reflect.Value{vb, vxs.Index(i)})[0])
	}
	return vb.Interface()
}

func All(xs, f interface{}) bool {
	chk := ty.Check(
		new(func(func(ty.A) bool, []ty.A) bool),
		f, xs)
	vf, vxs := chk.Args[0], chk.Args[1]

	xsLen := vxs.Len()
	for i := 0; i < xsLen; i++ {
		vy := vf.Call([]reflect.Value{vxs.Index(i)})[0]
		if !vy.Bool() {
			return false
		}
	}
	return true
}

func Exists(xs, f interface{}) bool {
	chk := ty.Check(
		new(func(func(ty.A) bool, []ty.A) bool),
		f, xs)
	vf, vxs := chk.Args[0], chk.Args[1]

	xsLen := vxs.Len()
	for i := 0; i < xsLen; i++ {
		vy := vf.Call([]reflect.Value{vxs.Index(i)})[0]
		if vy.Bool() {
			return true
		}
	}
	return false
}

func In(x, xs interface{}) bool {
	chk := ty.Check(
		new(func(ty.A, []ty.A) bool),
		x, xs)
	vxs := chk.Args[1]

	length := vxs.Len()
	for i := 0; i < length; i++ {
		if reflect.DeepEqual(vxs.Index(i).Interface(), x) {
			return true
		}
	}
	return false
}
