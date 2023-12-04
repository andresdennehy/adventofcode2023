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
    //winning_sum := 0;
    scratchcard_counts := make(map[int]int, 0);

    current_card := 1;
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

        matching_numbers := make([]int, 0)
        for number_we_have, _ := range numbers_we_have_map {
            if _, ok := winning_numbers_map[number_we_have]; ok {
                matching_numbers = append(matching_numbers, number_we_have);
            }
        }

        scratchcard_counts[current_card] = len(matching_numbers);
        current_card += 1;
    }

    // Iterate over suffixes in reverse order
    for i:=len(scratchcard_counts)-1; i>0; i-- {
        subproblem_sum := 1;
        // We always have the card we are iterating over
        for j:=1; j<=scratchcard_counts[i] && j<=len(scratchcard_counts); j++ {
            subproblem_sum += scratchcard_counts[i+j];
        }
        scratchcard_counts[i] = subproblem_sum;
    }

    // This is just horrible code
    global_sum := 1;
    for i:=1; i<len(scratchcard_counts); i++ {
        global_sum += scratchcard_counts[i];
    }

    fmt.Printf("Total number of scratchcards is %v\n", global_sum);
}
