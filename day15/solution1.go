package main

import (
    "bufio"
    "bytes"
    "fmt"
    "os"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

// I'm sure the input list isn't too long, but I feel like the right way is
// to scan data by commas
func ScanCSV(data []byte, atEOF bool) (advance int, token []byte, err error) {
    commaidx := bytes.IndexByte(data, ',')
    if commaidx > 0 {
        // we need to return the next position
        buffer := data[:commaidx]
        return commaidx + 1, bytes.TrimSpace(buffer), nil
    }

    // if we are at the end of the string, just return the entire buffer
    if atEOF {
        // but only do that when there is some data. If not, this might mean
        // that we've reached the end of our input CSV string
        if len(data) > 0 {
            return len(data), bytes.TrimSpace(data), nil
        }
    }

    // when 0, nil, nil is returned, this is a signal to the interface to read
    // more data in from the input reader. In this case, this input is our
    // string reader and this pretty much will never occur.
    return 0, nil, nil
}

func main() {
    file, err := os.Open("input.txt");
    check(err);
    defer file.Close();

    scanner := bufio.NewScanner(file);
    scanner.Split(ScanCSV);

    var result int;
    for scanner.Scan() {
        var currentValue int;
        // For each word
        word := scanner.Text();

        for _, c := range word {
            currentValue += int(c);
            currentValue *= 17;
            currentValue = currentValue % 256;
        }
        result += currentValue;
    }

    fmt.Printf("Result is %d\n", result);

}