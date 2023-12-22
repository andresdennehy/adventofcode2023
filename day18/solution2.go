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
    length int64;
}

func main() {
    file, err := os.Open("input.txt");
    check(err);
    defer file.Close();

    var edges []Edge = make([]Edge, 0);
    directions := map[string]string {
        "0": "R",
        "1": "D",
        "2": "L",
        "3": "U",
    }

    scanner := bufio.NewScanner(file);
    for scanner.Scan() {
        var line = strings.Fields(scanner.Text());
        // Trim parentheses
        color := line[2][2:len(line[2])-1];
        fmt.Printf("Color: %s\n", color)
        // Length is first five characters of hex
        length, err := strconv.ParseUint(color[:len(color)-1], 16, 32);
        check(err)
        direction := directions[color[len(color)-1:]]
        edges = append(edges, Edge{direction, int64(length)});
    }

    // Add first edge to end of list to make it a loop
    edges = append(edges, edges[0])

    var trenchArea int64;
    var trenchPerimeter int64;
    var currentX, currentY int64;
    // Calculate area of trench with shoelace formula
    for _, edge := range edges[:len(edges)-1] {
        var nextX, nextY int64 = currentX, currentY;
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

    trenchArea = int64(math.Abs(float64(trenchArea))) / 2;
    // Pick's theorem
    interiorPoints := trenchArea - trenchPerimeter / 2 + 1;

    fmt.Printf("Result is: %d\n", trenchPerimeter + interiorPoints);

}