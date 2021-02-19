//1. Напишите программу, в которой неявно будет срабатывать паника.
//Сделайте отложенную функцию, которая будет обрабатывать эту панику и печатать предупреждение в консоль.
//Критерий выполнения задания — программа не завершается аварийно.
//2. Дополните программу собственной ошибкой, хранящей время её возникновения.
//3. Напишите функцию, которая создаёт файл в файловой системе
//и использует отложенный вызов функций для безопасного закрытия файла.



package main

import (
	"github.com/ArtemZar/Go-level-2/lesson-1/creating_file"
	"github.com/ArtemZar/Go-level-2/lesson-1/myself_error"
	"github.com/ArtemZar/Go-level-2/lesson-1/panic_recover"
)

func main()  {
	//задание1
	panic_recover.NewPanic()

	//задание2
	myself_error.StartMyselfError()

	//задание3
	creating_file.CreatingNewFile()
}