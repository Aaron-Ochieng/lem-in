package utils

import "lem-in/models"


func MoveAnts(paths []models.Path, antsperroom map[int][]int, turns int) []string {
	moves := make([]string, turns)
	count := 0
	for i, p := range antsperroom {
		path := paths[i].Rooms
		for j , ant := range p {
			count++
			for k, room := range path {
				count-1
				moves[co]
			}
		} 
	}
	return moves
}