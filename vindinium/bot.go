package vindinium

import (
	"math/rand"
	"time"
)

type Bot interface {
	Move(state State) string
}

func getRandomDirection() string {
	directions := []string{"Stay", "North", "South", "East", "West"}
	return directions[rand.Intn(len(directions))]
}

type RandomBot struct{}

func (b RandomBot) Move(state State) string {
	return getRandomDirection()
}

type FighterBot struct{}

func (b FighterBot) Move(state State) string {
	return getRandomDirection()
}

type SlowBot struct{}

func (b SlowBot) Move(state State) string {
	time.Sleep(time.Second * 2)
	return getRandomDirection()
}
