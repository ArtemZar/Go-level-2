package panic_recover

import "fmt"

func NewPanic() {
	defer func() {
		if v := recover(); v != nil {
			fmt.Println("recovered", v)
		}
	}()
	panic("A-Х-Т-У-Н-Г!!!")
}