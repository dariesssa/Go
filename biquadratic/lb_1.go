package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Решение биквадратного уравнения")

	var valA, valB, valC float64
	var success bool = false

	if len(os.Args) == 4 {
		var err1, err2, err3 error

		valA, err1 = strconv.ParseFloat(os.Args[1], 64)
		valB, err2 = strconv.ParseFloat(os.Args[2], 64)
		valC, err3 = strconv.ParseFloat(os.Args[3], 64)

		if err1 != nil || err2 != nil || err3 != nil {
			fmt.Println("Ошибка в аргументах, переход к ручному вводу.")
		} else if valA == 0 {
			fmt.Println("Ошибка: из аргументов считан A = 0, переход к ручному вводу")
		} else {
			fmt.Printf("Коэффициенты успешно считаны: A = %v, B = %v, C = %v\n", valA, valB, valC)
			success = true
		}
	}

	if !success {
		for {
			fmt.Println("Ведите коэффициент А: ") // цикл для А
			_, err := fmt.Scanln(&valA)

			if err != nil {
				fmt.Println("Ошибка: введите корректное число.")

				var discard string
				fmt.Scanln(&discard)
				continue
			}

			if valA == 0 {
				fmt.Println("Ошибка: коэффициент A не может быть равен нулю!")
				continue
			}
			break
		}

		for {
			fmt.Println("Введите коэффициент B: ") // цикл для B
			_, err := fmt.Scanln(&valB)

			if err != nil {
				fmt.Println("Ошибка: введите корректное число.")

				var discard string
				fmt.Scanln(&discard)
				continue
			}
			break
		}

		for {
			fmt.Println("Введите коэффициент C: ") // цикл для C
			_, err := fmt.Scanln(&valC)

			if err != nil {
				fmt.Println("Ошибка: введите корректное число.")

				var discard string
				fmt.Scanln(&discard)
				continue
			}
			break
		}
	}

	D := valB*valB - 4*valA*valC //считаем дискриминант и находим корни
	sqrtD := math.Sqrt(D)

	if D < 0 {
		fmt.Println("Дискриминант < 0. Корней нет.")

	} else if D == 0 {
		t := -valB / (2 * valA)

		fmt.Println("Промежуточное t =", t)

		if t < 0 {
			fmt.Println("Корней нет (t < 0)")
		} else if t == 0 {
			fmt.Println("Корень: x = 0")
		} else {
			fmt.Println("Корни: x1 =", math.Sqrt(t), ", x2 =", -math.Sqrt(t))
		}

	} else if D > 0 {
		t1 := (-valB + sqrtD) / (2 * valA)
		t2 := (-valB - sqrtD) / (2 * valA)

		fmt.Printf("Промежуточные корни: t1 = %v, t2 = %v\n", t1, t2)

		switch {
		case t1 < 0:
			fmt.Println("Для t1 корней x нет (t1 < 0)")
		case t1 == 0:
			fmt.Println("Из t1 найден корень: x = 0")
		case t1 > 0:
			fmt.Printf("Из t1 найдены корни: x1 = %v, x2 = %v\n", math.Sqrt(t1), -math.Sqrt(t1))
		}

		switch {
		case t2 < 0:
			fmt.Println("Для t2 корней x нет (t2 < 0)")
		case t2 == 0:
			fmt.Println("Из t2 найден один корень: x = 0")
		case t2 > 0:
			fmt.Printf("Из t2 найдены корни: x1 = %v, x2 = %v\n", math.Sqrt(t2), -math.Sqrt(t2))
		}
	}

}
