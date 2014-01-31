package vindinium

type Bot interface {
	Move(state State) string
}

type RandomBot struct{}

func (b RandomBot) Move(state State) string {
	return "Stay"
}

type FighterBot struct{}

func (b FighterBot) Move(state State) string {
	return "Stay"
}

type SlowBot struct{}

func (b SlowBot) Move(state State) string {
	return "Stay"
}
