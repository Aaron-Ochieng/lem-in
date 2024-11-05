# Lem-in: Digital Ant Farm Simulator

`lem-in` is a Go-based program that simulates an ant farm where ants navigate a colony of interconnected rooms (nodes) and tunnels (edges) to find the quickest path from a starting room to an exit. The program reads colony specifications from a file, finds the shortest path(s), and outputs each move of the ants as they progress from room to room.

## Project Overview

The goal of this project is to develop a program that:
1. Parses a text file describing a colony's layout, rooms, and tunnels.
2. Simulates the movement of ants from a `##start` room to a `##end` room.
3. Determines the quickest path(s) while avoiding traffic jams and idle time.
4. Outputs each move until all ants have successfully exited the colony.

## How It Works

1. **Ant Colony Layout**:
   - The colony consists of rooms connected by tunnels.
   - The program reads a list of rooms (with coordinates) and tunnels from the input file.
   - Each room is represented by its name and coordinates in the format `"name x y"` (e.g., `Room1 2 3`).
   - Tunnels connect pairs of rooms in the format `"room1-room2"` (e.g., `1-2`).

2. **Ant Simulation**:
   - All ants start in the `##start` room.
   - The goal is to bring them to the `##end` room with as few moves as possible.
   - Paths are calculated to avoid traffic jams and make optimal use of available routes.

3. **Output**:
   - Displays the initial setup of ants, rooms, and tunnels.
   - Each line shows ants’ movements from one room to another in the format `Lx-y` where `x` is the ant number and `y` is the room name (e.g., `L1-2 L2-3`).

## Input File Format

The input file should contain:
- Number of ants
- Room specifications, including `##start` and `##end`.
- Tunnel specifications connecting the rooms.

### Example Input File

```plaintext
3
##start
0 1 0
##end
1 5 0
2 9 0
3 13 0
0-2
2-3
3-1
```

### Example Output

```plaintext
L1-2
L1-3 L2-2
L1-1 L2-3 L3-2
L2-1 L3-3
L3-1
```

## Features

- **Dynamic Pathfinding**: Calculates the shortest path for ants to exit the colony efficiently.
- **Traffic Management**: Ensures ants avoid congestion by using alternate routes when possible.
- **Error Handling**: Validates input for missing rooms, incorrect formatting, circular paths, self-loops, and more, displaying error messages if the format is invalid.

## Instructions

### Running the Program

```sh
$ go run . <filename>
```

Replace `<filename>` with the path to the file containing the colony’s data.

### Room and Tunnel Requirements

1. Room names cannot start with `L` or `#`.
2. Room names must not contain spaces.
3. Each tunnel connects exactly two rooms, and there can be no duplicate tunnels.
4. The `##start` and `##end` rooms can contain multiple ants, but each intermediate room can contain only one ant at a time.

## Example Usage

### Example 1

Given the file `test0.txt` with content:

```plaintext
3
##start
0 1 0
##end
1 5 0
2 9 0
3 13 0
0-2
2-3
3-1
```

Running the command:

```sh
$ go run . test0.txt
```

Produces the output:

```plaintext
L1-2
L1-3 L2-2
L1-1 L2-3 L3-2
L2-1 L3-3
L3-1
```

### Example 2

File `test1.txt`:

```plaintext
3
2 5 0
##start
0 1 2
##end
1 9 2
3 5 4
0-2
0-3
2-1
3-1
2-3
```

```sh
$ go run . test1.txt
```

Output:

```plaintext
L1-2 L2-3
L1-1 L2-1 L3-2
L3-1
```

## Error Handling

Common errors include:
- **Invalid Room Format**: Rooms missing coordinates or with invalid names.
- **Missing `##start` or `##end`**: Start and end rooms are mandatory.
- **Circular Paths or Self-Loops**: Rooms linking to themselves, causing infinite loops.
- **Improper Ant Count**: Negative or missing ant count.

For such cases, the program will output an error message:

```plaintext
ERROR: invalid data format
```

## Code Structure

- **Input Parsing**: Reads and validates file data, extracting ants, rooms, and tunnels.
- **Pathfinding**: Calculates optimal paths for ants to traverse the colony.
- **Output Generator**: Prints each move in the format specified.

## Development Notes

- **Language**: Go
- **Testing**: Recommended to write unit tests for each module to ensure correct functionality.
- **Allowed Packages**: Only Go’s standard library is permitted.

## Contributing

Contributions to this project are welcome. Please follow these steps to contribute:
1. Fork the repository.
2. Create a new feature branch.
3. Commit your changes with clear messages.
4. Submit a pull request for review.

## License

This project is licensed under the MIT License. See the LICENSE file for details.

## Authors

[Aaron Ochieng](https://learn.zone01kisumu.ke/git/aaochieng)
[Moses Onyango](https://learn.zone01kisumu.ke/git/moonyango)
[Andrew Osindo](https://learn.zone01kisumu.ke/git/aosindo)
