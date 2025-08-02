package utils

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func PromptInt(message string) int {
	reader := bufio.NewReader(os.Stdin)

	for {
		PrintTomato(message)

		// Read line + remove extra space
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		// Try to convert input to number
		num, err := strconv.Atoi(input)
		if err != nil || num <= 0 {
			PrintTomato("Please enter a positive integer.")
			continue // Repeat until valid
		}

		return num
	}
}
