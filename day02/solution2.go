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

func main() {

    file, err := os.Open("input.txt");
    check(err);
    defer file.Close();

    scanner := bufio.NewScanner(file);
    cube_power := 0;

    for scanner.Scan() {
        line := scanner.Text();
        var game_id int;
        var game string;

        split_line := strings.Split(line, ": ");
        fmt.Sscanf(split_line[0], "Game %d", &game_id);
        game = split_line[1];

        plays := strings.Split(game, "; ")
        var min_red, min_green, min_blue int;
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
            if red > min_red {
                min_red = red;
            }
            if green > min_green {
                min_green = green;
            }
            if blue > min_blue {
                min_blue = blue;
            }
        }

        cube_power += min_red * min_green * min_blue;
    }

    fmt.Printf("Result is %v\n", cube_power);
}
