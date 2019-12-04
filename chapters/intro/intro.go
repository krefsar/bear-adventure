package chapters

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/krefsar/bear-adventure/characters"
	"github.com/krefsar/bear-adventure/player"
	"github.com/krefsar/bear-adventure/powers"
)

func RunIntro(player *player.Player) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("You're enjoying a nice snooze in your den, dreaming of honey and berries.")
	Buzzter := characters.InitBuzzter()
	Buzzter.Say("\"Buzz! Buzz!\"")
	fmt.Println("An annoying buzzing sound starts ringing in your ear. Maybe it's best to ignore it.")
	fmt.Println()

	fmt.Println("\"Buzz! Buzz buzz buzz! Buzz buzz!\" You can't quite make out what the buzzing is, maybe if you focused your hearing...")
	fmt.Println()

	hearingPower := powers.Power{
		Name: "Bear Ears",
	}
	player.GainPower(&hearingPower)

	var playerPowers []string
	for _, power := range player.Powers {
		playerPowers = append(playerPowers, strings.ToUpper(power.Name))
	}
	fmt.Printf("Available powers: %s\n", strings.Join(playerPowers, ", "))

	hearingPowerUsed := false
	for hearingPowerUsed != true {
		fmt.Print("What do you do? ")
		playerAction, _ := reader.ReadString('\n')
		playerAction = strings.TrimSuffix(playerAction, "\n")

		if strings.Contains(strings.ToLower(playerAction), "ears") {
			hearingPowerUsed = true
		} else {
			fmt.Println("Doing that won't stop this annoying buzzing. If only you could hear it more clearly...")
		}

		fmt.Println()
	}

	fmt.Println("You use your mighty BEAR EARS to focus in on the sound of the buzzing.")
	fmt.Println()
	fmt.Printf("\"Buzz buzz! Wake up, %s! You've got to wake up, we need your help!\"", player.Name)
	fmt.Println()
	fmt.Println("It's Buzzter, your Best Bee Friend! What's got him so worked up today? Can't he give you just a few more days of sleep?")
	fmt.Println()
	fmt.Println("\"Come on you lazy bear, I even brought you a treat if you get up. See if you can tell what it is without opening your eyes.\"")
	fmt.Println()
	fmt.Println("A present? Buzzter does know how to get your attention. Maybe there's a way to tell what he brought using your other senses.")
	fmt.Println()

	smellingPower := powers.Power{
		Name: "Bear Nose",
	}
	player.GainPower(&smellingPower)

	playerPowers = nil 
	for _, power := range player.Powers {
		playerPowers = append(playerPowers, strings.ToUpper(power.Name))
	}
	fmt.Printf("Available powers: %s\n", strings.Join(playerPowers, ", "))

	smellingPowerUsed := false
	for smellingPowerUsed != true {
		fmt.Print("What do you do? ")
		playerAction, _ := reader.ReadString('\n')
		playerAction = strings.TrimSuffix(playerAction, "\n")

		if strings.Contains(strings.ToLower(playerAction), "nose") {
			smellingPowerUsed = true
		} else {
			fmt.Println("Doing that won't help you figure out what Buzzter brought. Is there a way to sense it without opening your eyes?")
		}

		fmt.Println()
	}

	fmt.Println("Using your powerful BEAR NOSE, you sniff deeply. Your DEN is full of many scents, but is anything different?")
	fmt.Println("Then you smell it, the unmistakable scent... of HONEY! Buzzter knows you all too well. You smack your lips in appreciation.")
	fmt.Println()
	fmt.Printf("\"Ha, so you've smelled your present, %s? Well there's plenty more for you if you help us out.\"\n", player.Name)
	fmt.Println("Hmm, you don't know if doing work is worth missing out on sleep, but you do really love honey...")
	fmt.Println()
	fmt.Printf("\"The BEES have been buzzing about some new flower they're searching for, and it sounds like they need your help to find it, %s.\n", player.Name)
	fmt.Println("Meet me at THE HIVE and we'll find out more together. They're offering some serious HONEY!\"")
	fmt.Println()
	fmt.Println("You couldn't care less about a flower, but you're going to need some more HONEY if you're going to make it through the upcoming winter.")
	fmt.Println()
	fmt.Printf("\"See you there, %s! Buzz buzz!\"\n", player.Name)
	fmt.Println()

}