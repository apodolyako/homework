package main

import (
	"fmt"
)

func inputData(msg string) (num int) {

	fmt.Println(msg)
	fmt.Scanln(&num)
	return num
}

func evenFix(num int) string {
	if num%2 == 0 {
		return "четное"
	}
	return "не четное"
}

func main() {
	var num int
	var str string

	num = inputData("введите число:")

	str = evenFix(num)

	fmt.Println("Введенное число ", num, "-", str)

}
