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

type Direction struct {
    dx int;
    dy int;
}

type Beam struct {
    x int;
    y int;
    direction Direction;
}

type Tile struct {
    x int;
    y int;
}

func (b *Beam) move(grid *[][]byte) (*Beam, bool) {
    // Check if we reached mirrors or splitters
    switch (*grid)[b.y][b.x] {
        case '/':
            if b.direction.dx == 0 {
                b.direction.dx = -b.direction.dy;
                b.direction.dy = 0;
            } else {
                b.direction.dy = -b.direction.dx;
                b.direction.dx = 0;
            }
        case '\\':
            if b.direction.dx == 0 {
                b.direction.dx = b.direction.dy;
                b.direction.dy = 0;
            } else {
                b.direction.dy = b.direction.dx;
                b.direction.dx = 0;
            }
        case '|':
            if b.direction.dx != 0 {
                // Change direction, return a new beam
                b.direction.dx = 0;
                b.direction.dy = -1;
                return &Beam{b.x, b.y, Direction{dx: 0, dy: 1}}, true;
            }
        case '-':
            if b.direction.dy != 0 {
                // Change direction, return a new beam
                b.direction.dx = -1;
                b.direction.dy = 0;
                return &Beam{b.x, b.y, Direction{dx: 1, dy: 0}}, true;
            }
    }

    // Stop if we reached the edge
    if (b.x == 0 && b.direction.dx < 0) || (b.x == len((*grid)[0])-1 && b.direction.dx > 0) {
        b.direction.dx = 0;
    }
    if (b.y == 0 && b.direction.dy < 0) || (b.y == len((*grid))-1 && b.direction.dy > 0) {
        b.direction.dy = 0;
    }

    if (b.x > 0 && b.direction.dx == -1) || (b.x < len((*grid)[0])-1 && b.direction.dx == 1) {
        b.x += b.direction.dx;
    }
    if (b.y > 0 && b.direction.dy == -1) || (b.y < len((*grid))-1 && b.direction.dy == 1) {
        b.y += b.direction.dy;
    }

    return nil, false;
}

func printGrid(grid [][]byte) {
    for i := 0; i < len(grid); i++ {
        fmt.Println(string(grid[i]));
    }
    fmt.Println();
}

func runBeams(grid [][]byte, startX, startY int, startDirection Direction) int {
    beamQueue := []*Beam{&Beam{startX, startY, startDirection}};
    seen := make(map[Tile]Direction);
    energized := make(map[Tile]bool);
    energized[Tile{startX, startY}] = true;

    for len(beamQueue) > 0 {
        currentBeam := beamQueue[0];
        beamQueue = beamQueue[1:];
        // If we've seen this beam before, pop it
        if direction, ok := seen[Tile{currentBeam.x, currentBeam.y}]; ok && direction == currentBeam.direction {
            continue;
        }
        // If it stopped, pop to never be seen again
        if currentBeam.direction.dx == 0 && currentBeam.direction.dy == 0 {
            continue;
        }
        seen[Tile{currentBeam.x, currentBeam.y}] = currentBeam.direction;
        if newBeam, newBeamCreated := currentBeam.move(&grid); newBeamCreated {
            beamQueue = append(beamQueue, newBeam);
        }
        beamQueue = append(beamQueue, currentBeam);
        if _, ok := energized[Tile{currentBeam.x, currentBeam.y}]; !ok {
            energized[Tile{currentBeam.x, currentBeam.y}] = true;
        }
    }

    return len(energized);
}

func max(nums ...int) int {
    var max int;
    for _, num := range nums {
        if num > max {
            max = num;
        }
    }
    return max;
}

func main() {
    file, err := os.Open("input.txt");
    check(err);
    defer file.Close();

    var grid [][]byte;
    scanner := bufio.NewScanner(file);
    for scanner.Scan() {
        var line = scanner.Text();
        var row []byte;
        for i := 0; i < len(line); i++ {
            row = append(row, line[i]);
        }
        grid = append(grid, row);
    }

    printGrid(grid);
    var maxEnergized int;
    for i := 0; i < len(grid); i++ {
        energizedTop := runBeams(grid, i, 0, Direction{dx: 0, dy: 1});
        energizedBottom := runBeams(grid, i, len(grid)-1, Direction{dx: 0, dy: -1});
        maxEnergized = max(maxEnergized, energizedTop, energizedBottom);
    }
    for j := 0; j < len(grid[0]); j++ {
        energizedLeft := runBeams(grid, 0, j, Direction{dx: 1, dy: 0});
        energizedRight := runBeams(grid, len(grid[0])-1, j, Direction{dx: -1, dy: 0});
        maxEnergized = max(maxEnergized, energizedLeft, energizedRight);
    }

    fmt.Printf("Result is %d\n", maxEnergized);

}