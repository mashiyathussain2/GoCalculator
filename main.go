package main

import (
	"fmt"
	"math"
)

func Add(x, y int) int {
	return x + y
}

func Subtract(x, y int) int {
	return x - y
}

func Multiply(x, y int) int {
	return x * y
}

func Division(x, y int) (string int) {
	if y < 0 && x <= 0 {
		return 0
	} else if y == 0 {
		fmt.Println("Denominator should not be zero")
		return int(math.NaN())
	}
	return x / y
}
func Square(x float32) float32 {
	if x == 0 {
		return 0
	}
	return x * x
}
func Name(x string) string {
	return x
}
func Password(x string) string {
	return x
}
func main() {

}
