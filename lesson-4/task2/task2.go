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
	ctx := context.Background()
	var cancel context.CancelFunc

	longJobFunctionDoneCh := make(chan error)
	go func(ctx context.Context, cancel context.CancelFunc) {
		longJobFunctionDoneCh <- longJobFunction()
	}(ctx, cancel)

	// имитация сигнала  kill
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, syscall.SIGTERM)
	time.Sleep(5 * time.Second)
	signalChannel <- syscall.SIGTERM
	
	<-signalChannel
	ctx, cancel = context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	select {
	case <-ctx.Done():
		fmt.Println("сработал контекст таймаут")
	case <-longJobFunctionDoneCh:
	}
}

// longJobFunction функция с длительным временем работы
//
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