package rooms

type Room interface {
	GetName() string
	OnEnter()
	PrintDescription()
	Execute(string)
}