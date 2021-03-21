package main

import (
	"testing"
)

//TestFoo поиск максимального значения работы конкурентных счетчиков (1000)
//
func TestFoo(t *testing.T) {
	got := Foo()
	if got != nil {
		t.Errorf("Тест не пройден. Результат работы функции Foo: %v; должно быть nil", got)
	}
}