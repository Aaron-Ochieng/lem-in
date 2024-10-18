package models

type AntColony struct {
	NumberOfAnts int
	Rooms        map[string][2]int
	Links        map[string][]string
	Start        string
	End          string
}
