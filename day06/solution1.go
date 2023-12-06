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
    var times []int;
    var record_distances []int;

    // Parse times
    scanner.Scan()
    for _, time := range strings.Fields(scanner.Text())[1:] {
        parsed_time, err := strconv.Atoi(time)
        check(err)
        times = append(times, parsed_time)
    }

    scanner.Scan()
    for _, record_distance := range strings.Fields(scanner.Text())[1:] {
        parsed_distance, err := strconv.Atoi(record_distance)
        check(err);
        record_distances = append(record_distances, parsed_distance)
    }

    var ways_to_beat int;
    for race:=0; race<len(times); race++ {
        var ways_to_beat_race int;
        race_time, record_distance := times[race], record_distances[race]
        for release_time:=1; release_time<race_time; release_time++ {
            if release_time * (race_time - release_time) > record_distance {
                ways_to_beat_race++;
            }
        }
        fmt.Printf("Ways to beat race %v: %v\n", race, ways_to_beat_race);
        if ways_to_beat == 0 {
            ways_to_beat = ways_to_beat_race;
        } else {
            ways_to_beat *= ways_to_beat_race;
        }
    }

    fmt.Printf("Result is %v\n", ways_to_beat);
}
