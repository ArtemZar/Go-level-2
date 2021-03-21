/*
Написать функцию, которая принимает на вход
структуру in (struct или кастомную struct) и
values map[string]interface{}
(key - название поля структуры, которому нужно присвоить value этой мапы).
Необходимо по значениям из мапы изменить входящую структуру in с помощью пакета reflect.
Функция может возвращать только ошибку error.
Написать к данной функции тесты
*/

package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name string
	Age int
}

func main() {
	NewUser := User{
		Name:"Artem",
		Age: 34,
	}

	values := map[string]interface{}{
		"Name": "Ivan",
		"Age": 19,
	}


f(&NewUser, values)
	fmt.Println(NewUser)

}

func f(in interface{}, values map[string]interface{}) {
	if in==nil{ // проверка не пустая ли структура
		return
	}

	v := reflect.ValueOf(in)

	if v.Kind() == reflect.Ptr { // проверка указатель или значение
		v = v.Elem() // разиминовывание если указатель
	}

	if  !v.CanSet() { // проверка чтобы это значение было устанавливаемым
		return
	}

	if v.Kind() != reflect.Struct{ // проверка структура или нет
		return
	}

	for i := 0; i < v.NumField(); i++ { // итерация полей структуры
		typeFild := v.Field(i)
		newName, ok := values[v.Type().Field(i).Name] // проверка на соответсвие полей структуры ключам мапы
		if ok {
			func(any, newVel interface{}) { // функция для реализации type switch
				switch any.(type) {
				case int:
					typeFild.SetInt(reflect.ValueOf(newVel).Int())
				case string:
					typeFild.SetString(reflect.ValueOf(newVel).String())
				default:
					fmt.Println("???")
				}
			}(typeFild.Interface(), newName)
		}
	}
}