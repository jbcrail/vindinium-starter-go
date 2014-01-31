package vindinium

type Game struct {
	Id       string `json:"id"`
	Turn     int    `json:"turn"`
	MaxTurns int    `json:"maxTurns"`
	Heroes   []Hero `json:"heroes"`
	Board    Board  `json:"board"`
	Finished bool   `json:"finished"`
}
