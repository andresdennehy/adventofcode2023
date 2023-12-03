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


func parse_number(row, col int, lines [][]byte) int {
    // Look to left and right of number for other numbers
    start, end := col, col;
    for start > 0 && lines[row][start-1] >= 48 && lines[row][start-1] <= 57 {
        start--;
    }
    for end < len(lines[0]) - 1 && lines[row][end+1] >= 48 && lines[row][end+1] <= 57 {
        end++;
    }
    number, err := strconv.Atoi(string(lines[row][start:end+1]));
    check(err);

    // Parse number only once
    for i:=start; i<=end; i++ {
        lines[row][i] = '.';
    }

    return number
}



func look_for_numbers(row, col int, lines [][]byte) int {
    part_numbers := make([]int, 0);

    // Look for numbers in all directions
    if row > 0 {
        for j:=col-1; j<=col+1; j++ {
            if 48 <= lines[row-1][j] && lines[row-1][j] <= 57 {
                part_numbers = append(part_numbers, parse_number(row-1, j, lines));
            }
        }
    }
    if row < len(lines) - 1 {
        for j:=col-1; j<=col+1; j++ {
            if 48 <= lines[row+1][j] && lines[row+1][j] <= 57 {
                part_numbers = append(part_numbers, parse_number(row+1, j, lines));
            }
        }
    }
    if col > 0 {
        for i:=row-1; i<=row+1; i++ {
            if 48 <= lines[i][col-1] && lines[i][col-1] <= 57 {
                part_numbers = append(part_numbers, parse_number(i, col-1, lines));
            }
        }
    }
    if col < len(lines[row]) - 1 {
        for i:=row-1; i<=row+1; i++ {
            if 48 <= lines[i][col+1] && lines[i][col+1] <= 57 {
                part_numbers = append(part_numbers, parse_number(i, col+1, lines));
            }
        }
    }

    if len(part_numbers) == 2 {
        return part_numbers[0] * part_numbers[1];
    }
    return 0
}


func main() {

    file, err := os.Open("input.txt");
    check(err);
    defer file.Close();

    scanner := bufio.NewScanner(file);
    lines := make([][]byte, 0);
    part_numbers := 0;

    for scanner.Scan() {
        line := make([]byte, 0);
        for _, character := range scanner.Text() {
            line = append(line, byte(character));
        }
        lines = append(lines, line);
    }

    for row, line := range lines {
        for col, character := range line {
            // ASCII codes for numbers are 48-57
            if character == '*' {
                part_numbers += look_for_numbers(row, col, lines);
            }
        }
    }

    fmt.Printf("Result is %v\n", part_numbers);
}
