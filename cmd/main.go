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
	shortest := utils.OptimizePaths(paths)
	fmt.Println(shortest)
	shortopt := utils.Optimize2(paths, Antcolony)
	fmt.Println(shortopt)
	firstop := utils.PlaceAnts(Antcolony, shortest)
	secondop := utils.PlaceAnts(Antcolony, shortopt)
	fmt.Println(firstop)
	fmt.Println(secondop)

	turns1 := utils.GenerateTurns(firstop,shortest)
	turns2 := utils.GenerateTurns(secondop, shortopt)
    fmt.Println(turns1)
	fmt.Println(turns2)
	finalpath := shortest
	finalAntspalced := firstop
	turns := turns1
	if turns1 > turns2 {
		finalpath = shortopt
		finalAntspalced = secondop
		turns = turns2
	}
	fmt.Println(finalpath)

	moves := utils.MoveAnts(finalpath, finalAntspalced,turns)
	fmt.Println(moves)
}
