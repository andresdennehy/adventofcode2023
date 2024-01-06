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

type Coordinate struct {
    row int
    col int
}

type QueueItem struct {
    point Coordinate
    steps int
}

func BFS(grid [][]byte, start Coordinate, steps int) int {
    // Advances the grid but tracks how many steps it did so far
    // If it did `steps` steps and it didn't reach that one in the same
    // number of steps, it sums one to the result
    var result int;

    // Create a queue
    var queue []QueueItem;
    queue = append(queue, QueueItem{start, 0});

    // Create a visited map
    visited := make(map[QueueItem]bool);

    for len(queue) > 0 {
        // Get the first element
        current := queue[0];
        queue = queue[1:];

        // Check if it's visited
        if _, ok := visited[current]; ok {
            continue;
        }

        // Mark it as visited
        visited[current] = true;

        // Check if it's the end
        if current.steps == steps {
            result++;
            continue;
        }

        // Add the neighbours to the queue
        if current.point.row > 0 && grid[current.point.row - 1][current.point.col] != '#' {
            queue = append(queue, QueueItem{Coordinate{current.point.row - 1, current.point.col}, current.steps + 1});
        }
        if current.point.row < len(grid) - 1 && grid[current.point.row + 1][current.point.col] != '#' {
            queue = append(queue, QueueItem{Coordinate{current.point.row + 1, current.point.col}, current.steps + 1});
        }
        if current.point.col > 0 && grid[current.point.row][current.point.col - 1] != '#' {
            queue = append(queue, QueueItem{Coordinate{current.point.row, current.point.col - 1}, current.steps + 1});
        }
        if current.point.col < len(grid[0]) - 1 && grid[current.point.row][current.point.col + 1] != '#'{
            queue = append(queue, QueueItem{Coordinate{current.point.row, current.point.col + 1}, current.steps + 1});
        }
    }

    return result;
}

func main() {
    file, err := os.Open("input.txt");
    check(err);
    defer file.Close();

    var grid [][]byte;
    scanner := bufio.NewScanner(file);
    for scanner.Scan() {
        line := scanner.Text();
        grid = append(grid, []byte(line));
    }

    var start Coordinate;
    // Find the starting point
    for i, row := range grid {
        for j, point := range row {
            if point == 'S' {
                start = Coordinate{i, j};
                break;
            }
        }
    }

    fmt.Printf("Result is: %v\n", BFS(grid, start, 64));

}