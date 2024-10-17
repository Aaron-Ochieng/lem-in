package utils

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ParseFile(filename string) {
	filebytes, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("err reading file:", err)
	}
	filestring := string(filebytes)
	fileslice := strings.Split(filestring, "\n")

	numberOfAnts := 0
	roomMap := make(map[string][2]int)
	linkMap := make(map[string]string)

	for i, v := range fileslice {
		if i == 0 {
			n, err := strconv.Atoi(v)
			if err != nil {
				fmt.Println(err)
				return
			}
			numberOfAnts = n
		} else {
			if strings.HasPrefix(v, "#") {
				continue
			} else {
				if strings.Contains(v, " ") {
					x_coord, err := strconv.Atoi(strings.Fields(v)[1])
					if err != nil {
						fmt.Println(err)
					}
					y_coord, err := strconv.Atoi(strings.Fields(v)[2])
					if err != nil {
						fmt.Println(err)
					}
					arr := [2]int{x_coord, y_coord}
					roomMap[strings.Fields(v)[0]] = arr
				}
				if strings.Contains(v, "-") {
					linkMap[strings.Split(v, "-")[0]] = strings.Split(v, "-")[1]
				}
			}
		}
	}
	fmt.Println(numberOfAnts)
	fmt.Println(roomMap)
	fmt.Println(linkMap)
}
