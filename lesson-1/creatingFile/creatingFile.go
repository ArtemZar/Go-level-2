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

	defer newFile.Close()
	_, _ = fmt.Fprintln(newFile, "data")

	if err := newFile.Close(); err != nil {
		if err != nil {
			return err
		}
	}

	return err
}
