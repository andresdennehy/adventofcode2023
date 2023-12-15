package main

import (
    "bufio"
    "fmt"
    "os"
    //"strconv"
    //"strings"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

type Coordinate struct {
    row int;
    col int;
}

func printGrid(grid [][]byte) {
    for i:=0; i<len(grid); i++ {
        for j:=0; j<len(grid[0]); j++ {
            fmt.Printf("%c", grid[i][j]);
        }
        fmt.Printf("\n");
    }
}

func tiltNorth(grid [][]byte) {
    // Slices are mutable, so we can just swap the values
    for i:=0; i<len(grid); i++ {
        for j:=0; j<len(grid[0]); j++ {
            if grid[i][j] == 'O' {
                currPos := Coordinate{i, j};
                // Swap with one above until it can't move anymore
                for currPos.row > 0 && grid[currPos.row-1][currPos.col] == '.' {
                    grid[currPos.row][currPos.col] = '.';
                    grid[currPos.row-1][currPos.col] = 'O';
                    currPos.row--;
                }
            }
        }
    }
}

func calculateScore(grid [][]byte) int {
    var result int;
    var southEdge int = len(grid);
    for i:=0; i<len(grid); i++ {
        for j:=0; j<len(grid[0]); j++ {
            if grid[i][j] == 'O' {
                result += southEdge - i;
            }
        }
    }
    return result;
}

func solve(scanner *bufio.Scanner) int {
    var grid [][]byte;

    for (*scanner).Scan() {
        line := (*scanner).Text();
        row := make([]byte, len(line));
        for i := 0; i < len(line); i++ {
            row[i] = line[i];
        }
        grid = append(grid, row);
    }

    printGrid(grid);
    fmt.Printf("\nTilting north\n")
    tiltNorth(grid);

    return calculateScore(grid);
}

func main() {

    file, err := os.Open("input.txt");
    check(err);
    defer file.Close();

    scanner := bufio.NewScanner(file);
    fmt.Printf("Result is %d\n", solve(scanner));

}