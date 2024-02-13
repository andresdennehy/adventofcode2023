package main

import (
    "bufio"
    "fmt"
    "os"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

type Point struct {
	row, col int
}

type State struct {
	point Point
	steps int
	path []Point
}

func isValid(row, col, width, height int, visited map[Point]int, steps int) bool {
	if col >= 0 && col < width && row >= 0 && row < height {
        return true
	}
	return false
}

func isInPath(path []Point, point Point) bool {
    for _, p := range path {
        if p.row == point.row && p.col == point.col {
            return true;
        }
    }
    return false;
}

func longestHike(trails []string) int {
	height, width := len(trails), len(trails[0])

	var start, end Point;
	for i, c := range trails[0] {
	    if c == '.' {
            start = Point{0, i};
            break;
        }
	}
	for i, c := range trails[height-1] {
        if c == '.' {
            end = Point{height-1, i};
            break;
        }
    }
    fmt.Printf("Start: %v, End: %v\n", start, end)

	// Initialize visited array
	visited := make(map[Point]int)

	queue := []State{State{start, 0, []Point{}}}

	// Possible moves: up, down, left, right
	moves := []Point{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	maxDistance := 0;

	for len(queue) > 0 {
		currentState := queue[0];
		queue = queue[1:];

		row, col, steps := currentState.point.row, currentState.point.col, currentState.steps;

		if currentState.point.row == end.row && currentState.point.col == end.col {
			if steps > maxDistance {
                maxDistance = steps;
            }
            continue;
		}

		if isInPath(currentState.path, currentState.point) {
            continue;
        }

		if prev, ok := visited[currentState.point]; ok && prev > currentState.steps {
            continue;
        }

        visited[currentState.point] = currentState.steps;
        currPath := make([]Point, len(currentState.path))
        copy(currPath, currentState.path)
        currPath = append(currPath, currentState.point)

		switch trails[row][col] {
		case '>':
		    queue = append(queue, State{Point{row, col+1}, steps + 1, currPath})
		    continue;
        case '<':
            queue = append(queue, State{Point{row, col-1}, steps + 1, currPath})
            continue;
        case '^':
            queue = append(queue, State{Point{row-1, col}, steps + 1, currPath})
            continue;
        case 'v':
            queue = append(queue, State{Point{row+1, col}, steps + 1, currPath})
            continue;
        case '.':
            for _, move := range moves {
                newRow, newCol := row+move.row, col+move.col
                if isValid(newRow, newCol, width, height, visited, steps) && trails[newRow][newCol] != '#' {
                    queue = append(queue, State{Point{newRow, newCol}, steps + 1, currPath})
                }
            }
		}
	}

	return maxDistance // No valid path found
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var trails []string;
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		trails = append(trails, line)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	result := longestHike(trails)
	fmt.Printf("The longest hike is %d steps long.\n", result)
}