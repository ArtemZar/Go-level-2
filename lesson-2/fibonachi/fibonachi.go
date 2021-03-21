// Package main implements functions to calculate the Fibonacci number
//
// The FindFibonachiElement returns some Fibonacci number
//
// FindFibonachiElement(uint32) uint32
//
// So, you are allowed to call this function and get some Fibonacci number
package fibonachi

import (
	"fmt"
)

var mapFibonachi = map[uint32]uint32{
	0: 0,
	1: 1,
	2: 1,
}

func Fibonachi() {
	var fibonachiNumber uint32
	for true {
		fmt.Scan(&fibonachiNumber)
		_, exist := mapFibonachi[fibonachiNumber]
		if exist == true {
			fmt.Println(mapFibonachi[fibonachiNumber])
		} else {
			fmt.Println(FindFibonachiElement(fibonachiNumber))
		}
	}
}

// FindFibonachiElement рекурсивная функция с использованием мапы
//
// принимает порядковый номер числа Фибоначчи (тип uint32)
//
// возвращает само число Фибоначчи (тип uint32)
func FindFibonachiElement(fnum uint32) uint32 {
	if fnum == 0 {
		return 0
	} else if fnum == 1 || fnum == 2 {
		return 1
	} else {
		mapFibonachi[fnum] = FindFibonachiElement(fnum-2) + FindFibonachiElement(fnum-1)
		return mapFibonachi[fnum]
	}
}
