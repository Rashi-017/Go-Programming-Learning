package main

import "fmt"

type MoterVehicle interface {
	milage() float64
}

type BMW struct {
	distance float64
	speed    float64
}

func (b BMW) milage() float64 {
	return b.distance * b.speed
}

type OD struct {
	distance float64
	speed    float64
}

func (o OD) milage() float64 {
	return o.distance * o.speed
}

func totalmilage(m []MoterVehicle) {
	totalmilage := 0.0
	for _, val := range m {
		totalmilage = totalmilage + val.milage()
	}
	fmt.Printf("total milage given  %f ", totalmilage)
}

func function2() {
	b := BMW{120.90, 30.90}
	o := OD{20.00, 40.00}

	person := []MoterVehicle{b, o}
	totalmilage(person)
}
