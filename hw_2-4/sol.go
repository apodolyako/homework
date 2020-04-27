package main

import "fmt"

// удаляем выбранный  те i-ый элемент среза
func remove(s []int, i int) []int {
	var buf []int

	//проверяю если переданный элемент последний тогда в buf копирую все кроме последнего элементы
	if i == len(s)-1 {
		for j := 0; j < len(s)-1; j++ {
			buf = append(buf, s[j])
		}

	} else {
		for j := 0; j < len(s); j++ {

			if j == i {
				j = j + 1
			}
			buf = append(buf, s[j])
		}
	}
	return buf
}

func main() {

	const n = 21
	var x []int

	for i := 0; i < n; i++ {
		x = append(x, i+1)
	}

	for j := 2; j <= x[len(x)-1]/2; j++ {
		for i := 0; i < len(x); i++ {
			if x[i] == 2*j {
				x = remove(x, i)
				break
			}
		}
	}

	fmt.Println("Массив простых чисел", x, "Кол-во элементов:", len(x))

}
