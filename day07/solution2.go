package main

import (
    "bufio"
    "fmt"
    "os"
    "sort"
    "strconv"
    "strings"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

type Hand struct {
    cards string
    value int
    bet int
}

var CardValues = map[byte]int{
    'A': 13,
    'K': 12,
    'Q': 11,
    'J': 0,
    'T': 9,
    '9': 8,
    '8': 7,
    '7': 6,
    '6': 5,
    '5': 4,
    '4': 3,
    '3': 2,
    '2': 1,
}

func is_less(first, second Hand) bool {
    // Stable lexicographic ordering, first on value
    // then on cards one by one
    if first.value < second.value {
        return true
    }
    if first.value == second.value {
        for i:= 0; i < len(first.cards); i++ {
            if CardValues[first.cards[i]] < CardValues[second.cards[i]] {
                return true
            } else if CardValues[first.cards[i]] > CardValues[second.cards[i]] {
                return false
            }
        }
    }
    return false
}

func assign_value(hand string) int {
    /*
    Five of a kind 7
    Four of a kind 6
    Full house 5
    Three of a kind 4
    Two pair 3
    One pair 2
    High card 1
    */
    card_counts := make(map[rune]int)
    var jokers int;
    for _, card := range hand {
        if card == 'J' {
            jokers++;
            continue
        }
        card_counts[card]++;
    }

    if jokers == 5 { return 7; }

    keys := make([]rune, 0)

	for key := range card_counts {
		keys = append(keys, key)
	}

    sort.SliceStable(keys, func(i, j int) bool {
        return card_counts[keys[i]] > card_counts[keys[j]]
    })

    card_counts[keys[0]] += jokers

    switch len(card_counts) {
        case 1:
            return 7 // Five of a kind
        case 5:
            return 1 // High card
        default:
            for _, value := range card_counts {
                switch value {
                    case 4:
                        return 6 // Four of a kind
                    case 3:
                        if len(card_counts) == 2 {
                            return 5 // Full house
                        }
                        return 4 // Three of a kind
                    case 2:
                        if len(card_counts) == 3 {
                            return 3 // Two pair
                        }
                }
            }
            return 2
	}
}

func main() {

    file, err := os.Open("input.txt");
    check(err);
    defer file.Close();

    scanner := bufio.NewScanner(file);
    hands := make([]Hand, 0);

    for scanner.Scan() {
        line := strings.Fields(scanner.Text());
        hand := line[0];
        bet, err := strconv.Atoi(line[1]);
        check(err);
        hands = append(hands, Hand{hand, assign_value(hand), bet});
    }

    sort.SliceStable(hands, func(i, j int) bool {
        return is_less(hands[i], hands[j])
    })
    fmt.Println(hands)

    var score int;
    for i, hand := range hands {
        score += (i+1) * hand.bet
        fmt.Printf("Adding %v * %v\n", (i+1), hand.bet)
    }

    fmt.Printf("Result is %v\n", score);
}
