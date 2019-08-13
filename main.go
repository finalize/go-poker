package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/shgysd/go-poker/pkg/deck"
	"github.com/shgysd/go-poker/pkg/game"
)

var d deck.Deck
var g game.Game

func init() {
	g.New(2)
	d.New()
	d.Shuffle()
	deal := d.Deal()
	for i := 0; i < g.Entries; i++ {
		var hand []map[string]int
		for j := 0; j < 2; j++ {
			hand = append(hand, deal())
		}
	}
}

func main() {
}

func question(q string) bool {
	result := true
	fmt.Print(q)
	answer := os.Stdin
	scanner := bufio.NewScanner(answer)
	fmt.Println(scanner.Text())
	for scanner.Scan() {
		i := scanner.Text()

		if i == "Y" || i == "y" {
			break
		} else if i == "N" || i == "n" {
			result = false
			break
		} else {
			fmt.Println("yかnで答えてください。")
			fmt.Print(q)
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return result
}
