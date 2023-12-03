package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
    //"strconv"
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
    id_sum := 0;
    red_limit := 12;
    green_limit := 13;
    blue_limit := 14;

    for scanner.Scan() {
        line := scanner.Text();
        var game_id int;
        var game string;
        var possible bool = true;

        split_line := strings.Split(line, ": ");
        fmt.Sscanf(split_line[0], "Game %d", &game_id);
        game = split_line[1];

        plays := strings.Split(game, "; ")
        for _, play := range plays {
            var red, green, blue int;
            cubes := strings.Split(play, ", ");
            for _, cube := range cubes {
                var color string;
                var value int;
                fmt.Sscanf(cube, "%d %s", &value, &color);
                switch color {
                    case "red":
                        red = value;
                    case "green":
                        green = value;
                    case "blue":
                        blue = value;
                }
            }
            if red > red_limit || green > green_limit || blue > blue_limit {
                fmt.Printf("Game %d broke limit: %s\n", game_id, play);
                possible = false;
                break;
            }
        }

        if possible {
            id_sum += game_id;
        }
    }

    fmt.Printf("Result is %v\n", id_sum);
}
