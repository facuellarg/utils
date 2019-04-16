package utils

//Combinatoria retorna una matriz con.
//Todas las posibles combinaciones de longitud a.
func Combinatoria(a int) [][]int {
	return combinatoriaArr(make([]int, a))
}

func combinatoriaArr(numero []int) [][]int {
	var arr [][]int
	if len(numero) == 2 {
		for i := 0; i < 10; i++ {
			for j := 0; j < 10; j++ {
				if i != j {
					arr = append(arr, []int{i, j})
				}
			}
		}
		return arr
	}
	agregar := false
	a := combinatoriaArr(numero[1:])
	for i := 0; i < 10; i++ {
		for j := 0; j < len(a); j++ {
			for k := 0; k < len(a[j]); k++ {
				if i == a[j][k] {
					agregar = false
				}
			}
			if agregar {
				arr = append(arr, append(a[j], i))
			}
			agregar = true
		}
	}
	return arr
}
