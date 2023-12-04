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

func main() {

    file, err := os.Open("input.txt");
    check(err);
    defer file.Close();

    scanner := bufio.NewScanner(file);
    winning_sum := 0;

    for scanner.Scan() {
        line := strings.Split(scanner.Text(), " | ");

        winning_numbers, numbers_we_have := line[0], line[1];
        winning_numbers = strings.Split(winning_numbers, ": ")[1];
        
        // With a hash map we avoid iterating over whole array
        winning_numbers_map := make(map[int]bool);
        for _, number := range strings.Fields(winning_numbers) {
            converted_number, err := strconv.Atoi(number);
            check(err)
            winning_numbers_map[converted_number] = true;
        }

        numbers_we_have_map := make(map[int]bool);
        for _, number := range strings.Fields(numbers_we_have) {
            converted_number, err := strconv.Atoi(number);
            check(err);
            numbers_we_have_map[converted_number] = true;
        }

        matching_numbers := make([]int, 0);
        for number_we_have, _ := range numbers_we_have_map {
            if _, ok := winning_numbers_map[number_we_have]; ok {
                matching_numbers = append(matching_numbers, number_we_have);
            }
        }

        score := 0;
        for index, _ := range matching_numbers {
            if index == 0 {
                score = 1;
            } else {
                score *= 2;
            }
        }

        winning_sum += score;
    }

    fmt.Printf("Result is %v\n", winning_sum);
}
