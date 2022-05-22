package main

import (
	"errors"
	"fmt"
	"log"
	"strings"
)

// Interface shape
// Made sure that triangle and square matched the interface by having a function called getArea int
type shape interface {
	getArea() (int, error)
}

// square with a field called sideLength
type square struct {
	sideLength float64
}

// Triangle struct with a field called height and base
type triangle struct {
	height float64
	base   float64
}

// Takes the interface shape as a parameter
func printArea(s shape) {
	fmt.Println(s.getArea())
}

// square method with receiver
func (s square) getArea() (int, error) {
	if s.sideLength <= 0 {
		return 0, errors.New("error: square sides should be greater than 0")
	}
	return int(s.sideLength * s.sideLength), nil
}

// triangle method with receiver
func (t triangle) getArea() (int, error) {
	if t.height <= 0 || t.base <= 0 {
		return 0, errors.New("error: triangle base and height should be greater than 0")
	}

	return int((t.base * t.height) * 0.5), nil
}

func main() {
	fmt.Println("(Area)Enter the shape you want to calculate(triangle/square):")

	// Create a variable and takes the user input
	var userInputShape string
	fmt.Scanln(&userInputShape)

	// User input validation
	if strings.ToLower(userInputShape) == "triangle" {
		var userBase, userHeight int
		fmt.Println("Enter the base of the triangle:")
		fmt.Scanln(&userBase)

		fmt.Println("Enter the height of the triangle:")
		fmt.Scanln(&userHeight)

		userTriangle := triangle{float64(userHeight), float64(userBase)}
		printArea(userTriangle)

	} else if strings.ToLower(userInputShape) == "square" {
		var userSide int
		fmt.Println("Enter the base of the square")
		fmt.Scanln(&userSide)

		userSquare := square{float64(userSide)}
		printArea(userSquare)
	} else {
		// Break the program if the user input is neither above
		log.Fatal("The user input isn't correct")
	}
}
