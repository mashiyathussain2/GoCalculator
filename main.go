package main

import (
	"fmt"
	"math"
	"net/http"
	"os"
	"path"

	"github.com/joho/godotenv"
	"github.com/labstack/echo"
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
func Addo(context echo.Context) error {
	return context.JSON(http.StatusOK, "version")
}

func main() {
	fmt.Println("Go Program")
	server := echo.New()
	server.GET(path.Join("/"), Addo)

	godotenv.Load()
	port := os.Getenv("PORT")

	address := fmt.Sprintf("%s:%s", "0.0.0.0", port)
	fmt.Println(address)
	server.Start(address)

}
