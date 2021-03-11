/*
Написать программу,
которая использует мьютекс для безопасного доступа к данным из нескольких потоков.
Выполните трассировку программы
*/

package main

import (
	"fmt"
	"log"
	"os"
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

	const flow = 1000
	wg := sync.WaitGroup{}
	wg.Add(flow)

	var (
		count int
		mutex sync.Mutex
	)

	for i := 0; i < flow; i++ {

		go func() {
			defer wg.Done()
			mutex.Lock()
			count++
			mutex.Unlock()
		}()
	}

	wg.Wait()
	fmt.Println(count)

}
