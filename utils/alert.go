package utils

import "github.com/gen2brain/beeep"

func Alert(message string) {
	go func(msg string) {
		_ = beeep.Beep(beeep.DefaultFreq, beeep.DefaultDuration)
		err := beeep.Alert("terminal-tomato 🍅", msg, "")
		if err != nil {
			PrintTomato("Unable to send notification: " + msg)
		}
	}(message)
}
