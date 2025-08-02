package utils

import "fmt"

const (
	Red   = "\033[31m"
	Reset = "\033[0m"
)

func PrintTomato(msg string) {
	fmt.Println(Red + "terminal-tomato 🍅  " + Reset + msg)
}
