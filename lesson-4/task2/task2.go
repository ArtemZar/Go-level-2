/*
Написать программу, которая при получении в канал сигнала SIGTERM
останавливается не позднее, чем за одну секунду (установить таймаут).
*/
package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

)

func main() {

	signalChannel := make(chan os.Signal, 1)

	signal.Notify(signalChannel, syscall.SIGTERM)

	fmt.Println("waiting signal for kill process ", os.Getpid())

	workersJob()

	// остановка программы по сигналу. workersJob завершает счет.
	 go func() {
		sig := <-signalChannel
		fmt.Printf("\ncaught the signal: %v\n", sig)
		fmt.Println("exit aftet 1 second")


	}()

	// имитация сигнала kill
	go func() {
		defer func() {
		signalChannel<-syscall.SIGTERM
		}()
	}()

	time.Sleep(1 * time.Second)
}

// workersJob имитация деятельности, счет от 1 до 10
//
// пне принимает не возвращает никаких значений
func workersJob() {
	var workers = make(chan struct{}, 1)

	for i := 1; i <= 10; i++ {
		workers <- struct{}{}

		go func(job int){
			defer func() {
				<-workers
			}()

			time.Sleep(1 * time.Second)
			fmt.Print(job, " ")
		}(i)
	}
}