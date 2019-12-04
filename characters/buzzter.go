package characters

import (
	"github.com/fatih/color"
)

func InitBuzzter() *Character {
	return &Character{
		Name: "Buzzter",
		TalkColor: color.Yellow,
	}
}