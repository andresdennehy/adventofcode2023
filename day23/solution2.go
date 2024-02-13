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

func isInList(list []Point, point Point) bool {
    for _, p := range list {
        if p == point {
            return true
        }
    }
    return false
}

func getNeighbors(trails []string, row, col int) []Point {
    height, width := len(trails), len(trails[0])
    neighbors := []Point{}

    if row > 0 && trails[row-1][col] != '#' {
        neighbors = append(neighbors, Point{row-1, col})
    }
    if row < height-1 && trails[row+1][col] != '#' {
        neighbors = append(neighbors, Point{row+1, col})
    }
    if col > 0 && trails[row][col-1] != '#' {
        neighbors = append(neighbors, Point{row, col-1})
    }
    if col < width-1 && trails[row][col+1] != '#' {
        neighbors = append(neighbors, Point{row, col+1})
    }

    return neighbors
}

func isValid(row, col, width, height int, visited map[Point]int, steps int) bool {
	if col >= 0 && col < width && row >= 0 && row < height {
        return true
	}
	return false
}

func convertToGraph(trails []string) map[Point]map[Point]int {
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

	// Initialize graph, to convert from unweighted to weighted graph
	vertices := make([]Point, 0)
	graph := make(map[Point]map[Point]int)

	for i := 0; i < height; i++ {
        for j := 0; j < width; j++ {
            if trails[i][j] != '#' {
                adjCount := len(getNeighbors(trails, i, j))
                if adjCount > 2 {
                    vertices = append(vertices, Point{i, j})
                }
            }
        }
    }
    vertices = append(vertices, start, end)
    fmt.Printf("Vertices: %v\n", vertices)

    // Create weighted graph
    for _, point := range vertices {
        queue := []Point{point}
        seen := make(map[Point]bool)
        seen[point] = true
        dist := 0
        for len(queue) > 0 {
            newQueue := make([]Point, 0)
            dist++
            for _, popped := range queue {
                for _, neighbor := range getNeighbors(trails, popped.row, popped.col) {
                    if _, ok := seen[neighbor]; !ok {
                        if isInList(vertices, neighbor) {
                            if _, ok := graph[point]; !ok {
                                graph[point] = make(map[Point]int)
                            }
                            graph[point][neighbor] = dist
                            seen[neighbor] = true
                        } else {
                            newQueue = append(newQueue, neighbor)
                            seen[neighbor] = true
                        }
                    }
                }
            }
            queue = newQueue
        }
    }
    fmt.Printf("Graph: %v\n", graph)

    return graph

}

func dfs(graph map[Point]map[Point]int, end Point, cur Point, pathSet map[Point]bool, totalDist int, result *int) {
	if cur == end {
		if totalDist > *result {
			*result = totalDist
		}
		return
	}
	for neighbor, weight := range graph[cur] {
		if !pathSet[neighbor] {
			pathSet[neighbor] = true
			dfs(graph, end, neighbor, pathSet, totalDist+weight, result)
			delete(pathSet, neighbor)
		}
	}
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var trails []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		trails = append(trails, line)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	graph := convertToGraph(trails)
    result := 0
    visited := make(map[Point]bool)
    dfs(graph, Point{len(trails)-1, len(trails[0])-2}, Point{0, 1}, visited, 0, &result)

	fmt.Printf("The longest hike is %d steps long.\n", result)
}