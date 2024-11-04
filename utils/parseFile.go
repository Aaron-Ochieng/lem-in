package utils

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"lem-in/models"
)

func fileContents(filename string) (res []string, err error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		text := scanner.Text()
		if text != "" && (!strings.Contains(text, "#") || strings.Contains(text, "##end") || strings.Contains(text, "##start")) {
			res = append(res, text)
		}
	}
	return
}
func validateRoomCoordinates(s []string) bool {
	if len(s) != 3 {
		return false
	}

	_, err := strconv.Atoi(s[1])

	if err != nil {
		return false
	}

	if _, err = strconv.Atoi(s[2]); err != nil {
		return false
	}

	return true
}

func roomConnection(s string) bool {
	val := strings.Split(s, "-")
	return len(val) == 2 && val[0] != val[1]
}

func splitRoomCordinates(s string) (string, bool) {
	val := strings.Split(s, "")
	return val[0], validateRoomCoordinates(val)
}

func ParseFile(filename string) (colony *models.AntColony, err error) {
	contents, err := fileContents(filename)
	if err != nil {
		return
	}

	// Number of ants
	val, err := strconv.Atoi(contents[0])

	if err != nil {
		err = errors.New("invalid number of ants")
		return
	}

	if val == 0 {
		err = errors.New("number of ants cannot be 0")
	}

	for i := 1; i < len(contents); i++ {
		// Locate start room
		if strings.Contains(contents[i], "##start") {
			roomname, bl := splitRoomCordinates(contents[i+1])
			if !bl {
				err = errors.New("invalid room coordinates")
				return
			}
			colony.Start = roomname
		}

		// Locate end room
		if strings.Contains(contents[i], "##end") {
			roomname, bl := splitRoomCordinates(contents[i+1])
			if !bl {
				err = errors.New("invalid room coordinates")
				return
			}
			colony.End = roomname
		}

		// Room with coordinates
		if strings.Contains(contents[i], " ") {
			temp := strings.Split(contents[i], " ")
			if !validateRoomCoordinates(temp) {
				err = errors.New("invalid room coordinates")
				return
			}
			// check if room already exists
			_, ok := colony.Links[temp[0]]
			if ok {
				err = errors.New("room already exist")
				return
			}
			// append it to a map
			colony.Links[temp[0]] = []string{}
		}

	}

	return
}
