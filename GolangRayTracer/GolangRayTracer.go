package main

import (
	"fmt"
	"github.com/mikee385/GolangRayTracer/GolangRayTracer/color"
)

func main() {
	red := color.New(1.0, 0.0, 0.0)
	green := color.New(0.0, 1.0, 0.0)
	result := red.Add(green)
	fmt.Println(result.Red, result.Green, result.Blue)

	fmt.Println(color.New(0.5, 0.5, 0.5))

	fmt.Println(color.White())
}
