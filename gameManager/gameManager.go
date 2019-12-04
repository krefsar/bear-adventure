package gameManager

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"

	"github.com/krefsar/bear-adventure/characters"
	"github.com/krefsar/bear-adventure/player"
)

type GameManager struct {
	Characters []*characters.Character
	GameOver bool
	Player player.Player

	// prompt colors
	InfoColor func(string, ...interface{})
	NarrationColor func(string, ...interface{})
	PromptColor color.Attribute
}

func (g *GameManager) Info(message string, more ...interface{}) {
	g.InfoColor(message, more...)
}

func (g *GameManager) Narrate(message string, more ...interface{}) {
	g.NarrationColor(message, more...)
}

func (g *GameManager) Prompt(message string, more ...interface{}) (output string) {
	reader := bufio.NewReader(os.Stdin)

	promptColor := color.New(g.PromptColor)
	promptColor.Printf(fmt.Sprintf("%s ", message), more...)

	output, _ = reader.ReadString('\n')
	output = strings.TrimSuffix(output, "\n")
	return output
}

func Initialize() *GameManager {
	npcs := createCharacters()
	player := player.Player{}

	return &GameManager{
		Characters: npcs,
		GameOver: false,
		Player: player,

		InfoColor: color.Blue,
		NarrationColor: color.White,
		PromptColor: color.FgGreen,
	}
}

func createCharacters() []*characters.Character {
	var npcs []*characters.Character

	npcs = append(npcs, characters.InitBuzzter())
	return npcs
}