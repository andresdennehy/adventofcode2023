package main

import (
    "bufio"
    "fmt"
    "os"
    "sort"
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

type CoordinateSet struct {
    rocks map[Coordinate]bool;
}

func hashCoordinateSet(cs CoordinateSet) string {
    // Extract keys
    var keys []Coordinate
    for coord := range cs.rocks {
        keys = append(keys, coord)
    }

    // Sort keys
    sort.Slice(keys, func(i, j int) bool {
        if keys[i].row != keys[j].row {
            return keys[i].row < keys[j].row
        }
        return keys[i].col < keys[j].col
    })

    // Implement your own hashing logic here
    // This is a simple example; you might want a more robust hash function
    hash := ""
    for _, coord := range keys {
        hash += fmt.Sprintf("(%d,%d)", coord.row, coord.col)
    }
    return hash
}


func printGrid(grid [][]byte) {
    for i:=0; i<len(grid); i++ {
        for j:=0; j<len(grid[0]); j++ {
            fmt.Printf("%c", grid[i][j]);
        }
        fmt.Printf("\n");
    }
    fmt.Printf("\n");
}

func tiltNorth(grid [][]byte) {
    for i:=0; i<len(grid); i++ {
        for j:=0; j<len(grid[0]); j++ {
            if grid[i][j] == 'O' {
                currPos := Coordinate{i, j};
                for currPos.row > 0 && grid[currPos.row-1][currPos.col] == '.' {
                    grid[currPos.row][currPos.col] = '.';
                    grid[currPos.row-1][currPos.col] = 'O';
                    currPos.row--;
                }
            }
        }
    }
}

func tiltSouth(grid [][]byte) {
    for i:=len(grid)-1; i>=0; i-- {
        for j:=0; j<len(grid[0]); j++ {
            if grid[i][j] == 'O' {
                currPos := Coordinate{i, j};
                for currPos.row < len(grid)-1 && grid[currPos.row+1][currPos.col] == '.' {
                    grid[currPos.row][currPos.col] = '.';
                    grid[currPos.row+1][currPos.col] = 'O';
                    currPos.row++;
                }
            }
        }
    }
}

func tiltWest(grid [][]byte) {
    for i:=0; i<len(grid); i++ {
        for j:=0; j<len(grid[0]); j++ {
            if grid[i][j] == 'O' {
                currPos := Coordinate{i, j};
                for currPos.col > 0 && grid[currPos.row][currPos.col-1] == '.' {
                    grid[currPos.row][currPos.col] = '.';
                    grid[currPos.row][currPos.col-1] = 'O';
                    currPos.col--;
                }
            }
        }
    }
}

func tiltEast(grid [][]byte) {
    for i:=0; i<len(grid); i++ {
        for j:=len(grid[0])-1; j>=0; j-- {
            if grid[i][j] == 'O' {
                currPos := Coordinate{i, j};
                for currPos.col < len(grid[0])-1 && grid[currPos.row][currPos.col+1] == '.' {
                    grid[currPos.row][currPos.col] = '.';
                    grid[currPos.row][currPos.col+1] = 'O';
                    currPos.col++;
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

func gridToCoordinateSet(grid [][]byte) CoordinateSet {
    result := CoordinateSet{};
    result.rocks = make(map[Coordinate]bool);
    for i:=0; i<len(grid); i++ {
        for j:=0; j<len(grid[0]); j++ {
            if grid[i][j] == 'O' {
                coord := Coordinate{i, j};
                result.rocks[coord] = true;
            }
        }
    }
    return result;
}

func solve(scanner *bufio.Scanner) int {
    var grid [][]byte;
    var target int = 1000000000;

    for (*scanner).Scan() {
        line := (*scanner).Text();
        row := make([]byte, len(line));
        for i := 0; i < len(line); i++ {
            row[i] = line[i];
        }
        grid = append(grid, row);
    }

    printGrid(grid);

    stopCache := false;
    // Cache tells us if we've seen this grid before, in what iteration
    cache := make(map[string]int);
    for i:=0; i<target; i++ {
        tiltNorth(grid);
        tiltWest(grid);
        tiltSouth(grid);
        tiltEast(grid);
        if !stopCache {
            gridSet := gridToCoordinateSet(grid);
            hash := hashCoordinateSet(gridSet);
            if cycleStart, ok := cache[hash]; ok {
                fmt.Printf("Found cycle at iteration %d\n", i);
                fmt.Printf("Previous iteration where we saw this was %d\n", cycleStart);
                stopCache = true;
                cycleLength := i - cycleStart;
                fmt.Printf("Cycle length is %d\n", cycleLength);
                target = i + (target - cycleStart) % cycleLength;
                fmt.Printf("New target is %d\n", target);
            } else {
                cache[hash] = i;
            }
        }
    }

    return calculateScore(grid);
}

func main() {

    file, err := os.Open("input.txt");
    check(err);
    defer file.Close();

    scanner := bufio.NewScanner(file);
    fmt.Printf("Result is %d\n", solve(scanner));

}