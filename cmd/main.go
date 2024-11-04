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
	paths, antsperpath, turns:= utils.FindPaths(Antcolony)
	moves := utils.MoveAnts(paths, antsperpath, turns)

	for _, move := range moves {
		fmt.Println(move)
	} 
}
