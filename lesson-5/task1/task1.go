// Напишите программу, которая запускает n потоков и дожидается завершения их всех

package main

import (
	"fmt"
	"sync"
)

const flous = 5 // количество потоков

func main() {
	wg := new(sync.WaitGroup)

	for i := 1; i <= flous; i++ {
		wg.Add(1)
		go work(i, wg)
	}

	wg.Wait() // ожидание завершения всех горутин в группе
	fmt.Println("Все потоки завершили выполнение")
}

func work(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Горутина %d начала выполнение \n", id)
	//time.Sleep(1 * time.Second)
	fmt.Printf("Горутина %d завершила выполнение \n", id)
}
