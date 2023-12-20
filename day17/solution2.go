package main

import (
    "bufio"
    "fmt"
    "math"
    "os"
    "strconv"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

type Direction int;

const (
    UP Direction = iota
    LEFT
    DOWN
    RIGHT
)

type Tile struct {
    row int
    col int
    direction Direction
    streak int
}

func getLeft(dir Direction) Direction {
    return (dir + 1) % 4;
}

func getRight(dir Direction) Direction {
    return (dir + 3) % 4;
}

func min(a, b int) int {
    if a < b {
        return a;
    }
    return b;
}

func DP(grid [][]int) int {
    height := len(grid);
    width := len(grid[0]);
    var minHeatLoss int = math.MaxInt32;

    visited := make(map[Tile]int);

    queue := []Tile{Tile{0, 0, RIGHT, 0}, Tile{0, 0, DOWN, 0}};

    for len(queue) > 0 {
        var tile = queue[0];
        queue = queue[1:];

        if tile.row == height - 1 && tile.col == width - 1 && tile.streak >= 4 {
            minHeatLoss = min(minHeatLoss, visited[tile]);
        }

        for _, dir := range [3]Direction{tile.direction, getLeft(tile.direction), getRight(tile.direction)} {
            var row, col int = tile.row, tile.col;

            switch dir {
            case UP:
                row = tile.row - 1;
            case DOWN:
                row = tile.row + 1;
            case LEFT:
                col = tile.col - 1;
            case RIGHT:
                col = tile.col + 1;
            }
            if row < 0 || row >= height || col < 0 || col >= width {
                continue;
            }

            var nextStreak int = 1;
            if dir == tile.direction {
                nextStreak = tile.streak + 1;
            }

            if (tile.direction == dir && tile.streak < 10) || (tile.direction != dir && tile.streak >= 4) {
                nextTile := Tile{row, col, dir, nextStreak};
                totalHeatLoss := visited[tile] + grid[nextTile.row][nextTile.col];
                if val, found := visited[nextTile]; !found || val > totalHeatLoss {
                    visited[nextTile] = totalHeatLoss;
                    fmt.Printf("  Updating distance to %d, %d: %d direction=%v\n", nextTile.row, nextTile.col, totalHeatLoss, nextTile.direction);
                    queue = append(queue, nextTile);
                }
            }
        }
    }

    return minHeatLoss;
}

func main() {
    file, err := os.Open("input.txt");
    check(err);
    defer file.Close();

    var grid [][]int;
    scanner := bufio.NewScanner(file);
    for scanner.Scan() {
        var line = scanner.Text();
        var row []int;
        for i := 0; i < len(line); i++ {
            num, err := strconv.Atoi(string(line[i]));
            check(err);
            row = append(row, num);
        }
        grid = append(grid, row);
    }

    fmt.Printf("Grid is %d x %d\n", len(grid), len(grid[0]));
    fmt.Printf("Minimum path sum is %d\n", DP(grid));

}