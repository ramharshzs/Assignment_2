package main

import "fmt"

type car struct {
	company string
	wheels  string
}

type bike struct {
	company string
	wheels  string
}

type truck struct {
	company string
	wheels  string
}

func determineType(v interface{}) {

	switch v.(type) {
	case car:
		fmt.Println("A car struct")
	case bike:
		fmt.Println("A bike struct")
	case truck:
		fmt.Println("A truck struct")
	default:
		fmt.Println("Unknown vehicle")
	}

}

func createVehicle() {
	x := car{company: "BMW", wheels: "4 wheels"}
	y := bike{company: "BMW", wheels: "2 wheels"}
	z := truck{company: "BMW", wheels: "10 wheels"}
	determineType(x)
	determineType(y)
	determineType(z)
}
