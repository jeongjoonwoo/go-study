package main

import (
	"fmt"
	"math"
)

type cal interface {
	volume() float64
	surfaceArea() float64
}

type Cube struct {
	horizontal float64
	vertical   float64
	height     float64
}

type Cylinder struct {
	radius float64
	height float64
}

func (cu Cube) volume() float64 {
	return cu.horizontal * cu.vertical * cu.height
}

func (cu Cube) surfaceArea() float64 {
	return 2 * (cu.height*cu.horizontal + cu.height*cu.vertical + cu.vertical*cu.horizontal)
}

func (cy Cylinder) volume() float64 {
	return (cy.radius * cy.radius * math.Pi) * cy.height
}

func (cy Cylinder) surfaceArea() float64 {
	return (cy.radius*cy.radius*math.Pi)*2 + 2*cy.radius*math.Pi*cy.height
}

func printMeasure(m ...cal) {
	for _, value := range m {
		fmt.Printf("%0.2f, %0.2f\n", value.surfaceArea(), value.volume())
	}
}

func main() {
	cy1 := Cylinder{10, 10}
	cy2 := Cylinder{4.2, 15.6}
	cu1 := Cube{10.5, 20.2, 20}
	cu2 := Cube{4, 10, 23}

	printMeasure(cy1, cy2, cu1, cu2)
}
