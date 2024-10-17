package utils

import (
	"fmt"
	"os"
)

func ParseFile(filename string) {
	filebytes, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("err reading file:", err)
	}
	filestring := string(filebytes)
	fmt.Println(filestring)
}
