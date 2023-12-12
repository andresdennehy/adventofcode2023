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
    row int
    col int
}

func printGrid(grid [][]int) {
    // For debugging
    for i:=0; i<len(grid); i++ {
        for j:=0; j<len(grid[0]); j++ {
            if grid[i][j] == 1 {
                fmt.Printf("#")
            } else {
                fmt.Printf(".")
            }
        }
        fmt.Printf("\n")
    }
    fmt.Printf("\n")
}

func isIn(value int, array []int) bool {
    for _, element := range array {
        if element == value {
            return true
        }
    }
    return false
}

func findEmptyRows(grid [][]int) []int {
    emptyRows := make([]int, 0)

    for i:=0; i<len(grid); i++ {
        empty := true;
        for j:=0; j<len(grid[0]); j++ {
            if grid[i][j] == 1 {
                empty = false;
                break;
            }
        }
        if empty {
            emptyRows = append(emptyRows, i)
        }
    }

    return emptyRows
}

func findEmptyCols(grid [][]int) []int {
    emptyCols := make([]int, 0)

    for j:=0; j<len(grid[0]); j++ {
        empty := true;
        for i:=0; i<len(grid); i++ {
            if grid[i][j] == 1 {
                empty = false;
                break;
            }
        }
        if empty {
            emptyCols = append(emptyCols, j)
        }
    }

    return emptyCols
}

func calculateDistance(grid [][]int, start, end *Point, emptyRows, emptyCols []int) int {
    var distance int;

    rowDirection := 1;
    if start.row > end.row {
        rowDirection = -1
    }

    // Calculate height
    for i:=start.row; i!=end.row; i+=rowDirection {
        if isIn(i, emptyRows) {
            distance += 1000000
        } else {
            distance += 1
        }
    }

    colDirection := 1;
    if start.col > end.col {
        colDirection = -1
    }
    // Calculate width
    for j:=start.col; j!=end.col; j+=colDirection {
        if isIn(j, emptyCols) {
            distance += 1000000
        } else {
            distance += 1
        }
    }

    return distance
}

func main() {

    file, err := os.Open("input.txt");
    check(err);
    defer file.Close();

    scanner := bufio.NewScanner(file);
    var grid [][]int;
    var galaxies []*Point;

    for scanner.Scan() {
        line := scanner.Text()
        gridLine := make([]int, 0)
        for _, character := range line {
            if character == '.' {
                gridLine = append(gridLine, 0)
            } else {
                gridLine = append(gridLine, 1)
            }
        }
        grid = append(grid, gridLine)
    }

    printGrid(grid);
    emptyRows := findEmptyRows(grid);
    emptyCols := findEmptyCols(grid);

    for i:=0; i<len(grid); i++ {
        for j:=0; j<len(grid[0]); j++ {
            if grid[i][j] == 1 {
                galaxies = append(galaxies, &Point{row: i, col: j})
            }
        }
    }

    var total_distance int;
    for i:=0; i<len(galaxies); i++ {
        for j:=i+1; j<len(galaxies); j++ {
            distance := calculateDistance(grid, galaxies[i], galaxies[j], emptyRows, emptyCols)
            total_distance += distance;
        }
    }
    fmt.Printf("Result is %d\n", total_distance)
}
