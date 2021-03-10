package creatingFile

import (
	"fmt"
	"os"
)

func CreatingNewFile() error {
	newFile, err := os.Create("newFile.txt")
	if err != nil {
		return err
	}

	defer func() {
		err := newFile.Close()
		if err != nil {
			fmt.Println("File is not close")
		}
	}()

	_, _ = fmt.Fprintln(newFile, "data")

	return err
}
