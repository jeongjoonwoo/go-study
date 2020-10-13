package main

import "fmt"

const g float32 = 9.8

type calType func(float32, float32) float32

func calMechEnergy(f calType, m float32, v float32) float32 {
	result := f(m, v)
	return result
}

func main() {
	var m, v, h float32
	fmt.Scanln(&m, &v, &h)

	kinEnergy := func(m float32, v float32) float32 {
		return m * v * v / 2
	}
	potEnergy := func(m float32, h float32) float32 {
		return m * g * h
	}

	ke := calMechEnergy(kinEnergy, m, v)
	pe := calMechEnergy(potEnergy, m, h)
	fmt.Println(ke, pe, ke+pe)
}
