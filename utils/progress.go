package utils

import (
	"fmt"
	"strings"
	"time"
)

func printStaticProgressBar(elapsed, duration time.Duration) {
	percentComplete := float64(elapsed) / float64(duration)
	if percentComplete > 1 {
		percentComplete = 1
	}

	barLength := 30
	amountFilled := int(percentComplete * float64(barLength))
	bar := strings.Repeat("▮", amountFilled) + strings.Repeat("▯", barLength-amountFilled)

	remaining := duration - elapsed
	if remaining < 0 {
		remaining = 0
	}

	// Round up seconds to avoid jumpy timer
	totalSeconds := int((remaining + time.Second - 1).Seconds())
	mins := totalSeconds / 60
	secs := totalSeconds % 60

	fmt.Printf("\r[%s] %02d:%02d remaining", bar, mins, secs)
}

func PrintAnimatedProgressBar(duration time.Duration) {
	start := time.Now()
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	printStaticProgressBar(0, duration)

	for now := range ticker.C {
		elapsed := now.Sub(start)
		if elapsed >= duration {
			break
		}
		printStaticProgressBar(elapsed, duration)
	}

	bar := strings.Repeat("▮", 30)
	fmt.Printf("\r[%s] 00:00 remaining\n\n", bar)
}
