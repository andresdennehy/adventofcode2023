package main

import (
    "bufio"
	"fmt"
	"os"
	"strings"
)

type Point struct {
	row, col int
}

type QueueItem struct {
    point       Point
    stepsTaken  int
}

func solve(raw string, steps int, start Point) int {
	grid := make(map[Point]byte)
	rows := strings.Split(raw, "\n")
	for rn, row := range rows {
		for cn, char := range row {
			grid[Point{rn, cn}] = byte(char)
		}
	}

    if grid[start] != 'S' {
        panic("Start is a wall")
    }

	seen := make(map[Point]int)
	parity := steps % 2
	fmt.Printf("Parity: %d\n", parity)
	queue := []QueueItem{{point: start, stepsTaken: 0}}

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		if seen[curr.point] > 0 {
			continue
		}
		seen[curr.point] = curr.stepsTaken

		if curr.stepsTaken > steps {
			continue
		}

		for _, neighbor := range getNeighbours(curr.point) {
			if grid[neighbor] == '#' {
				continue
			}
			if neighbor.row < 0 || neighbor.col < 0 {
                continue
            }
            if neighbor.row >= len(rows) || neighbor.col >= len(rows[0]) {
                continue
            }
			queue = append(queue, QueueItem{point: neighbor, stepsTaken: curr.stepsTaken + 1})
		}
	}

	count := 0
	for _, stepsTaken := range seen {
		if stepsTaken % 2 == parity {
			count++
		}
	}
	return count
}

func BFS(raw string, steps int) int {
	size := len(strings.Split(raw, "\n")[0])

	var rows []string
	for _, row := range strings.Split(raw, "\n") {
		rows = append(rows, strings.Repeat(row, 5))
	}
	blockOfRows := strings.Join(rows, "\n")

	// Repeat the entire block of rows five times
	var bigRows []string
	for i := 0; i < 5; i++ {
		bigRows = append(bigRows, blockOfRows)
	}
	bigRaw := strings.Join(bigRows, "")
	fmt.Printf("Big raw:\n%s\n", bigRaw)

	STEPS_A := size / 2
	STEPS_B := size + STEPS_A
	STEPS_C := size + STEPS_B

	fmt.Printf("STEPS_A: %d\n", STEPS_A)
	fmt.Printf("STEPS_B: %d\n", STEPS_B)
	fmt.Printf("STEPS_C: %d\n", STEPS_C)

	bigStart := Point{STEPS_C, STEPS_C}

	s1 := solve(bigRaw, STEPS_A, bigStart)
	s2 := solve(bigRaw, STEPS_B, bigStart)
	s3 := solve(bigRaw, STEPS_C, bigStart)

	N := (steps - STEPS_A) / size
	D1 := s2 - s1
	D2 := s3 - s2
	D3 := D2 - D1

	return s1 + N*D1 + N*(N-1)/2*D3
}

func getNeighbours(p Point) []Point {
	return []Point{
		{p.row - 1, p.col},
		{p.row + 1, p.col},
		{p.row, p.col - 1},
		{p.row, p.col + 1},
	}
}


func main() {

    file, error := os.Open("input.txt")
    if error != nil {
        panic(error)
    }
    defer file.Close()

	scanner := bufio.NewScanner(file)
	var data string;
	for scanner.Scan() {
        data += scanner.Text() + "\n"
    }

    fmt.Printf("Data:\n%s\n", data)

    fmt.Printf("Result: %d\n", BFS(data, 26501365))

}
