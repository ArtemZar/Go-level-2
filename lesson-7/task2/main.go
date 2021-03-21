/*
Написать функцию, которая принимает на вход имя файла и название функции.
Необходимо подсчитать в этой функции количество вызовов асинхронных функций.
Результат работы должен возвращать количество вызовов int и ошибку error.
Разрешается использовать только go/parser, go/ast и go/token.
*/

package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"sync"
	"time"
)

const (
	fileName = "main.go" // название файла, по которому мы итерируемся для составления AST
	funcName = "f"
)

func main() {

	f()

	result, err := myParser(fileName, funcName)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("количество вызовов асинхронных функций: %v\n", result)

}

// f функция запуска горутин
//
func f() {
	wg := sync.WaitGroup{}
	wg.Add(3)

	go func() {
		defer wg.Done()
		time.Sleep(1 * time.Second)
	}()

	go func() {
		defer wg.Done()
		time.Sleep(1 * time.Second)
	}()

	go func() {
		defer wg.Done()
		time.Sleep(1 * time.Second)
	}()

	wg.Wait()
}

// myParser парсит файл, находит заданную функцию и подсчитывает количество запусков асинхронных функций
//
func myParser(fileName, funcName string) (int, error) {
var counter int
	fset := token.NewFileSet()
	astFile, err := parser.ParseFile(fset, fileName, nil, parser.ParseComments) // режим parser.ParseComments анализирует все, включая комментарии
	if err != nil {
		return 0, err
	}

	for _, f := range astFile.Decls {
		fn, ok := f.(*ast.FuncDecl)
		if !ok {
			continue
		}

		if fn.Name.Name == funcName {
			fmt.Printf("Function %v found\n", funcName)
			ast.Inspect(fn, func(n ast.Node) bool {
				// Find go funcs
				//
				ret, ok := n.(*ast.GoStmt)
				if ok {
					fmt.Printf("\tGO statement found on line %d\n", fset.Position(ret.Pos()).Line)
					counter++
					return true
				}
				return true
			})
		}

	}


	return counter, err
}

