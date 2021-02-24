/**
1. Напишите программу, в которой неявно будет срабатывать паника.
Сделайте отложенную функцию, которая будет обрабатывать эту панику и печатать предупреждение в консоль.
Критерий выполнения задания — программа не завершается аварийно.
2. Дополните программу собственной ошибкой, хранящей время её возникновения.
3. Напишите функцию, которая создаёт файл в файловой системе
и использует отложенный вызов функций для безопасного закрытия файла.
**/

package main

import (
	"github.com/ArtemZar/Go-level-2/lesson-1/creatingFile"
	"os"

	"fmt"
	"github.com/ArtemZar/Go-level-2/lesson-1/myselfError"
	"github.com/ArtemZar/Go-level-2/lesson-1/panicRecover"
)

func main() {
	// задание1
	panicRecover.NewPanic()

	// задание2
	myselfError.StartMyselfError()

	// задание3
	if err := creatingFile.CreatingNewFile(); err != nil {
		fmt.Println("Error function:", err)
		os.Exit(1)
	}

}
