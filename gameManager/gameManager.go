package gameManager

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/krefsar/bear-adventure/player"
)

type GameManager struct {
	CurrentLocation *Location
	GameOver bool
	Player player.Player
}

func Initialize() *GameManager {
	player := player.Player{}

	gm := &GameManager{
		Player: player,
	}

	return gm
}

func (g *GameManager) EndGame() {
	g.GameOver = true
}

func (g *GameManager) EnterLocation(locationName string) {
	newLocation := locationMap[locationName]
	g.CurrentLocation = newLocation
	newLocation.PrintDescription()
}

func (g *GameManager) ExecuteCommand(playerCommand string) {
	tokens := strings.Fields(strings.ToLower(playerCommand))

	command := tokens[0]

	switch command {
	case "go":
		if len(tokens) > 1 {
			direction := tokens[1]
			nextLocation := g.GoDirection(direction)

			if nextLocation != "" {
				g.EnterLocation(nextLocation)
			}
		} else {
			fmt.Println("Go in which direction?")
		}
	case "look":
		g.CurrentLocation.PrintDescription()
	case "quit":
		g.EndGame()
	default:
		fmt.Println("Sorry, that command wasn't recognized.")
	}
}

func (g *GameManager) GoDirection(direction string) (nextLocation string) {
	if isValidDirection(direction) {
		switch direction {
			case "n":
				direction = "north"
			case "s":
				direction = "south"
			case "e":
				direction = "east"
			case "w":
				direction = "west"
		}

		return g.CurrentLocation.GoDirection(direction)
	} else {
		fmt.Printf("%s is not an available direction. Available directions are North, South, East, and West.\n", strings.Title(direction))
		return ""
	}
}

func (g *GameManager) Prompt(message string, more ...interface{}) (output string) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf(message + " ", more...)

	output, _ = reader.ReadString('\n')
	output = strings.TrimSuffix(output, "\n")
	return output
}

func (g *GameManager) Welcome() {
	fmt.Println("Welcome to Bear Adventure!")
	fmt.Println("A game by Evan Martinez")
	fmt.Println()
}

type PrintDescriptionFunc func()

type Location struct {
	ConnectionMap map[string]string
	Name string
	PrintDescription PrintDescriptionFunc
}

func (l *Location) GoDirection(direction string) (nextLocation string) {
	connectedLocation, present := l.ConnectionMap[direction]

	if !present {
		fmt.Printf("Cannot go %s from here.\n", strings.Title(direction))
		return ""
	} else {
		return connectedLocation
	}
}

var Intro1 = &Location{
	Name: "Intro 1",
	PrintDescription: func() {
		fmt.Println("This is the Intro 1 Location")
	},
	ConnectionMap: map[string]string{
		"south": "Intro2",
	},
}

var Intro2 = &Location{
	Name: "Intro 2",
	PrintDescription: func() {
		fmt.Println("This is the Intro 2 Location")
	},
	ConnectionMap: map[string]string{
		"north": "Intro1",
	},
}

var locationMap = map[string]*Location{
	"Intro1": Intro1,
	"Intro2": Intro2,
}

func isValidDirection(direction string) bool {
	switch direction {
	case "n", "north", "s", "south", "e", "east", "w", "west":
		return true
	default:
		return false
	}
}