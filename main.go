package main

import (
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
	for len(g.Players) != 1 {
		next := g.Next(deal)
		for g.Round != "Done" {
			next()
		}
		g.Players = g.Players[:len(g.Players)-1]
	}
}
