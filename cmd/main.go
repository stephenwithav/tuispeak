package main

import (
	"log"
	"os"
	"os/exec"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/stephenwithav/tuispeak"
)

func main() {
	speaker := exec.Command("espeak", "-f", "/dev/stdin")
	speakerW, err := speaker.StdinPipe()
	if err != nil {
		log.Fatal(`Could not start espeak.`)
	}
	defer speakerW.Close()
	style := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color(`#FFFFFF`)).
		Height(7).
		Width(20)

	model := tuispeak.NewModel(
		[]string{"Hello", "You", `There`},
		speakerW,
		style,
	)
	if _, err := tea.NewProgram(model).Run(); err != nil {
		log.Fatalf(`Error running program: %s`, err.Error())
		os.Exit(1)
	}
}
