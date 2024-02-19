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

type Hailstone struct {
    x, y, z float64
    dx, dy, dz float64
}

func findIntersection(a, b Hailstone) (float64, float64, bool) {
    // If they are parallel, they will never intersect
    if a.dy / a.dx == b.dy / b.dx && a.y != b.y && a.x != b.x {
        return 0, 0, false
    }

    numerator := b.y - a.y + (b.dy / b.dx) * (a.x - b.x)
    denominator := a.dy - a.dx * (b.dy / b.dx)
    p1 := numerator / denominator

    if p1 < 0 {
        return 0, 0, false
    }

    p2 := (a.x - b.x + a.dx * p1) / b.dx

    if p2 < 0 {
        return 0, 0, false
    }

    x := a.x + a.dx * p1
    y := a.y + a.dy * p1

    return x, y, true
}

func isInTestArea(x, y, minX, maxX, minY, maxY float64) bool {
    return x >= minX && x <= maxX && y >= minY && y <= maxY
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var hailstones []Hailstone;
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var x, y, z int
		var dx, dy, dz int
		fmt.Sscanf(scanner.Text(), "%d, %d, %d @ %d, %d, %d", &x, &y, &z, &dx, &dy, &dz)
		hailstones = append(hailstones, Hailstone{float64(x), float64(y), float64(z), float64(dx), float64(dy), float64(dz)})
	}

	range_start := float64(200000000000000)
	range_end := float64(400000000000000)

    collisions := 0
	for i := 0; i < len(hailstones); i++{
        for j := i + 1; j < len(hailstones); j++{
            h1 := hailstones[i]
            h2 := hailstones[j]
            // Find the intersection point
            x, y, intersect := findIntersection(h1, h2)
            if !intersect {
                continue
            }
            // Check if the intersection point is within the test area
            if isInTestArea(x, y, range_start, range_end, range_start, range_end) {
                fmt.Printf("Collision between %d and %d at %d, %d\n", i, j, x, y)
                collisions++
            }
        }
    }

	fmt.Println(collisions)
}