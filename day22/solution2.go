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


type brick struct {
    fromX, fromY, fromZ, toX, toY, toZ int
}

var bricks []brick;

func min(a int, b int) int { if a < b { return a }; return b }
func max(a int, b int) int { if a > b { return a }; return b }

func (b brick) intersect(other brick) bool {
	return max(b.fromX, other.fromX) <= min(b.toX, other.toX) &&
           max(b.fromY, other.fromY) <= min(b.toY, other.toY) &&
           max(b.fromZ, other.fromZ) <= min(b.toZ, other.toZ)
}

func (b brick) fall() brick {
	return brick{ b.fromX, b.fromY, b.fromZ - 1, b.toX, b.toY, b.toZ - 1 }
}

func (b brick) canFall(bricks []brick) bool {
	var fallen = b.fall()
	for _, bb := range bricks {
		if b != bb && fallen.intersect(bb) {
			return false
		}
	}
	return b.fromZ > 1
}

func main() {
    file, err := os.Open("input.txt");
    check(err);
    defer file.Close();

    scanner := bufio.NewScanner(file);
    for scanner.Scan() {
        line := scanner.Text();
        var startX, startY, startZ int;
        var endX, endY, endZ int;
        fmt.Sscanf(line, "%d,%d,%d~%d,%d,%d", &startX, &startY, &startZ, &endX, &endY, &endZ);

        bricks = append(bricks, brick{startX, startY, startZ, endX, endY, endZ});
    }

    // Sort bricks slice by Z
    sort.Slice(bricks, func(i, j int) bool {
    	return bricks[i].fromZ < bricks[j].fromZ
    })
    // Make bricks fall as far down as possible
    for i, _ := range bricks {
		for bricks[i].canFall(bricks) {
			bricks[i] = bricks[i].fall()
		}
	}

	var removableBlocks int;
	for i, b := range bricks {
		var prev = b
		bricks[i] = brick{ -1, -1, -1, -1, -1, -1 }
		var canFall = false
		for j, bb := range bricks {
			if i != j && bb.canFall(bricks) {
				canFall = true
				break
			}
		}
		if canFall {
            var fallen = make([]brick, len(bricks))
			if copy(fallen, bricks) != len(bricks) { panic("copy failed") }
            for j, bb := range fallen {
                if i != j && bb.canFall(fallen) {
                    removableBlocks++
					for fallen[j].canFall(fallen) {
						fallen[j] = fallen[j].fall()
					}
                }
            }
		}
		bricks[i] = prev
	}

    fmt.Printf("Result: %v\n", removableBlocks);

}