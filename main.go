package main

import (
	"fmt"

	"github.com/krefsar/bear-adventure/gameManager"
)

func main() {
	gm := gameManager.Initialize()

	gm.Welcome()

	nameAccepted := false
	for nameAccepted != true {
		bearName := gm.Prompt("You are a bear. What is your name?")
		confirmName := gm.Prompt("Your name is %s (y/n)?", bearName)

		if confirmName == "y" {
			nameAccepted = true
			gm.Player.Name = bearName
		}

		fmt.Println()
	}

	fmt.Printf("Great! You are a bear named %s, and your adventure is about to begin!\n", gm.Player.Name)
	fmt.Println()

	gm.EnterLocation("DenLocation")

	for !gm.GameOver {
		fmt.Println()
		playerCommand := gm.Prompt(">")
		gm.ExecuteCommand(playerCommand)
	}

	fmt.Println()
	fmt.Println("GAME OVER")
	fmt.Println("Thanks for playing")
}