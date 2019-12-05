package rooms

import (
	"fmt"

	"github.com/krefsar/bear-adventure/gameManager"
)

type Intro1 struct {
	Name string
	GM gameManager.GameManager
}

type Intro2 struct {
	Name string
	GM gameManager.GameManager
}

func (i Intro2) GetName() string {
	return i.Name
}

func (i Intro2) OnEnter() {
	i.PrintDescription()
}

func (i Intro2) Execute(command string) {
}

func (i Intro1) GetName() string {
	return i.Name
}

func (i Intro1) OnEnter() {
	i.PrintDescription()
}

func (i Intro1) Execute(command string) {
	if strings.Contains(strings.toLower(command), "ears") {
		gm.EnterRoom("Intro2")
	} else {
		fmt.Println("Doing that won't stop this annoying buzzing. If only you could hear it more clearly...")
	}
}

func (i Intro1) PrintDescription() {
	fmt.Println("Zzz... Zzz...")
	fmt.Println("It's your first day of hibernation, and you're happily snoozing in your Den. There's nothing more important to a bear than sleep, except maybe honey. And berries.")

	fmt.Println("\"Buzz! Buzz!\"")
	fmt.Println("An annoying buzzing sound starts ringing in your ear. Maybe it's best to ignore it.")
	fmt.Println()

	fmt.Println("\"Buzz! Buzz buzz buzz! Buzz buzz!\" You can't quite make out what the buzzing is, maybe if you focused your hearing...")
}

func (i Intro2) PrintDescription() {
	fmt.Println("You use your mighty BEAR EARS to focus in on the sound of the buzzing.")
	fmt.Println()
	fmt.Printf("\"Buzz buzz! Wake up, %s! You've got to wake up, we need your help!\"", i.GM.Player.Name)
	fmt.Println()
	fmt.Println("It's Buzzter, your Best Bee Friend! What's got him so worked up today? Can't he give you just a few more days of sleep?")
	fmt.Println()
	fmt.Println("\"Come on you lazy bear, I even brought you a treat if you get up. See if you can tell what it is without opening your eyes.\"")
	fmt.Println()
	fmt.Println("A present? Buzzter does know how to get your attention. Maybe there's a way to tell what he brought using your other senses.")
	fmt.Println()
}

func createIntroRoom1(gm gameManager.GameManager) Room {
	return Intro1{
		Name: "Intro 1",
		GM: gm,
	}
}

func createIntroRoom2(gm gameManager.GameManager) Room {
	return Intro2{
		Name: "Intro 2",
		GM: gm,
	}
}

func CreateIntroRooms(gm gameManager.GameManager) map[string]Room {
	return map[string]Room{
		"Intro1": createIntroRoom1(gm),
		"Intro2": createIntroRoom2(gm),
	}
}