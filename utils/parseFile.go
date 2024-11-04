package utils

import (
	"bufio"
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

func ParseFile(filename string) (colony *models.AntColony, err error) {
	contents, err := fileContents(filename)
	if err != nil {
		return
	}
	return
}
