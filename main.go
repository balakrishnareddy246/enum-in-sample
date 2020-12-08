package main

import "fmt"

type Direction int

const (
	North Direction = iota
	South
	East
	West
)

func main() {

	// Declaring a variable myDirection with type Direction
	var myDirection Direction
	myDirection = West

	if myDirection == West {
		fmt.Println("myDirection is West:", myDirection)
	}
}
