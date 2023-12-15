package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func findUnknownPipes(springs string) []int {
    // Find all indices of unknown pipes
    indices := make([]int, 0);

    for i:=0; i<len(springs); i++ {
        if springs[i] == '?' {
            indices = append(indices, i);
        }
    }

    return indices;
}

func isValid(springs string, springList []int) bool {
    damagedGroups := strings.Fields(strings.ReplaceAll(springs, ".", " "));

    if len(damagedGroups) != len(springList) {
        return false;
    }

    for i:=0; i<len(damagedGroups); i++ {
        if springList[i] != len(damagedGroups[i]) {
            return false;
        }
    }

    return true;
}

func bruteForce(springs string, springList []int) int {
    unknownPipes := findUnknownPipes(springs);

    if len(unknownPipes) == 0 {
        if isValid(springs, springList) {
            return 1;
        } else {
            return 0;
        }
    }

    var arrangements int;
    arrangements += bruteForce(springs[:unknownPipes[0]] + "." + springs[unknownPipes[0]+1:], springList);
    arrangements += bruteForce(springs[:unknownPipes[0]] + "#" + springs[unknownPipes[0]+1:], springList);

    return arrangements
}

func solve(scanner *bufio.Scanner) int {
    var result int;

    for (*scanner).Scan() {
        line := strings.Fields((*scanner).Text());
        springs := line[0]
        springList := make([]int, 0);
        for _, number := range strings.Split(line[1], ",") {
            springCount, err := strconv.Atoi(number);
            check(err);
            springList = append(springList, springCount);
        }
        combinations := bruteForce(springs, springList);
        fmt.Printf("Springs: %s, List: %v, Combinations: %d\n", springs, springList, combinations);
        result += combinations;
    }

    return result;
}

func main() {

    file, err := os.Open("input.txt");
    check(err);
    defer file.Close();

    scanner := bufio.NewScanner(file);
    fmt.Printf("Result is %d\n", solve(scanner));

}