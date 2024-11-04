package utils

import (
	"fmt"
	"lem-in/models"
)


func MoveAnts(paths []models.Path, antsperroom map[int][]int, turns int) []string {
	moves := make([]string, turns)
	for i, p := range antsperroom {
		path := paths[i].Rooms
		for j , ant := range p {
			for k, room := range path[1:] {
				moves[j+k] += fmt.Sprintf("L%v-%v ",ant,room)
			}
		} 
	}
	return moves
}
