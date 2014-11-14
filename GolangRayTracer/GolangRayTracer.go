package main

import (
	"fmt"
	"github.com/mikee385/GolangRayTracer/GolangRayTracer/color"
	"github.com/mikee385/GolangRayTracer/GolangRayTracer/table"
)

func main() {
	var red = color.New(1.0, 0.0, 0.0)
	var green = color.New(0.0, 1.0, 0.0)
	var result = red.Add(green)
	fmt.Println(result.Red, result.Green, result.Blue)

	fmt.Println(color.New(0.5, 0.5, 0.5))

	fmt.Println(color.White())

	var table = table.New(8, 6)
	for row := 0; row < table.GetHeight(); row++ {
		for column := 0; column < table.GetWidth(); column++ {
			table.Set(row, column, color.Green())
		}
	}
	for row := 0; row < table.GetHeight(); row++ {
		for column := 0; column < table.GetWidth(); column++ {
			fmt.Printf("%v, %v = %v\n", row, column, table.Get(row, column))
		}
	}
}
