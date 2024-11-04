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
	shortopt := utils.Optimize2(paths, Antcolony)
	firstop := utils.PlaceAnts(Antcolony, shortest)
	secondop := utils.PlaceAnts(Antcolony, shortopt)
	turns1 := utils.GenerateTurns(firstop,shortest)
	turns2 := utils.GenerateTurns(secondop, shortopt)
	finalpath := shortest
	finalAntspalced := firstop
	turns := turns1
	if turns1 > turns2 {
		finalpath = shortopt
		finalAntspalced = secondop
		turns = turns2
	}
	moves := utils.MoveAnts(finalpath, finalAntspalced,turns)
	for _, move := range moves {
		fmt.Println(move)
	}
}
