package config

var PreviousState string
var CurrentState string

func SetPreviousState(value string) {
	PreviousState = value
}

func SetCurrentState(value string) {
	CurrentState = value
}
