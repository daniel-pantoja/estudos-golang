// TIPOS:
// bool (true/false)
// string - sequência de bytes
// int
// float (float64/float32) - decimal

package main

import (
	"fmt"
)

func main() {
	fmt.Printf("Type: %T - Value: %v\n", true, true)
	fmt.Printf("Type: %T - Value: %v\n", "steph", "daniel")
	fmt.Printf("Type: %T - Value: %v\n", 2, 17676)
	fmt.Printf("Type: %T - Value: %v\n", "2", "322")
	fmt.Printf("Type: %T - Value: %v\n", 1.444, 1.555)
}