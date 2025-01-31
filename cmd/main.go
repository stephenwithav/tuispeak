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

	// Create 2/3 Models.  What contains them?
	// The model.  Ugh.
	// Perhaps a BaseModel with a container holding them?
	// Use lipgloss.JoinHorizontal(lipgloss.Left, containers.View())
	// Create a divisor container.  Each is 1/3 width.
	// Initialize list on tea.WindowSizeMsg type?  Nah, before.
	model := tuispeak.NewModel(
		[]tuispeak.Container{
			{
				Questions: []string{"Hello", "You", `There`},
				Title:     `To Dos`,
			},
			{
				Questions: []string{"Hello", "You", `There`},
				Title:     `In Progress`,
			},
		},
		speakerW,
		style,
	)
	if _, err := tea.NewProgram(model).Run(); err != nil {
		log.Fatalf(`Error running program: %s`, err.Error())
		os.Exit(1)
	}
}
