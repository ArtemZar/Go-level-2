/*
Смоделировать ситуацию “гонки”,
и проверить программу на наличии “гонки”
*/

package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	const count = 1000
	var (
		counter int // общий экземпляр - счетчик
	)

	f, err := os.Create("output.txt")
	if err != nil {
		log.Fatal("File didn’t created")
	}
	defer f.Close()

	for i := 0; i < count; i += 1 {
		go func() {
			counter += 1 // Выполняем инкремент при выполнении потока
			f.WriteString(strconv.Itoa(counter))
		}()
	}

	for i := 0; i < count; i += 1 {
		go func() {
			counter -= 1 // Выполняем декремент при выполнении потока
			f.WriteString(strconv.Itoa(counter))
		}()
	}
	time.Sleep(5 * time.Second)
	fmt.Println(counter)
}
