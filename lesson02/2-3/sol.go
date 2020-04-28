package main

import "fmt"

func calcF(num int) int {

	switch num {
	case 0:
		return 0
	case 1:
		return 1
	}
	return calcF(num-1) + calcF(num-2)
}

func main() {
	const n = 100

	for i := 0; i < n; i++ {
		fmt.Println(calcF(i))
	}

}
