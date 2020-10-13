package main

import "fmt"

var g float32 = 9.8

type mul struct {
	m  float32
	v  float32
	h  float32
	ke float32
	pe float32
	me float32
}

func (data *mul) keMul() float32 {
	return data.m * data.v * data.v / 2
}

func (data *mul) peMul() float32 {
	return data.h * data.m * g
}

func main() {
	var objectNum int
	var m, v, h float32

	fmt.Scanln(&objectNum)

	object := make([]mul, objectNum)

	for i := 0; i < objectNum; i++ {
		fmt.Scanln(&m, &v, &h)
		object[i].m = m
		object[i].v = v
		object[i].h = h
		object[i].ke = object[i].keMul()
		object[i].pe = object[i].peMul()
	}

	for _, o := range object {
		fmt.Println(o.ke, o.pe, o.ke+o.pe)
	}

}
