package main

import (
	"calculator/math"
	"fmt"

	"github.com/labstack/echo/v4"
)

func main() {

	// Echo instance
	e := echo.New()

	soma := math.Soma(12, 6)
	fmt.Println(soma)

	subtracao := math.Subtracao(22, 3)
	fmt.Println(subtracao)
}