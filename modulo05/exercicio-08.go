// Inverter valores
// Given a set of numbers, return the additive inverse of each. Each positive becomes negatives, and the negatives become positives.

// Lógica:
// Passo 1: Multiplicar por -1 => n * (-1)
// Passo 2: Percorrer o array usando o for range
// Passo 3: Usar o append para adicionar mais um item ao array

package kata

func Invert(arr []int) []int {
	var invertNumbers []int

	for _, number := range arr {
		invertNumbers = append(invertNumbers, number*(-1))
	}

	return invertNumbers
}