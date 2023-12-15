package main

import (
    "bufio"
    "bytes"
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

type Lens struct {
    label string
    focalLength int
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

func hash(label string) int {
    var hash int;
    for _, c := range label {
        hash += int(c);
        hash *= 17;
        hash = hash % 256;
    }
    return hash
}

func deleteLens(lenses []*Lens, label string) []*Lens {
    for i, lens := range lenses {
        if lens.label == label {
            if len(lenses) == 1 {
                lenses = []*Lens{};
            } else if i == len(lenses) - 1 {
                lenses = lenses[:i];
            } else {
                lenses = append(lenses[:i], lenses[i+1:]...);
            }
            break;
        }
    }
    return lenses
}

func isIn(lenses []*Lens, label string) bool {
    for _, lens := range lenses {
        if lens.label == label {
            return true;
        }
    }
    return false
}

func replaceLens(lenses []*Lens, label string, focalLength int) []*Lens {
    for i, lens := range lenses {
        if lens.label == label {
            lenses[i].focalLength = focalLength;
            break;
        }
    }
    return lenses
}

func main() {
    file, err := os.Open("input.txt");
    check(err);
    defer file.Close();

    scanner := bufio.NewScanner(file);
    scanner.Split(ScanCSV);

    boxes := make(map[int][]*Lens);
    for scanner.Scan() {
        // For each word
        word := scanner.Text();

        if word[len(word)-1] == '-' {
            label := word[:len(word)-1];
            wordHash := hash(label);
            if _, ok := boxes[wordHash]; ok {
                boxes[wordHash] = deleteLens(boxes[wordHash], label);
            }
        } else {
            // Only - and = are used in the input
            // We are in the = case
            parts := strings.Split(word, "=");
            label := parts[0]
            // Assume no errors in strconv.Atoi :-)
            focalLength, _ := strconv.Atoi(parts[1]);
            wordHash := hash(label);
            if _, ok := boxes[wordHash]; ok {
                if isIn(boxes[wordHash], label) {
                    boxes[wordHash] = replaceLens(boxes[wordHash], label, focalLength);
                } else {
                    boxes[wordHash] = append(boxes[wordHash], &Lens{label, focalLength});
                }
            } else {
                boxes[wordHash] = []*Lens{&Lens{label, focalLength}};
            }
        }
    }

    var result int;
    for boxNumber, boxLenses := range boxes {
        for lensNumber, lens := range boxLenses {
            fmt.Printf("Box %d, label %v, lens %d, focal length %d\n", boxNumber+1, lens.label, lensNumber+1, lens.focalLength);
            result += (boxNumber+1) * (lensNumber+1) * lens.focalLength;
        }
    }

    fmt.Printf("Result is %d\n", result);

}