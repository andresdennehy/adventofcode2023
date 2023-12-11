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

func GCD(a, b int) int {
      for b != 0 {
              t := b
              b = a % b
              a = t
      }
      return a
}

func LCM(a, b int, integers ...int) int {
      result := a * b / GCD(a, b)

      for i := 0; i < len(integers); i++ {
              result = LCM(result, integers[i])
      }

      return result
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
    a_nodes := make([]string, 0)

    scanner.Scan();
    orders := scanner.Text();

    // Empty line
    scanner.Scan();

    // Rest
    for scanner.Scan() {
        line := strings.Split(scanner.Text(), " = ");
        left, right := parseLeftRight(line[1]);
        nodes[line[0]] = []string{left, right};
        if line[0][len(line[0])-1] == 'A' {
            a_nodes = append(a_nodes, line[0])
        }
    }
    fmt.Println(a_nodes)

    var current_node string;
    var current_order int;
    a_node_steps := make([]int, 0);
    var steps int;
    for _, a_node := range a_nodes {
        current_node = a_node;
        current_order = 0;
        steps = 0;
        for current_node[len(current_node)-1] != 'Z' {
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
        fmt.Printf("Node %v took %v steps\n", a_node, steps)
        a_node_steps = append(a_node_steps, steps)
    }


    fmt.Printf("Result is %v\n", LCM(a_node_steps[0], a_node_steps[1], a_node_steps[2:]...));
}
