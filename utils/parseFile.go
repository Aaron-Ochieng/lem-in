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

// ParseFile reads and validates an ant colony configuration file
func ParseFile(filename string) (*models.AntColony, error) {
	contents, err := fileContents(filename)
	if err != nil {
		return nil, err
	}

	if len(contents) == 0 {
		return nil, errors.New("empty file")
	}

	colony := &models.AntColony{
		Rooms: make([]models.Room, 0),
		Links: make(map[string][]string),
	}

	// Parse number of ants
	antCount, err := strconv.Atoi(contents[0])
	if err != nil {
		return nil, errors.New("invalid number of ants")
	}
	if antCount <= 0 {
		return nil, errors.New("number of ants must be positive")
	}
	colony.NumberOfAnts = antCount

	// Parse rooms and connections
	for i := 1; i < len(contents); i++ {
		line := contents[i]

		switch {
		case strings.Trim(line, " ") == "##start":
			if i+1 >= len(contents) {
				return nil, errors.New("missing start room definition")
			}
			roomName, ok := parseRoom(contents[i+1], colony)
			if !ok {
				return nil, errors.New("invalid start room coordinates")
			}
			if _, exists := colony.Links[roomName]; exists {
				return nil, fmt.Errorf("duplicate room name: %s", roomName)
			}
			colony.Links[roomName] = []string{}
			colony.Start = roomName
			i++ // Skip the next line since we processed it

		case strings.HasPrefix(line, "##end"):
			if i+1 >= len(contents) {
				return nil, errors.New("missing end room definition")
			}
			roomName, ok := parseRoom(contents[i+1], colony)
			if !ok {
				return nil, errors.New("invalid end room coordinates")
			}
			if _, exists := colony.Links[roomName]; exists {
				return nil, fmt.Errorf("duplicate room name: %s", roomName)
			}
			colony.Links[roomName] = []string{}
			colony.End = roomName
			i++ // Skip the next line since we processed it

		case strings.Contains(line, " "):
			roomName, ok := parseRoom(line, colony)
			if !ok {
				return nil, errors.New("invalid room coordinates")
			}
			if _, exists := colony.Links[roomName]; exists {
				return nil, fmt.Errorf("duplicate room name: %s", roomName)
			}
			colony.Links[roomName] = []string{}

		case strings.Contains(line, "-"):
			if err := parseConnection(line, colony); err != nil {
				return nil, err
			}
		}
	}

	// Validate colony configuration
	if err := validateColony(colony); err != nil {
		return nil, err
	}

	return colony, nil
}

// fileContents reads non-empty and non-comment lines from a file
func fileContents(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		if text != "" && (!strings.Contains(text, "#") || strings.Contains(text, "##end") || strings.Contains(text, "##start")) {
			lines = append(lines, text)
			models.FileContents += text + "\n"
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}

	return lines, nil
}

// parseRoom parses a room definition line and adds it to the colony
func parseRoom(line string, colony *models.AntColony) (string, bool) {
	parts := strings.Split(line, " ")
	if len(parts) != 3 {
		return "", false
	}

	x, err := strconv.Atoi(parts[1])
	if err != nil {
		return "", false
	}

	y, err := strconv.Atoi(parts[2])
	if err != nil {
		return "", false
	}

	// Check for duplicate coordinates
	for _, room := range colony.Rooms {
		if room.Coord_X == x && room.Coord_Y == y {
			return "", false
		}
	}

	room := models.Room{
		Name:    parts[0],
		Coord_X: x,
		Coord_Y: y,
	}
	colony.Rooms = append(colony.Rooms, room)
	return room.Name, true
}

// parseConnection parses a room connection line and adds it to the colony
func parseConnection(line string, colony *models.AntColony) error {
	parts := strings.Split(line, "-")
	if len(parts) != 2 || parts[0] == parts[1] {
		return errors.New("invalid room connection")
	}

	// Verify both rooms exist
	if _, exists := colony.Links[parts[0]]; !exists {
		return fmt.Errorf("room does not exist: %s", parts[0])
	}
	if _, exists := colony.Links[parts[1]]; !exists {
		return fmt.Errorf("room does not exist: %s", parts[1])
	}

	// Add bidirectional connection
	colony.Links[parts[0]] = append(colony.Links[parts[0]], parts[1])
	colony.Links[parts[1]] = append(colony.Links[parts[1]], parts[0])
	return nil
}

// validateColony performs final validation of the colony configuration
func validateColony(colony *models.AntColony) error {
	if colony.Start == "" {
		return errors.New("no colony starting point defined")
	}
	if colony.End == "" {
		return errors.New("no colony ending point defined")
	}

	// Verify start and end rooms exist in links
	if _, exists := colony.Links[colony.Start]; !exists {
		return errors.New("start room not found in connections")
	}
	if _, exists := colony.Links[colony.End]; !exists {
		return errors.New("end room not found in connections")
	}

	return nil
}
