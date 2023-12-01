package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
)

var numbers = map[string]int{
    "zero": 0,
    "one": 1,
    "two": 2,
    "three": 3,
    "four": 4,
    "five": 5,
    "six": 6,
    "seven": 7,
    "eight": 8,
    "nine": 9,
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func check_for_number(line string, i int) int {
    if line[i] <= 57 && line[i] >= 48 {
        number, _ := strconv.Atoi(string(line[i]));
        return number;
    } else {
        // Check for string numbers
        for number_string, number := range numbers {
            if i+len(number_string)-1 < len(line) && line[i:i+len(number_string)] == number_string {
                return number;
            }
        }
    }
    return -1;
}

func main() {

    file, err := os.Open("input.txt");
    check(err);
    defer file.Close();

    scanner := bufio.NewScanner(file);
    value_sum := 0;

    for scanner.Scan() {
        line := scanner.Text();
        first_number := 0;
        last_number := 0;

        // Find first number
        for i:=0; i<len(line); i++ {
            res := check_for_number(line, i);
            if res != -1 {
                first_number = res;
                break;
            }
        }

        // Find last number
        for i:=len(line)-1; i>=0; i-- {
            res := check_for_number(line, i);
            if res != -1 {
                last_number = res;
                break;
            }
        }

        fmt.Printf("First number is %v, last number is %v\n", first_number, last_number);
        value_sum += first_number * 10 + last_number;
    }

    fmt.Printf("Result is %v\n", value_sum);
}
