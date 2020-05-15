package calculator

import (
	// Пакет для работы с ошибками
	"errors"
	// Описания типов для построения бинарных деревьев
	"go/ast"
	// Пакет для обработки исходных данных
	"go/parser"
	// Содержит константы для математической лексики +, -, /, * и т.д.
	"go/token"
	// Математические константы и выражения
	"math"
	// Конвертер строк
	"strconv"
	// Приведение имен констант к единому виду
	"strings"
)

// Calculate рассчитает данные выражения, переданного в строке
func Calculate(expr string) (float64, error) {
	// Поиск управляющих операторов в строке и построение дерева
	// Вернет корневой элемент в переменную root все математические действия, построенные
	// в виде дерева
	// Node будет содержать в себе действие для вычисления
	// exprNode() даст доступ к ветвям дерева, исходящим из корня/ноды
	//
	// type Expr interface {
	//  Node
	//  exprNode()
	// }
	root, err := parser.ParseExpr(expr)

	if err != nil {
		return -1, err
	} else {
		return eval(root)
	}
}

type Func struct {
	Name string
	Args int
	Func func(args ...float64) float64
}

var funcMap map[string]Func

// Создание массива функций для обработки данных
func init() {
	funcMap = make(map[string]Func)
	funcMap["sqrt"] = Func{
		Name: "sqrt",
		Args: 1,
		Func: func(args ...float64) float64 {
			return math.Sqrt(args[0])
		},
	}
	funcMap["abs"] = Func{
		Name: "abs",
		Args: 1,
		Func: func(args ...float64) float64 {
			return math.Abs(args[0])
		},
	}
	funcMap["log"] = Func{
		Name: "log",
		Args: 2,
		Func: func(args ...float64) float64 {
			return math.Log(args[0]) / math.Log(args[1])
		},
	}
	funcMap["ln"] = Func{
		Name: "ln",
		Args: 1,
		Func: func(args ...float64) float64 {
			return math.Log(args[0])
		},
	}
	funcMap["sin"] = Func{
		Name: "sin",
		Args: 1,
		Func: func(args ...float64) float64 {
			return math.Sin(args[0])
		},
	}
	funcMap["cos"] = Func{
		Name: "cos",
		Args: 1,
		Func: func(args ...float64) float64 {
			return math.Cos(args[0])
		},
	}
	funcMap["tan"] = Func{
		Name: "tan",
		Args: 1,
		Func: func(args ...float64) float64 {
			return math.Tan(args[0])
		},
	}
	funcMap["arcsin"] = Func{
		Name: "arcsin",
		Args: 1,
		Func: func(args ...float64) float64 {
			return math.Asin(args[0])
		},
	}
	funcMap["arccos"] = Func{
		Name: "arccos",
		Args: 1,
		Func: func(args ...float64) float64 {
			return math.Acos(args[0])
		},
	}
	funcMap["arctan"] = Func{
		Name: "arctan",
		Args: 1,
		Func: func(args ...float64) float64 {
			return math.Atan(args[0])
		},
	}
	funcMap["max"] = Func{
		Name: "max",
		Args: 2,
		Func: func(args ...float64) float64 {
			return math.Max(args[0], args[1])
		},
	}
	funcMap["min"] = Func{
		Name: "min",
		Args: 2,
		Func: func(args ...float64) float64 {
			return math.Min(args[0], args[1])
		},
	}
}

func eval(expr ast.Expr) (float64, error) {
	switch expr.(type) {
	case *ast.BasicLit:
		return basic(expr.(*ast.BasicLit))
	// Бинарные выражения
	case *ast.BinaryExpr:
		return binary(expr.(*ast.BinaryExpr))
	// Вложенные вычисления
	case *ast.ParenExpr:
		return eval(expr.(*ast.ParenExpr).X)
	// Обработка посредством математических функций
	case *ast.CallExpr:
		return call(expr.(*ast.CallExpr))
	// Случай для обработки констант
	case *ast.Ident:
		return ident(expr.(*ast.Ident))
	default:
		return -1, errors.New("Не удалось распознать оператор")
	}
}

func basic(lit *ast.BasicLit) (float64, error) {
	switch lit.Kind {
	case token.INT:
		i, err := strconv.ParseInt(lit.Value, 10, 64)

		if err != nil {
			return -1, err
		} else {
			return float64(i), nil
		}
	case token.FLOAT:
		i, err := strconv.ParseFloat(lit.Value, 64)

		if err != nil {
			return -1, err
		} else {
			return i, nil
		}
	default:
		return -1, errors.New("Неизвестный аргумент")
	}
}

