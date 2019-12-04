package characters

import (
)

type Character struct {
	Name string
	TalkColor func(string, ...interface{})
}

func (c *Character) Say(message string) {
	c.TalkColor(message)
}