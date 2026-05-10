package main

import "fmt"

type Player struct {
	score int
	name  string
}

func (p *Player) SetScore(s int) {
	old := p.score
	p.score = s
	fmt.Printf("score changed: %d → %d\n", old, p.score)
}

func (p *Player) Score() int {
	return p.score
}
