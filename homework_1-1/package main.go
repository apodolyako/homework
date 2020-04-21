package main

import (
	"fmt"
	"math"
)

func main() {
	const exchRate float64 = 75.25
	var num float64
	var cent float64
	// num = 45.5

	fmt.Println("Введите сколько рублей поменять?")
	fmt.Scanln(&num)
	// fmt.Println("Привет", num)
	sum := num / exchRate
	cent = math.Round((sum - math.Round(num/exchRate)) * 100)

	fmt.Println("Курсу 1$ = ", exchRate, "; Итого Долларов: ", math.Round(sum), "и ", cent, " центов")
}
