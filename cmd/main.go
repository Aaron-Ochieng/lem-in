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
	Antcolony, err := utils.ParseFile(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	paths := utils.FindPaths(Antcolony)
	fmt.Println(paths)
	ants := utils.PlaceAnts(Antcolony, paths)
	fmt.Println(ants)
}
