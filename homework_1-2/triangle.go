package main

import (
	"fmt"
	"math"
)

func main() {

	var sizeA float64 = 10
	var sizeB float64 = 20
	var sizeC float64

	sizeC = math.Sqrt(math.Pow(sizeA, 2) + math.Pow(sizeB, 2))
	fmt.Println("Дано: прямоугольный треугольник с катетами A=", sizeA, " В=", sizeB)
	fmt.Println("Тогда: ")
	fmt.Println("Площадь треугольника: S=", sizeA*sizeB/2, "Гепотенуза: C=", sizeC, " Периметр: P=", sizeA+sizeB+sizeC)

}
