package timer

import (
	"fmt"
	"github.com/luketucich/terminal-tomato/utils"
	"time"
)

type Timer struct {
	workTime       time.Duration
	shortBreakTime time.Duration
	longBreakTime  time.Duration
	cycles         int
}

func NewTimer(workTime time.Duration, shortBreakTime time.Duration, longBreakTime time.Duration) *Timer {
	utils.PrintTomato(fmt.Sprintf(
		"Starting Pomodoro\n\n"+
			"  Work Time   : %v\n"+
			"  Short Break : %v\n"+
			"  Long Break  : %v\n",
		workTime, shortBreakTime, longBreakTime))

	return &Timer{
		workTime:       workTime,
		shortBreakTime: shortBreakTime,
		longBreakTime:  longBreakTime,
		cycles:         0,
	}
}

func (t *Timer) Start() {
	for {
		utils.Alert(fmt.Sprintf("Work session started! Focus for %v", t.workTime))
		utils.PrintTomato("Work session started!")
		utils.PrintAnimatedProgressBar(t.workTime)

		t.cycles++

		if t.cycles > 0 && t.cycles%4 == 0 {
			utils.Alert(fmt.Sprintf("Work session ended! Take a long break for %v", t.longBreakTime))
			utils.PrintTomato("Long break started! Relax and recharge.")
			utils.PrintAnimatedProgressBar(t.longBreakTime)
		} else {
			utils.Alert(fmt.Sprintf("Work session ended! Take a short break for %v", t.shortBreakTime))
			utils.PrintTomato("Short break started! Take a quick rest.")
			utils.PrintAnimatedProgressBar(t.shortBreakTime)
		}
	}
}
