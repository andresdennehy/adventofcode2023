package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
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
    value_sum := 0;

    for scanner.Scan() {
        line := scanner.Text();
        first_number := 0;
        last_number := 0;

        // Find first number
        for _, char := range line {
            if char <= 57 && char >= 48 {
                first_number, err = strconv.Atoi(string(char));
                if err != nil {
                    panic(err);
                }
                break;
            }
        }
        for i:=len(line)-1; i>=0; i-- {
            if byte(line[i]) <= 57 && byte(line[i]) >= 48 {
                last_number, err = strconv.Atoi(string(line[i]));
                if err != nil {
                    panic(err);
                }
                break;
            }
        }

        fmt.Printf("First number is %v, last number is %v\n", first_number, last_number);
        value_sum += first_number * 10 + last_number;
    }

    fmt.Printf("Result is %v\n", value_sum);
}
