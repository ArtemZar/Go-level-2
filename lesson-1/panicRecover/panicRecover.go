package panicRecover

import "fmt"

func NewPanic() {

	names := []string{
		"Artem",
		"Alexei",
		"Alexandr",
	}

	defer func() {
		if v := recover(); v != nil {
			fmt.Println("recovered panic: ", v)
		}
	}()

	fmt.Println("My favorite name is:", names[len(names)])

}
