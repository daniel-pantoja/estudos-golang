// Listas

// 2) Maps - heterogêneos (pode misturar os tipos)
// var := map[key]value{}

package main

import "fmt"

func main() {
		idade := map[string]int{}
		idade["Daniel"] = 18
		idade["Ramilly"] = 19
		idade["Neymar"] = 33
		fmt.Println(idade)
		fmt.Println(idade["Daniel"])
		fmt.Println(idade["Ramilly"])
		fmt.Println(idade["Neymar"])

		anoNasc := map[string]int{
			"Daniel": 2006,
			"Ramilly": 2005,
			"Neymar": 1992,
		}
		fmt.Println(anoNasc)
		fmt.Println(anoNasc["Daniel"])
		fmt.Println(anoNasc["Ramilly"])
		fmt.Println(anoNasc["Neymar"])
		anoNasc["Miguel"] = 2008
		fmt.Println(anoNasc)
		fmt.Println(anoNasc["Miguel"])
		anoNasc["Futebol"] = 1863 
		fmt.Println(anoNasc)
		fmt.Println(anoNasc["Futebol"])

		altura := map[string]float64{
			"Daniel": 1.64,
			"Ramilly": 1.63,
			"Miguel": 1.67,
		}
		fmt.Println(altura)
		fmt.Println(altura["Daniel"])
		fmt.Println(altura["Ramilly"])
		fmt.Println(altura["Miguel"])
}