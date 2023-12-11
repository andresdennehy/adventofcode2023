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
    row int
    col int
}

// Can't use constant arrays in Go
var toDown = []byte{'S', '|', 'F', '7'}
var toUp = []byte{'S', '|', 'J', 'L'}
var toLeft = []byte{'S', '-', 'J', '7'}
var toRight = []byte{'S', '-', 'F', 'L'}

var fromUp = []byte{'|', 'L', 'J'}
var fromDown = []byte{'|', '7', 'F'}
var fromRight = []byte{'-', 'L', 'F'}
var fromLeft = []byte{'-', '7', 'J'}

func characterInList(b byte, list []byte) bool {
    for _, item := range list {
        if item == b {
            return true
        }
    }
    return false
}

func coordinateInList(c *Coordinate, list []*Coordinate) bool {
    for _, item := range list {
        if item.row == c.row && item.col == c.col {
            return true
        }
    }
    return false
}

func DFS(pipe_map [][]byte, start *Coordinate) int {
    // Iterative DFS that returns length of path
    seen := []*Coordinate{};
    stack := []*Coordinate{start};
    for len(stack) > 0 {
        current := stack[len(stack)-1];
        seen = append(seen, current);
        stack = stack[:len(stack)-1];

        cell := pipe_map[current.row][current.col]

        // We are going down
        if current.row < len(pipe_map)-1 &&
          characterInList(cell, toDown) &&
          characterInList(pipe_map[current.row+1][current.col], fromUp) &&
          !coordinateInList(&Coordinate{current.row+1, current.col}, seen) {
            stack = append(stack, &Coordinate{current.row+1, current.col})
        }
        // We are going up
        if current.row > 0 &&
          characterInList(cell, toUp) &&
          characterInList(pipe_map[current.row-1][current.col], fromDown) &&
          !coordinateInList(&Coordinate{current.row-1, current.col}, seen) {
            stack = append(stack, &Coordinate{current.row-1, current.col})
        }
        // We are going left
        if current.col > 0 &&
          characterInList(cell, toLeft) &&
          characterInList(pipe_map[current.row][current.col-1], fromRight) &&
          !coordinateInList(&Coordinate{current.row, current.col-1}, seen) {
            stack = append(stack, &Coordinate{current.row, current.col-1})
        }
        // We are going right
        if current.col < len(pipe_map[0])-1 &&
          characterInList(cell, toRight) &&
          characterInList(pipe_map[current.row][current.col+1], fromLeft) &&
          !coordinateInList(&Coordinate{current.row, current.col+1}, seen) {
            stack = append(stack, &Coordinate{current.row, current.col+1})
        }
    }
    return len(seen)
}

func main() {

    file, err := os.Open("input.txt");
    check(err);
    defer file.Close();

    scanner := bufio.NewScanner(file);
    var pipe_map[][]byte;

    for scanner.Scan() {
        line := scanner.Text()
        pipe_line := []byte{}
        for i:=0; i<len(line); i++ {
            pipe_line = append(pipe_line, line[i])
        }
        pipe_map = append(pipe_map, pipe_line)
    }

    var start_position *Coordinate;
    for row, pipe_line := range pipe_map {
        for col, pipe := range pipe_line {
            if pipe == 'S' {
                start_position = &Coordinate{row, col}
            }
        }
    }

    fmt.Printf("Result is %d\n", DFS(pipe_map, start_position) / 2)
}
