package printer

import (
	"fmt"
	"os"
	"os/exec"
)

const (
	screenTitle = "Go implementation of 15 puzzle\n\n"
	gameLegend  = "\n\nUse arrow keys to move tiles\nPress 's' to shuffle field\nPress 'q', Esc or Ctrl+C to exit"
)

func Listen(ch <-chan string) {
	for {
		msg := <-ch

		clearCmd := exec.Command("clear")
		clearCmd.Stdout = os.Stdout
		clearCmd.Run()

		fmt.Printf("%s%s%s", screenTitle, msg, gameLegend)
	}
}
