/*
Написать многопоточную программу,
в которой будет использоваться явный вызов планировщика.
Выполните трассировку программы
 */

package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"runtime/trace"
	"sync"
)

func main() {
	f, err := os.Create("trace.proff")
	if err != nil {
		log.Fatal("File didn’t created")
	}
	defer f.Close()

	if err := trace.Start(f); err != nil {
		log.Fatal("Trace didn’t started")
	}
	defer trace.Stop()

	//fmt.Println(runtime.NumCPU()) // проверка числа ядер процессора

	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for c := 'a'; c <= 'z'; c += 1 {
			fmt.Printf("%c", c)
			runtime.Gosched()// вызов планировщика
		}
	}()
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i += 1 {
			fmt.Printf("%d", i)
			runtime.Gosched()// вызов планировщика
		}
	}()

	wg.Wait()
}
