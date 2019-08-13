package deck

import (
	"math/rand"
	"time"
)

// Deck デッキの構造体
type Deck struct {
	cards []map[string]int
}

// New デッキを初期化
func (deck *Deck) New() *Deck {
	suits := []string{"spade", "club", "heart", "dia"}
	card := make([]map[string]int, 0, 52)
	for _, v := range suits {
		for i := 1; i <= 13; i++ {
			cards := make(map[string]int)
			cards[v] = i
			card = append(card, cards)
		}
	}
	deck.cards = card
	return deck
}

// Shuffle デッキを初期化
func (deck *Deck) Shuffle() *Deck {
	for i := range deck.cards {
		rand.Seed(time.Now().UnixNano())
		var j = rand.Float32() * float32(i+1)
		var swap = deck.cards[i]
		deck.cards[i] = deck.cards[int(j)]
		deck.cards[int(j)] = swap
	}
	return deck
}

func (deck *Deck) Deal() func() map[string]int {
	i := 0
	return func() map[string]int {
		i++
		return deck.cards[i]
	}
}
