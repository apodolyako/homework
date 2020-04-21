package main

import "fmt"

func main() {
	var sumDep float64
	var percent float64

	fmt.Println("Введите сумму депозита:")
	fmt.Scanln(&sumDep)
	fmt.Println("Банковский процент:")
	fmt.Scanln(&percent)

	fmt.Println("Сумма вклада верез 5 лет:", sumDep+sumDep*percent/100*5)

}
