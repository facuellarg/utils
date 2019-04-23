package utils

var elem [10]int
var combinatoria [][]int

func comb(act []int, n int, r int) {

	if n == 0 {
		combinatoria = append(combinatoria, act)
	} else {
		for i := 0; i < r; i++ {
			if !contains(act, elem[i]) { // Controla que no haya repeticiones
				comb(append(act, elem[i]), n-1, r)
			}
		}
	}
	return
}

func contains(container []int, elemento int) bool {
	for _, v := range container {
		if v == elemento {
			return true
		}
	}
	return false
}

//Combinatoria retorna una matriz con.
//Todas las posibles combinaciones de longitud a.
func Combinatoria(tamaño int) [][]int {
	elem = [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	var act []int
	comb(act, tamaño, len(elem))
	return combinatoria
}
