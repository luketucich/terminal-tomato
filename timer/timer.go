package timer

import (
	"bufio"
	"fmt"
	"github.com/luketucich/terminal-tomato/utils"
	"os"
	"strings"
	"time"
)

type Timer struct {
	workTime       time.Duration
	shortBreakTime time.Duration
	longBreakTime  time.Duration
	cycles         int
	controlChan    chan string
	sessionDone    chan bool
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
		controlChan:    make(chan string),
		sessionDone:    make(chan bool),
	}
}

func (t *Timer) listenForCommands() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		command := strings.TrimSpace(strings.ToLower(scanner.Text()))
		if command != "" {
			t.controlChan <- command
		}
	}
}

func (t *Timer) displayProgressBar(duration time.Duration) {
	utils.PrintAnimatedProgressBar(duration)
	t.sessionDone <- true
}

func (t *Timer) workSession() {
	utils.Alert(fmt.Sprintf("Work session started! Focus for %v", t.workTime))
	utils.PrintTomato("Work session started!")

	go t.displayProgressBar(t.workTime)

	t.cycles++
}

func (t *Timer) breakSession() {
	if t.cycles > 0 && t.cycles%4 == 0 {
		utils.Alert(fmt.Sprintf("Work session ended! Take a long break for %v", t.longBreakTime))
		utils.PrintTomato("Long break started! Relax and recharge.")
		go t.displayProgressBar(t.longBreakTime)
	} else {
		utils.Alert(fmt.Sprintf("Work session ended! Take a short break for %v", t.shortBreakTime))
		utils.PrintTomato("Short break started! Take a quick rest.")
		go t.displayProgressBar(t.shortBreakTime)
	}

}

func (t *Timer) TimerLoop() {
	go t.listenForCommands()

	isWorkSession := true
	t.workSession()

	for {
		select {

		// Handle command input
		case cmd := <-t.controlChan:
			switch cmd {

			case "stop":
				utils.PrintTomato("Timer stopped. Goodbye!")
				return
			case "skip":
				utils.PrintTomato("Skipping current session...")
				go func() { t.sessionDone <- true }()
			default:
				utils.PrintTomato(fmt.Sprintf("Unknown command: %s. Available: stop, skip", cmd))
			}

		// Handle work + break sessions
		case <-t.sessionDone:

			if isWorkSession {
				t.breakSession()
				isWorkSession = false
			} else {
				t.workSession()
				isWorkSession = true
			}
		}
	}
}
