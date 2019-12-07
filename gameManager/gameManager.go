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
	Player *player.Player
}

var currentPlayer = &player.Player{}

func Initialize() *GameManager {
	gm := &GameManager{
		Player: currentPlayer,
	}

	return gm
}

func (g *GameManager) EndGame() {
	g.GameOver = true
}

func (g *GameManager) EnterLocation(locationName string) {
	newLocation := locationMap[locationName]
	g.CurrentLocation = newLocation

	fmt.Printf("%s\n", strings.ToUpper(newLocation.Name))
	g.Look()
}

func (g *GameManager) Look() {
	currentLocation := g.CurrentLocation
	currentLocation.PrintDescription()
	currentLocation.PrintItems()
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
		g.Look()
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

		return g.CurrentLocation.GetLocationInDirection(direction)
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

type Item struct {
	DescriptionName string
	Name string
}

type Location struct {
	ConnectionMap map[string]string
	Items []Item
	Name string
	PrintDescription PrintDescriptionFunc
}

func (l *Location) GetLocationInDirection(direction string) (nextLocation string) {
	connectedLocation, present := l.ConnectionMap[direction]

	if !present {
		fmt.Printf("Cannot go %s from here.\n", strings.Title(direction))
		return ""
	} else {
		return connectedLocation
	}
}

func (l *Location) PrintItems() {
	if len(l.Items) > 0 {
		// map to description names
		var descriptionNames []string
		for _, item := range l.Items {
			descriptionNames = append(descriptionNames, item.DescriptionName)
		}

		// change the last element to "and <item>"
		if len(descriptionNames) > 1 {
			descriptionNames[len(descriptionNames) - 1] = fmt.Sprintf("and %s", descriptionNames[len(descriptionNames) - 1])
		}

		// join the description names
		var joinedString string
		if len(descriptionNames) == 2 {
			joinedString = strings.Join(descriptionNames, " ")
		} else {
			joinedString = strings.Join(descriptionNames, ", ")
		}

		fmt.Printf("You can see %s.\n", joinedString)
	}
}

var BedItem = Item{
	Name: "Bed",
	DescriptionName: "a BED",
}

var JarItem = Item{
	Name: "Jar of Honey",
	DescriptionName: "a JAR OF HONEY",
}

var MeatItem = Item{
	Name: "Meat Scraps",
	DescriptionName: "some SCRAPS OF MEAT",
}

var CoolerItem = Item{
	Name: "Cooler",
	DescriptionName: "a COOLER",
}

var TentItem = Item{
	Name: "Shredded Tent",
	DescriptionName: "a SHREDDED TENT",
}

var DenLocation = &Location{
	Name: "Bear Den",
	ConnectionMap: map[string]string{
		"south": "CampLocation",
	},
	Items: []Item{BedItem, JarItem, MeatItem},
	PrintDescription: func() {
		fmt.Println("Your den is perfect for a bear. It's where you like to hang out when you need some alone time, especially when the other forest animals start to get to you.")
	},
}

var CampLocation = &Location{
	Name: "Camp",
	ConnectionMap: map[string]string{
		"north": "DenLocation",
	},
	Items: []Item{CoolerItem, TentItem},
	PrintDescription: func() {
		fmt.Println("The Camp is littered with things the humans left behind. It's messy, but that's because they left in such a hurry!")
	},
}

var locationMap = map[string]*Location{
	"CampLocation": CampLocation,
	"DenLocation": DenLocation,
}

func isValidDirection(direction string) bool {
	switch direction {
	case "n", "north", "s", "south", "e", "east", "w", "west":
		return true
	default:
		return false
	}
}