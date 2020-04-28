package main

import (
	"fmt"
)

func inputData(msg string) (num int) {

	fmt.Println(msg)
	fmt.Scanln(&num)
	return num
}

func evenFix(num int, split int) string {
	if num%split == 0 {
		return "делиться без остатка на"
	}
	return "не делиться без остатка на"
}

func main() {
	var num int
	var str string
	const spliter = 3

	num = inputData("введите число:")

	str = evenFix(num, spliter)

	fmt.Println("Введенное число ", num, "-", str, spliter)

}
