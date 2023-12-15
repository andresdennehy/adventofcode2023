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

func printGrid(grid [][]byte) {
    for i:=0; i<len(grid); i++ {
        for j:=0; j<len(grid[0]); j++ {
            fmt.Printf("%c", grid[i][j]);
        }
        fmt.Printf("\n");
    }
}

func findMirrorLine(grid [][]byte) int {
    // Brute force approach
    // In a nxm grid, O(n^2*m + m^2*n)
    // For each possible column, check all columns (m^2) and check if they are mirrored (n)
    // For each possible row, check all rows (n^2) and check if they are mirrored (m)
    var matches bool;
    width := len(grid[0]);
    height := len(grid);

    // Vertical case
    // Assume mirror line is to the left of col
    for col:=1; col<width; col++ {
        matches = true;
        smudges := 0;
        checkColLoop:
        for j:=col-1; j>=0 && col+(col-j-1) < width; j-- {
            for row:=0; row<height; row++ {
                if grid[row][j] != grid[row][col+(col-j-1)] {
                    if smudges > 0 {
                        matches = false;
                        break checkColLoop;
                    }
                    smudges++;
                }
            }
        }
        if matches && smudges == 1 {
            return col;
        }
    }

    // Horizontal case
    // Assume mirror line is above row
    for row:=1; row<height; row++ {
        matches = true;
        smudges := 0;
        checkRowLoop:
        for i:=row-1; i>=0 && row+(row-i-1) < height; i-- {
            for col:=0; col<width; col++ {
                if grid[i][col] != grid[row+(row-i-1)][col] {
                    if smudges > 0 {
                        matches = false;
                        break checkRowLoop;
                    }
                    smudges++;
                }
            }
        }
        if matches && smudges == 1 {
            return 100*row;
        }
    }

    return 0;
}

func solve(scanner *bufio.Scanner) int {
    var result int;
    var grid [][]byte;
    var i int;

    for (*scanner).Scan() {
        line := (*scanner).Text();

        // If line is empty, we have reached the end of the grid
        if len(line) == 0 {
            localResult := findMirrorLine(grid);
            fmt.Printf("Grid %d: %d\n", i, localResult);
            result += localResult;
            grid = [][]byte{};
            i++;
            continue;
        }

        gridRow := make([]byte, len(line));
        for i:=0; i<len(line); i++ {
            gridRow[i] = line[i];
        }
        grid = append(grid, gridRow);
    }

    fmt.Printf("Grid %d: %d\n", i+1, result);
    result += findMirrorLine(grid);

    return result;
}

func main() {

    file, err := os.Open("input.txt");
    check(err);
    defer file.Close();

    scanner := bufio.NewScanner(file);
    fmt.Printf("Result is %d\n", solve(scanner));

}