/*
Написать программу, которая при получении в канал сигнала SIGTERM
останавливается не позднее, чем за одну секунду (установить таймаут).
*/
package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, syscall.SIGTERM)

	doneCh := make(chan error)

	go func() {
		doneCh <- longJobFunction()
	}()

	// имитация сигнала kill с задержкой
	go func() {
		time.Sleep(3 * time.Second)
		signalChannel <- syscall.SIGTERM
		fmt.Println("сигнал kill отправлен в канал")
	}()

	select {
	case <-ctx.Done():
		sig := <-signalChannel
		fmt.Printf("\ncaught the signal: %v\n", sig)
		fmt.Println("сработал контекст таймаут")
	case <-doneCh:
	}
}

// longJobFunction имитация долгой деятельности (счетчик с задержкой в секунду)
//
// позвращает error
func longJobFunction() error {
	fmt.Println("start longJobFunction")
	fmt.Println("waiting signal for kill process ", os.Getpid())
	for i := 1; i <= 20; i++ {
		fmt.Println(i)
		time.Sleep(1 * time.Second)
	}
	fmt.Println("finish longJobFunction")
	return nil
}
