package main

import (
	"calculator/math"
	"fmt"
)

func main() {

	soma := math.Soma(12, 6)
	fmt.Println(soma)

	subtracao := math.Subtracao(22, 3)
	fmt.Println(subtracao)
}