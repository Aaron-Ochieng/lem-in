package main

import (
	"fmt"

	"lem-in/utils"
)

func main() {
	filename, errmsg := utils.ParseArgs()
	if errmsg != "" {
		fmt.Println(errmsg)
		return
	}
	utils.ParseFile(filename)
}
