package main

import (
    "bufio"
    "fmt"
    "math"
    "os"
    "strconv"
    "strings"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

type Edge struct {
    direction string;
    length int;
}

func main() {
    file, err := os.Open("input.txt");
    check(err);
    defer file.Close();

    var edges []Edge = make([]Edge, 0);

    scanner := bufio.NewScanner(file);
    for scanner.Scan() {
        var line = strings.Fields(scanner.Text());
        direction := line[0];
        length, err := strconv.Atoi(line[1]);
        check(err);
        edges = append(edges, Edge{direction, length});
    }

    // Add first edge to end of list to make it a loop
    edges = append(edges, edges[0])

    var trenchArea int;
    var trenchPerimeter int;
    var currentX, currentY int;
    // Calculate area of trench with shoelace formula
    for _, edge := range edges[:len(edges)-1] {
        var nextX, nextY int = currentX, currentY;
        trenchPerimeter += edge.length;

        switch edge.direction {
        case "R":
            nextX = currentX + edge.length;
        case "L":
            nextX = currentX - edge.length;
        case "U":
            nextY = currentY - edge.length;
        case "D":
            nextY = currentY + edge.length;
        }

        trenchArea += (currentX - nextX) * (currentY + nextY);
        currentX, currentY = nextX, nextY;
    }

    trenchArea = int(math.Abs(float64(trenchArea))) / 2;
    // Pick's theorem
    interiorPoints := trenchArea - trenchPerimeter / 2 + 1;

    fmt.Printf("Result is: %d\n", trenchPerimeter + interiorPoints);

}