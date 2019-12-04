package main

import (
	"fmt"

	"github.com/krefsar/bear-adventure/chapters/intro"
	"github.com/krefsar/bear-adventure/gameManager"
)

func main() {
	gm := gameManager.Initialize()

	gm.Info("Welcome to Bear Adventure!")
	gm.Info("A game by Evan Martinez")
	fmt.Println()

	player := gm.Player

	nameAccepted := false
	for nameAccepted != true {
		bearName := gm.Prompt("You are a bear. What is your name?")

		confirmName := gm.Prompt("Your name is %s (y/n)?", bearName)

		if confirmName == "y" {
			nameAccepted = true
			player.Name = bearName
		}

		fmt.Println()
	}

	fmt.Printf("Great! You are a bear named %s, and your adventure is about to begin!\n", player.Name)
	fmt.Println()

	chapters.RunIntro(&player)

	fmt.Println("GAME OVER")
	fmt.Println("Thanks for playing")
}