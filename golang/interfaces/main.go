package main

import (
	"fmt"

	"github.com/dkkahn10/assignments/golang/interfaces/models"
)

type geometry interface {
	Perimeter() float64
	Area() float64
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.Area())
	fmt.Println(g.Perimeter())
}

func main() {
	newCircle := models.Circle{
		Radius: 24,
	}

	newRectangle := models.Rectangle{
		Length: 11,
		Width:  4,
	}

	measure(newCircle)
	measure(newRectangle)
}
