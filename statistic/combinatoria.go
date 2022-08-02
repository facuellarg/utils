package statistic

var elem []string
var combinatoria [][]string

func comb(act []string, n int, r int) {
	if n == 0 {
		aux := make([]string, len(act))
		copy(aux, act)
		combinatoria = append(combinatoria, act)

	} else {

		for i := 0; i < r; i++ {
			if !contains(act, elem[i]) {
				comb(append(act, elem[i]), n-1, r)
			}
		}
	}
}

func contains(container []string, elemento string) bool {
	for _, v := range container {
		if v == elemento {
			return true
		}
	}
	return false
}

//Combinatoria retorna una matriz con
//Todas las posibles combinaciones de longitud a
//Basado en el arreglo elementos
func Combinatoria(elementos []string, tamaño int) [][]string {
	elem = elementos
	var act []string
	if tamaño > len(elementos) {
		return nil
	}
	comb(act, tamaño, len(elem))
	return combinatoria
}
