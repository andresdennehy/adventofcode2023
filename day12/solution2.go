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

func countCharacters(springs string, character byte) int {
    var count int;

    for i:=0; i<len(springs); i++ {
        if springs[i] == character {
            count++;
        }
    }

    return count;
}

func max(a []int) int {
    var result int;
    for i:=0; i<len(a); i++ {
        if a[i] > result {
            result = a[i];
        }
    }
    return result;
}

func DP(springs string, springList []int) int {
    springList = append(springList, 0);
    springs = springs + ".";

    n := len(springs);
    m := len(springList);
    max_run := max(springList);
    // DP memos are n x m x k, where n = springs[0:n], m = springList[0:m], k = length of current run
    dp := make([][][]int, n);
    for i:=0; i<n; i++ {
        dp[i] = make([][]int, m);
        for j:=0; j<m; j++ {
            dp[i][j] = make([]int, max_run+1);
        }
    }

    // Bottom-up DP
    for i:=0; i<n; i++ {
        current_char := springs[i];
        for j:=0; j<m; j++ {
            for k:=0; k<=springList[j]; k++ {
                if i == 0 {
                    if j != 0 {
                        dp[i][j][k] = 0;
                        continue;
                    }

                    if current_char == '#' {
                        if k != 1 {
                            dp[i][j][k] = 0;
                            continue;
                        }
                        dp[i][j][k] = 1;
                        continue;
                    }

                    if current_char == '.' {
                        if k != 0 {
                            dp[i][j][k] = 0;
                            continue;
                        }
                        dp[i][j][k] = 1;
                        continue;
                    }

                    if current_char == '?' {
                        if k != 0 && k != 1 {
                            dp[i][j][k] = 0;
                            continue;
                        }
                        dp[i][j][k] = 1;
                        continue;
                    }
                }

                var answerDot, answerHash int;

                // Process answer
                if k != 0 {
                    answerDot = 0;
                    answerHash = dp[i-1][j][k-1];
                } else if j > 0 && k == 0 {
                    answerDot = dp[i-1][j-1][springList[j-1]];
                    answerDot += dp[i-1][j][0];
                } else if countCharacters(springs[0:i], byte('#')) == 0 {
                    answerDot = 1;
                }

                switch current_char {
                case '.':
                    dp[i][j][k] = answerDot;
                case '#':
                    dp[i][j][k] = answerHash;
                case '?':
                    dp[i][j][k] = answerDot + answerHash;
                }
            }
        }
    }

    return dp[n-1][m-1][0];
}

func solve(scanner *bufio.Scanner) int {
    var result int;

    for (*scanner).Scan() {
        line := strings.Fields((*scanner).Text());
        springs := line[0]
        springList := make([]int, 0);
        for _, number := range strings.Split(line[1], ",") {
            springCount, err := strconv.Atoi(number);
            check(err);
            springList = append(springList, springCount);
        }

        originalSprings := springs;
        originalSpringList := springList;
        // Unfold
        for i:=1; i<5; i++ {
            springs += "?" + originalSprings;
            springList = append(springList, originalSpringList...);
        }

        combinations := DP(springs, springList);
        fmt.Printf("Springs: %s, List: %v, Combinations: %d\n", springs, springList, combinations);
        result += combinations;
    }

    return result;
}

func main() {

    file, err := os.Open("input.txt");
    check(err);
    defer file.Close();

    scanner := bufio.NewScanner(file);
    fmt.Printf("Result is %d\n", solve(scanner));

}