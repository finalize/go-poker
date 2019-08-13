package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/shgysd/go-poker/pkg/deck"
)

var d deck.Deck

func init() {
	d.New()
	d.Shuffle()
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
