package main

import (
	"fmt"

	"github.com/shgysd/go-poker/pkg/deck"
	"github.com/shgysd/go-poker/pkg/game"
)

var d deck.Deck
var g game.Game
var deal = d.Deal()

func init() {
	g.New(2)
	d.New()
	d.Shuffle()
	g.DealCard(deal)
}

func main() {
	fmt.Println("### Welcome to Texas hold 'em ###")
	next := g.NextRound(deal)
	for len(g.Players) != 1 {
		for g.Round != "Done" {
			for _, p := range g.Players {
				p.Action()
			}
			next()
			fmt.Println(g.Round)
		}
		g.Players = g.Players[:len(g.Players)-1]
	}
	fmt.Println("### Thank you for playing! ###")
}
