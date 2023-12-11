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

func allZeros(sequence []int) bool {
    for _, value := range sequence {
        if value != 0 {
            return false
        }
    }
    return true
}

func main() {

    file, err := os.Open("input.txt");
    check(err);
    defer file.Close();

    scanner := bufio.NewScanner(file);
    sequences := make([][]int, 0)
    var result int;

    for scanner.Scan() {
        sequence := make([]int, 0)
        line := strings.Fields(scanner.Text());
        for _, value := range line {
            number, err := strconv.Atoi(value);
            check(err);
            sequence = append(sequence, number)
        }
        sequences = append(sequences, sequence)
    }

    for _, sequence := range sequences {
        last_numbers := []int{sequence[len(sequence)-1]}
        for !allZeros(sequence) {
            new_sequence := make([]int, 0)
            for i:=1; i<len(sequence); i++ {
                new_sequence = append(new_sequence, sequence[i] - sequence[i-1])
            }
            last_numbers = append(last_numbers, new_sequence[len(new_sequence)-1])
            sequence = new_sequence
        }

        var next_value int;
        for i:=len(last_numbers)-1; i>=0; i-- {
            next_value += last_numbers[i]
        }
        result += next_value
    }

    fmt.Printf("Result is %d\n", result)
}
