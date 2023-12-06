package main

import (
    "bufio"
    "fmt"
    "os"
    "math"
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
    var seeds []int;
    var mappings [][][]int;

    // Parse all inputs
    for scanner.Scan() {
        line := scanner.Text();
        if strings.HasPrefix(line, "seeds:") {
            for _, seed_string := range strings.Fields(strings.Split(line, ": ")[1]) {
                seed_number, err := strconv.Atoi(seed_string)
                check(err)
                seeds = append(seeds, seed_number);
            }
        } else if strings.HasSuffix(line, "map:") {
            var object_mapping [][]int;
            scanner.Scan();
            for scanner.Text() != "" {
                mapping := make([]int, 3);
                values := strings.Split(scanner.Text(), " ");;
                for i, component := range values {
                    mapping[i], err = strconv.Atoi(component);
                    check(err);
                }
                object_mapping = append(object_mapping, mapping);
                scanner.Scan();
            }
            mappings = append(mappings, object_mapping);
        }
    }

    lowest_location_seed := math.MaxInt64;
    // Iterate over all seeds and find the location with lowest number
    for _, seed := range seeds {
        current_value := seed;
        for _, object_mapping := range mappings {
            var destination_value int;
            for _, mapping := range object_mapping {
                destination_range_start, source_range_start, range_length := mapping[0], mapping[1], mapping[2];
                if current_value >= source_range_start && current_value < source_range_start + range_length {
                    destination_value = destination_range_start + (current_value - source_range_start);
                    current_value = destination_value;
                    break;
                }
            }
        }
        // current_value at this point is location
        if current_value < lowest_location_seed {
            lowest_location_seed = current_value;
        }
    }

    fmt.Printf("Result is %v\n", lowest_location_seed);
}
