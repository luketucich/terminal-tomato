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
	workTimeMins := int(workTime.Minutes())
	shortBreakTimeMins := int(shortBreakTime.Minutes())
	longBreakTimeMins := int(longBreakTime.Minutes())

	utils.PrintTomato(fmt.Sprintf(
		"Starting Pomodoro\n\n"+
			"  Work Time   : %2d mins\n"+
			"  Short Break : %2d mins\n"+
			"  Long Break  : %2d mins\n",
		workTimeMins, shortBreakTimeMins, longBreakTimeMins))

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
		time.Sleep(t.workTime)
		t.cycles++

		if t.cycles > 0 && t.cycles%4 == 0 {
			utils.Alert(fmt.Sprintf("Work session ended! Take a long break for %v", t.longBreakTime))
			time.Sleep(t.longBreakTime)
		} else {
			utils.Alert(fmt.Sprintf("Work session ended! Take a short break for %v", t.shortBreakTime))
			time.Sleep(t.shortBreakTime)
		}
	}
}
