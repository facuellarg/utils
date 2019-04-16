package utils

//Combinatoria retorna una matriz con.
//Todas las posibles combinaciones de longitud a.
func Combinatoria(numero int) [][]int {
	var arr [][]int
	if numero < 2 {
		return [][]int{}
	}
	if numero == 2 {
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
	a := Combinatoria(numero - 1)
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
