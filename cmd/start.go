package cmd

import (
	"github.com/luketucich/terminal-tomato/timer"
	"github.com/luketucich/terminal-tomato/utils"
	"github.com/spf13/cobra"
	"time"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start a Pomodoro timer",
	Run: func(cmd *cobra.Command, args []string) {
		workTime := utils.PromptInt("How many minutes do you want to work?")
		shortBreakTime := utils.PromptInt("How many minutes for a short break?")
		longBreakTime := utils.PromptInt("How many minutes for a long break?")

		t := timer.NewTimer(
			time.Minute*time.Duration(workTime),
			time.Minute*time.Duration(shortBreakTime),
			time.Minute*time.Duration(longBreakTime),
		)

		t.Start()
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
