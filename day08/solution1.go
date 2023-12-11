package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func parseLeftRight(s string) (string, string) {
    parts := strings.Split(s, ", ");
    left, right := parts[0], parts[1];
    left = strings.Replace(left, "(", "", 1);
    right = strings.Replace(right, ")", "", 1);
    return left, right
}

func main() {

    file, err := os.Open("input.txt");
    check(err);
    defer file.Close();

    scanner := bufio.NewScanner(file);
    nodes := make(map[string][]string, 0)

    scanner.Scan();
    orders := scanner.Text();

    // Empty line
    scanner.Scan();

    // First element
    scanner.Scan();
    line := strings.Split(scanner.Text(), " = ");
    left, right := parseLeftRight(line[1]);
    first_element := line[0]
    nodes[line[0]] = []string{left, right};

    // Rest
    for scanner.Scan() {
        line := strings.Split(scanner.Text(), " = ");
        left, right := parseLeftRight(line[1]);
        nodes[line[0]] = []string{left, right};
    }

    // Find next node according to orders
    // Brute force...
    current_node := first_element;
    current_order := 0;
    var steps int;
    for current_node != "ZZZ" {
        if orders[current_order] == 'L' {
            current_node = nodes[current_node][0]
        } else {
            current_node = nodes[current_node][1]
        }
        if current_order < len(orders) - 1 {
            current_order++
        } else {
            current_order = 0
        }
        steps++
    }

    fmt.Printf("Result is %v\n", steps);
}
