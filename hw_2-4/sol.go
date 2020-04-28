package main

import "fmt"

func initX(n int) []int {
	var x []int
	for i := 0; i < n; i++ {
		x = append(x, i+1)
	}
	return x
}

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

//удаляю элементы среза кратные p
func deleteX(x []int, p int) []int {
	for j := 2; j <= x[len(x)-1]/2; j++ {
		for i := 0; i < len(x); i++ {
			if x[i] == p*j {
				x = remove(x, i)
				break
			}
		}
	}
	return x
}

//возвращаю в p след элемент >2
func nextP(x []int, p int) int {
	next := 0
	for i := 0; i < len(x); i++ {
		if x[i] > p {
			next = x[i]
			break
		}
	}
	return next
}

func simpleN(x []int) []int {
	p := 2
	for p != 0 {
		x = deleteX(x, p) //возврат среза без элементов кратных p
		p = nextP(x, p)
	}
	return x
}

func main() {

	const n = 525 //задаю кол-во n - натуральных чисел (подбором)..

	x := initX(n) //Инициализирую срез из заданного ряда натуральных чисел

	x = simpleN(x) // Получаем массив из простых чисел не больших по значению n

	fmt.Println("Массив простых чисел", x)
	fmt.Println("Кол-во элементов:", len(x))

}
