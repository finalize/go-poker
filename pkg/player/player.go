package player

import (
	"bufio"
	"fmt"
	"os"
)

// Player プレイヤーの情報
type Player struct {
	Name   string
	Hand   []map[string]int
	Status string
}

// SetHand 手札をセット
func (p *Player) SetHand(hand map[string]int) *Player {
	p.Hand = append(p.Hand, hand)
	return p
}

// Action アクションを分岐
func (p *Player) Action() {
	fmt.Println(p.Name + " is playing")
	fmt.Println("*** Check: 0 Call: 1 Bet: 2 Raise: 3 ***")
	fmt.Println("----- Your hand -----")
	for i := range p.Hand {
		for k, v := range p.Hand[i] {
			fmt.Printf("%s %d\n", k, v)
		}
	}
	fmt.Println("----- Your hand -----")
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Status:" + p.Status)
	for scanner.Scan() {
		t := scanner.Text()

		switch t {
		case "1":
			check(p)
		case "2":
			call()
		case "3":
			bet()
		case "4":
			raise()
		default:
			fmt.Println("Enter the number")
		}
		if p.Status != "" {
			break
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

func check(p *Player) {
	p.Status = "check"
	fmt.Println("check")
}

func call() {
	fmt.Println("call")

}

func bet() {
	fmt.Println("a")

}

func raise() {
	fmt.Println("a")

}
