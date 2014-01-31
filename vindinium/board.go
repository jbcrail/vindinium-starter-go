package vindinium

import (
	"strings"
)

type Position struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type Board struct {
	Size  int    `json:"size"`
	Tiles string `json:"tiles"`
}

func (b Board) String() string {
	var s string
	buf := make([]byte, b.Size*2)
	reader := strings.NewReader(b.Tiles)
	for {
		if _, err := reader.Read(buf); err != nil {
			break
		}
		s += string(buf) + "\n"
	}
	return s
}
