// Listas

// 1) Arrays e Slices - homogêneos (todos os elementos da lista vão ter o mesmo tipo)
// ex:[1, 2, 3, 4, 5, 6] - []int 
// ex:["Daniel", "Junior", "Pantoja"] - []string
// ex:[1.001, 2.002, 3.003] - []float64 

// 2) Maps - heterogêneos (pode misturar os tipos)
// ex:{"dani": 18, "Ramilly": 91}


package main

import "fmt"

func main() {
// Array - tamanho fixo
var array [2]string
array[0] = "Hello"
array[1] = "World"
fmt.Println(array[0])
fmt.Println(array[1])
fmt.Println(array[0], array[1])
fmt.Println(array)

numPrimos := [6]int{2, 3, 5, 7, 11, 13}
fmt.Println(numPrimos)
fmt.Println(numPrimos[0])
fmt.Println(numPrimos[1])
fmt.Println(numPrimos[3:6])
fmt.Println(numPrimos[4])
fmt.Println(numPrimos[5])

//Slice
//var slice []string
slice := make([]string, 5)
slice[0] = "Hello"
slice[1] = "World"
slice[2] = "World"
slice[3] = "World"
fmt.Println(slice[0], slice[1])
fmt.Println(slice)
fmt.Println(slice[0])
fmt.Println(slice[1])
fmt.Println(slice[2])
slice[2] = "Daniel"
fmt.Println(slice[2])
//len
fmt.Println(len(slice))
fmt.Println(slice[2])
fmt.Println(slice[3])
fmt.Println(slice[4])

numPares := []int{2, 4, 6, 8, 10, 12}
fmt.Println(numPares)

numPares = append(numPares, 14)
fmt.Println(numPares)
numPares = append(numPares, 16, 18, 20)
fmt.Println(numPares)
numPares = append(numPares, 16, 18, 20, 22)
fmt.Println(numPares)
numPares = append(numPares, 24, 26, 28, 30)
fmt.Println(numPares)
}