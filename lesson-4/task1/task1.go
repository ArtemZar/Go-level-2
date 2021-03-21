/*
С помощью пула воркеров написать программу,
которая запускает 1000 горутин, (конкурентный счетчик который будет считать до тысячи)
каждая из которых увеличивает число на 1.
Дождаться завершения всех горутин и убедиться,
что при каждом запуске программы итоговое число равно 1000.
*/

package main

import (
	"errors"
)

const  quantity = 1000
var controlSlice []int
var done = make(chan struct{})
var jobs = make(chan int)

func main() {

	// запускаем 1000 горутин (воркеров)
	for w := 1; w <= quantity; w++ {
		go worker(w, jobs)
	}

	// отправляем в канал исходное значение для работы воркеров
	jobs <- 0

	<-done
}

// worker функция для запуска горутин каждая из которых увеличивает полученное из канала jobs число на 1
//
// принимает порядковый номер id запускаемоей горутины (int) и канал для приема-передачи
//
// ничего не возвращает
func worker(id int, jobs chan int) {
	working := <-jobs
	//fmt.Println("worker", id, "input", working)
	working++
	controlSlice = append(controlSlice, working)
	//time.Sleep(time.Second)
	//fmt.Println("worker", id, "output", working)
	if working == quantity {
		close(jobs)
		close(done)
	} else {
		jobs <- working
	}
}

// Foo функция для тестирования
//
// не принимает аргументы. возвращает ошибку если максимальное значение счетчика не будет найдено
func Foo ()  error {
	main()
	for _, val := range controlSlice {
		if val == quantity {
			return  nil
		}
	}
	return errors.New("значение 1000 не найдено")
}




