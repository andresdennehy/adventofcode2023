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

    // Parse time
    scanner.Scan()
    var time_string string;
    for _, time := range strings.Fields(strings.Split(scanner.Text(), ":")[1]) {
        time_string += time;
    }
    race_time, err := strconv.Atoi(time_string)
    check(err)

    scanner.Scan()
    var record_distance_string string;
    for _, distance := range strings.Fields(strings.Split(scanner.Text(), ":")[1]) {
        record_distance_string += distance;
    }
    record_distance, err := strconv.Atoi(record_distance_string)
    check(err);

    fmt.Printf("Parsed time: %v\n", race_time);
    fmt.Printf("Parsed distance: %v\n", record_distance);

    var ways_to_beat int;
    for release_time:=1; release_time<race_time; release_time++ {
        if release_time * (race_time - release_time) > record_distance {
            ways_to_beat++;
        }
    }
    fmt.Printf("Result is %v\n", ways_to_beat);
}
