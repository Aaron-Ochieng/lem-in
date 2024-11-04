package utils

import (
	"lem-in/models"
)

// FindPaths finds all possible paths from start to end using BFS
func FindPaths(colony *models.AntColony) ([]models.Path,map[int][]int, int) {
	var allPaths []models.Path
	var queue [][]string
	queue = append(queue, []string{colony.Start})
	for i := range colony.Rooms {
		colony.Rooms[i].IsVisited = false
	}
	for len(queue) > 0 {
		currentPath := queue[0]
		queue = queue[1:]
		currentRoom := currentPath[len(currentPath)-1]
		if currentRoom == colony.End {
			allPaths = append(allPaths, models.Path{Rooms: currentPath})
			continue
		}
		for _, nextRoom := range colony.Links[currentRoom] {
			if !containsRoom(currentPath, nextRoom) {
				// Create a new path with the next room added
				newPath := make([]string, len(currentPath))
				copy(newPath, currentPath)
				newPath = append(newPath, nextRoom)
				queue = append(queue, newPath)
			}
		}
	}
	return ChooseOptimumPath(allPaths, colony)
}

// Helper function to check if a room is in a path
func containsRoom(path []string, room string) bool {
	for _, r := range path {
		if r == room {
			return true
		}
	}
	return false
}

func ChooseOptimumPath(paths []models.Path, Antcolony *models.AntColony) ([]models.Path, map[int][]int,int) {
	shortest := OptimizePaths(paths)
	shortopt := Optimize2(paths, Antcolony)
	firstop := PlaceAnts(Antcolony, shortest)
	secondop := PlaceAnts(Antcolony, shortopt)
	turns1 := GenerateTurns(firstop,shortest)
	turns2 := GenerateTurns(secondop, shortopt)
	finalpath := shortest
	finalAntspalced := firstop
	turns := turns1
	if turns1 > turns2 {
		finalpath = shortopt
		finalAntspalced = secondop
		turns = turns2
	}
	return finalpath, finalAntspalced,turns
}

func OptimizePaths(paths []models.Path) []models.Path {
	optimized := make([]models.Path, 0)
	optimized = append(optimized, paths[0])
	for i := 1; i < len(paths); i++ {
		if Check(paths[i].Rooms, optimized) {
			optimized = append(optimized, paths[i])
		}
	}
	return optimized
}

func Check(path []string, optimized []models.Path) bool {
	for _, optpath := range optimized {
		for k := 1; k < len(path)-1; k++ {
			for j :=1; j < len(optpath.Rooms)-1; j++{
				if path[k] == optpath.Rooms[j] {
					return false
				}
			}
		}
	}
	return true
}

func Optimize2(paths []models.Path, colony *models.AntColony) []models.Path {
	half := colony.NumberOfAnts / 2
	optimized := make([]models.Path, 0)
	optimized = append(optimized, paths[0])
	for i := 1; i < len(paths); i++ {
		if len(paths[i].Rooms)-1 <= half {
			v, index := Check2(paths[i].Rooms, optimized)
			if !v {
				if len(optimized[index].Rooms) != len(paths[i].Rooms) {
					optimized = Remove(optimized, index)
					optimized = append(optimized, paths[i])
				}
			} else {
				optimized = append(optimized, paths[i])
			}
		}
	}
	return optimized
}

func Check2(path []string, optimized []models.Path) (bool, int) {
	for g, optpath := range optimized {
		for k := 1; k < len(path)-1; k++ {
			for j :=1; j < len(optpath.Rooms)-1; j++{
				if path[k] == optpath.Rooms[j] {
					return false, g
				}
			}
		}
	}
	return true, 0
}

func Remove(optimized []models.Path, index int) []models.Path {
	new := make([]models.Path, 0)
	for i, path := range optimized {
		if i != index {
			new = append(new, path)
		}
	}
	return new
}

func GenerateTurns( option map[int][]int, paths []models.Path) int {
	maxturns := 0
	for i, path := range paths {
		rooms := len(path.Rooms)-1
		ants := len(option[i])
		turns := rooms + ants -1
		if turns > maxturns {
            maxturns = turns
        }
	}
	return maxturns
}

func PlaceAnts(colony *models.AntColony, paths []models.Path) map[int][]int {
	ants := colony.NumberOfAnts
	pathants := make(map[int][]int)
	ant := 1
	for ants > 0 {
		if PlaceRecursively(ant, paths, pathants, 0) {
			ant++
			ants--
		}
	}
	return pathants
}

func PlaceRecursively(ant int, paths []models.Path, pathants map[int][]int, path int) bool {
	if ant == 1 || path == len(paths)-1 {
		pathants[path] = append(pathants[path], ant)
		return true
	}
	rooms := len(paths[path].Rooms)-2
	antsinroom := len(pathants[path])-2
	det := rooms + antsinroom
	if det > (len(paths[path+1].Rooms)-2)+(len(pathants[path+1])-2){
		if PlaceRecursively(ant, paths, pathants, path+1) {
			return true
		}
	} else {
		pathants[path] = append(pathants[path], ant)
		return true
	}
	return false
}
