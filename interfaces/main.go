package main

import (
	"fmt"
	"interfaces/shapes"
)

func main() {

	s := shapes.Square{Side: 2}
	r := shapes.Rectangle{Length: 2, Breadth: 3}
	c := shapes.Circle{Radius: 2}

	fmt.Printf("\trectangle: %.3f\n", area(r))
	fmt.Printf("\tsquare: %.3f\n", area(s))
	fmt.Printf("\tcircle: %.3f\n", area(c))
}

// we need a single function to do this

func area(a shapes.Area) float64 {
	return a.Area()
}

/// WHAT ARE THE ADVANTAGES HERE ?
