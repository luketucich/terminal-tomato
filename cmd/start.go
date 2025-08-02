package cmd

import (
	"fmt"
	"github.com/luketucich/terminal-tomato/utils"
	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start a Pomodoro timer",
	Run: func(cmd *cobra.Command, args []string) {
		workTime := utils.PromptInt("How many minutes do you want to work?")
		shortBreakTime := utils.PromptInt("How many minutes for a short break?")
		longBreakTime := utils.PromptInt("How many minutes for a long break?")

		utils.PrintTomato(fmt.Sprintf(
			"Starting Pomodoro\n\n"+
				"  Work Time   : %2d mins\n"+
				"  Short Break : %2d mins\n"+
				"  Long Break  : %2d mins\n",
			workTime, shortBreakTime, longBreakTime))
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
