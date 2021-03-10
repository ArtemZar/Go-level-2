// Реализуйте функцию для разблокировки мьютекса с помощью defer

package main

import (
	"errors"
	"fmt"
	"log"
	"sync"
	"time"
)

var mu sync.Mutex
var a int

func main() {

	for i := 0; i < 1000; i++ {
		go func() {
			mu.Lock()
			defer mu.Unlock()
			if false {
				err := errors.New("my error")
				log.Println(err)
				return // альтернативное срабатывание defer (разблокировка)
			}
			a += 1
			return // срабатывание defer (разблокировка)
		}()
	}

	time.Sleep(1 * time.Second)
	fmt.Println(a)
}