func binary(expr *ast.BinaryExpr) (ret float64, err error) {
	x, err1 := eval(expr.X)
	y, err2 := eval(expr.Y)
	ret = -1

	if (err1 == nil) && (err2 == nil) {

		switch expr.Op {
		case token.ADD:
			ret = x + y
		case token.SUB:
			ret = x - y
		case token.MUL:
			ret = x * y
		case token.QUO:
			ret = x / y
		case token.REM:
			ret = float64(int64(x) % int64(y))
		case token.AND:
			ret = float64(int64(x) & int64(y))
		case token.OR:
			ret = float64(int64(x) | int64(y))
		case token.XOR:
			ret = math.Pow(x, y)
		default:
			err = errors.New("Неизвестный бинарный оператор")
		}
	} else {
		if err1 != nil {
			err = err1
		} else {
			err = err2
		}
	}

	return
}

func ident(id *ast.Ident) (float64, error) {
	switch n := strings.ToLower(id.Name); n {
	case "pi":
		return math.Pi, nil
	case "e":
		return math.E, nil
	case "phi":
		return math.Phi, nil
	default:
		return -1, errors.New("Неизвестная константа " + n)
	}
}


// Обработка функциональных операторов с помощью созданного массива функций для обработки
func call(c *ast.CallExpr) (float64, error) {
    switch t := c.Fun.(type) {
    case *ast.Ident:
    default:
        _ = t
        return -1, errors.New("Неизвестный тип функции")
    }

    ident := c.Fun.(*ast.Ident)

    args := make([]float64, len(c.Args))
    for i, expr := range c.Args {
        var err error
        args[i], err = eval(expr)
        if err != nil {
            return -1, err
        }
    }

    name := strings.ToLower(ident.Name)

    if val, ok := funcMap[name]; ok {
        if len(args) == val.Args {
            return val.Func(args...), nil
        } else {
            return -1, errors.New("Слишком много аргументов для " + name)
        }
    } else {
        return -1, errors.New("Неизвестная функция " + name)
    }
}


package main

import (
    "calculator"
    "fmt"
)

func main() {
    input := ""
    for {
        fmt.Print("> ")
        if _, err := fmt.Scanln(&input); err != nil {
            fmt.Println(err)
            continue
        }

        if input == "exit" {
            break
        }

        if res, err := calculator.Calculate(input); err == nil {
            fmt.Printf("Результат: %v\n", res)
        } else {
            fmt.Println("Не удалось произвести вычисление")
        }
    }
}



// package main

// import "fmt"

// type phones []int

// func (p phones) ViewList() {
//     for i, phone := range p {
//         fmt.Printf("\t %v) %v \n", i, phone)
//     }}
// }

// func main() {
//     addressBook := make(map[string]phones)

//     addressBook["Миша"] = phones{78293467382}
//     addressBook["Никита"] = phones{89167253764, 89635437382}

//     for name, ph := range addressBook {
//         fmt.Println(name)
//         ph.ViewList()
//     }
// }

// type Square struct {
//     edge float64
// }

// func (s Square) Area() float64 {
//     return s.edge * s.edge
// }

// type Circle struct {
//     radius float64
// }

// func (c Circle) Area() float64 {
//     return math.Pi * radius
// }

// func SummAreas(areas ...float64) float64 {
//     res := 0.0
//     for _, area := range areas {
//         res += area
//     }
//     return res
// }

// type Shape interface {
// 	Area() float64

// }

// func SummAreas(shapes ...Shape) float64 {
//     res := 0.0
//     for _, shape := range shapes {
//         if shape == nil {
//             continue
//         }
//         res += shape.Area()
//     }
//     return res
// }

// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"log"
// )

// type Message struct {
// 	Name string
// 	Body string
// 	Time int64
// }

// func main() {
// 	jsonString := "{\"Name\": \"Alice\", \"Body\": \"Hello\", \"Time\": 1294706395881547000}"
// 	b := []byte(jsonString)
// 	m := Message{}

// 	err := json.Unmarshal(b, &m)
// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}

// 	fmt.Println(m)
// }

// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"log"
// )

// type Message struct {
// 	Name string
// 	Body string
// 	Time int64
// }

// func main() {
// 	m := Message{"Alice", "Hello", 1294706395881547000}

// 	b, err := json.Marshal(m)
// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}

// 	fmt.Println(b)

// 	fmt.Println(string(b))
// }

// package main

// import (
// 	"fmt"
// 	// "stack"
// )

// func main() {
// 	addressBook := make(map[string][]int)

// 	addressBook["Alex"] = []int{89996543210}
// 	addressBook["Bob"] = []int{89167243812}
// 	addressBook["Bob"] = append(addressBook["Bob"], 89155243627)

// 	fmt.Println(addressBook)

// 	for name, numbers := range addressBook {
// 		fmt.Println("Абонент:", name)
// 		for i, number := range numbers {
// 			fmt.Printf("\t %v) %v \n", i+1, number)
// 		}
// 	}
// }

// func main() {
// 	stack.Push("Этот текст")
// 	stack.Push("Будет находиться в стеке")
// 	stack.Push("До первого обращения к pop")

// 	fmt.Println(stack.Pop())
// 	fmt.Println(stack.Pop())

// 	stack.Push("Добавим еще текста")

// 	fmt.Println(stack.Pop())
// 	fmt.Println(stack.Pop())

// }
