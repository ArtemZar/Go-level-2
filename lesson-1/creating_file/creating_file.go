package creating_file

import (
	"os"
	"fmt"
)

func CreatingNewFile()  {

new_file, err := os.Create("new_file.txt")
if err != nil {
panic(err)
}
defer new_file.Close()
_, _ = fmt.Fprintln(new_file, "data")
}


