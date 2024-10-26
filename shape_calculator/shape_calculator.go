package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type shape interface{
	area() float64
	perimeter() float64
}

type rectangle struct {
	width float64
	height float64
}

func (r rectangle) area() float64 {
	return r.width * r.height
}

func (r rectangle) perimeter() float64 {
	return 2 * (r.width + r.height)
}

type square struct {
	side float64
}

func (s square) area() float64 {
	return s.side * s.side
}

func (s square) perimeter() float64 {
	return 4 * s.side
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return 3.14 * c.radius * c.radius
}

func (c circle) perimeter() float64 {
	return 2 * 3.14 * c.radius
}

type triangle struct {
	base float64
	height float64
}

func (t triangle) area() float64 {
	return 0.5 * t.base * t.height
}

func (t triangle) perimeter() float64 {
	return 3 * t.base
}


func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Choose a shape (rectangle, square, circle, triangle):")
	shapeType, _ := reader.ReadString('\n')
	shapeType = strings.TrimSpace(shapeType)

	var shape shape

	switch shapeType {
	case "rectangle":
		fmt.Println("Enter width:")
		widthStr, _ := reader.ReadString('\n')
		width, _ := strconv.ParseFloat(strings.TrimSpace(widthStr), 64)

		fmt.Println("Enter height:")
		heightStr, _ := reader.ReadString('\n')
		height, _ := strconv.ParseFloat(strings.TrimSpace(heightStr), 64)

		shape = rectangle{width: width, height: height}

	case "square":
		fmt.Println("Enter side:")
		sideStr, _ := reader.ReadString('\n')
		side, _ := strconv.ParseFloat(strings.TrimSpace(sideStr), 64)

		shape = square{side: side}

	case "circle":
		fmt.Println("Enter radius:")
		radiusStr, _ := reader.ReadString('\n')
		radius, _ := strconv.ParseFloat(strings.TrimSpace(radiusStr), 64)

		shape = circle{radius: radius}

	case "triangle":
		fmt.Println("Enter base:")
		baseStr, _ := reader.ReadString('\n')
		base, _ := strconv.ParseFloat(strings.TrimSpace(baseStr), 64)

		fmt.Println("Enter height:")
		heightStr, _ := reader.ReadString('\n')
		height, _ := strconv.ParseFloat(strings.TrimSpace(heightStr), 64)

		shape = triangle{base: base, height: height}

	default:
		fmt.Println("Invalid shape type")
		return
	}

	fmt.Println("Area: ", shape.area())
	fmt.Println("Perimeter: ", shape.perimeter())
}
