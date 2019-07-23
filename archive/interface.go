package main

import "fmt"

type Vehicle interface {
	Move()
}

type Car struct {
	Speed uint32
}

func (car *Car) Move() {
	fmt.Printf("Car moves at speed %dkm/h\n", car.Speed)
}

type Plane struct {
	Speed uint32
}

func (plane *Plane) Move() {
	fmt.Printf("Plane moves at speed %dkm/h\n", plane.Speed)
}

func parse(v interface{}) {
	switch v.(type) {
	case Car:
		fmt.Println("v is a car")
	case Plane:
		fmt.Println("v is a plane")
	}
}

func main() {
	car := Car{Speed: 100}
	plane := Plane{Speed: 600}
	car.Move()
	plane.Move()
	parse(car)
	parse(plane)
}
