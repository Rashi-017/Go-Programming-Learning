package main

import "fmt"

type shape interface {
	printArea() float64
}
type triangle struct {
	height float64
	base   float64
}
type square struct {
	sidelength float64
}

func main() {
	t := triangle{
		height: 90.9,
		base:   90.0,
	}
	sq := square{
		sidelength: 60.0,
	}
	t.printArea()

	sq.printArea()

	getArea(t)
	getArea(sq)

}
func getArea(s shape) {
	fmt.Println(s.printArea())
}
func (t triangle) printArea() float64 {
	return t.height * t.base * 0.5
}
func (sq square) printArea() float64 {
	return sq.sidelength * sq.sidelength
}
