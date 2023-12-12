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

func printGrid(grid [][]byte) {
    // For debugging
    for i:=0; i<len(grid); i++ {
        for j:=0; j<len(grid[0]); j++ {
            fmt.Printf("%c", grid[i][j])
        }
        fmt.Printf("\n")
    }
    fmt.Printf("\n")
}

func isIn(a []int, b int) bool {
    for _, i := range a {
        if i == b {
            return true
        }
    }
    return false
}

func abs(a int) int {
    if a < 0 {
        return -a
    }
    return a
}

func expandGalaxies(grid [][]byte) [][]byte {
    /* Not memory efficient. There must be a way to calculate using
    math instead of actually manipulating the array. */

    // Expand empty cols
    for j:=0; j<len(grid[0]); j++ {
        empty := true;
        for i:=0; i<len(grid); i++ {
            if grid[i][j] == '#' {
                empty = false;
                break;
            }
        }
        if empty {
            for k:=0; k<len(grid); k++ {
                grid[k] = append(grid[k][:j+1], grid[k][j:]...)
                grid[k][j] = '.'
            }
            j++
        }
    }

    emptyRow := make([]byte, len(grid[0]))
    for i:=0; i<len(grid[0]); i++ {
        emptyRow[i] = '.'
    }

    // Expand empty rows
    for i:=0; i<len(grid); i++ {
        empty := true;
        for j:=0; j<len(grid[0]); j++ {
            if grid[i][j] == '#' {
                empty = false;
                break;
            }
        }
        if empty {
            grid = append(grid[:i+1], grid[i:]...)
            grid[i] = emptyRow
            i++
        }
    }

    return grid
}

func main() {

    file, err := os.Open("input.txt");
    check(err);
    defer file.Close();

    scanner := bufio.NewScanner(file);
    var grid [][]byte;
    var galaxies []*Point;

    for scanner.Scan() {
        line := scanner.Text()
        gridLine := make([]byte, 0)
        for i:=0; i<len(line); i++ {
            gridLine = append(gridLine, line[i])
        }
        grid = append(grid, gridLine)
    }

    fmt.Println("");
    grid = expandGalaxies(grid);
    printGrid(grid);

    for i:=0; i<len(grid); i++ {
        for j:=0; j<len(grid[0]); j++ {
            if grid[i][j] == '#' {
                galaxies = append(galaxies, &Point{row: i, col: j})
            }
        }
    }

    var total_distance int;
    for i:=0; i<len(galaxies); i++ {
        for j:=i+1; j<len(galaxies); j++ {
            distance := abs(galaxies[j].row - galaxies[i].row) + abs(galaxies[j].col - galaxies[i].col)
            total_distance += distance;
        }
    }
    fmt.Printf("Result is %d\n", total_distance)
}
