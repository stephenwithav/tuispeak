package main

import (
	"bytes"
	"io"
	"log"
	"os"
	"os/exec"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/stephenwithav/tuispeak"
	"gopkg.in/yaml.v3"
)

func readBoardsConfig() *Boards {
	configFile, err := os.Open(`boards.yml`)
	if err != nil {
		return nil
	}
	defer configFile.Close()

	buf, err := io.ReadAll(configFile)
	if err != nil {
		return nil
	}

	dec := yaml.NewDecoder(bytes.NewReader(buf))
	var boards Boards
	err = dec.Decode(&boards)
	if err != nil {
		return nil
	}

	return &boards
}

func newContainersFromBoards(boards *Boards) []tuispeak.Container {
	b := make([]tuispeak.Container, len(boards.Boards))
	for i, board := range boards.Boards {
		b[i] = tuispeak.Container{
			Questions: board.Questions,
			Title:     board.Title,
		}
	}

	return b
}

func main() {
	speaker := exec.Command("spd-say", "-y", "Diogo", "--pipe-mode")
	speakerW, err := speaker.StdinPipe()
	if err != nil {
		log.Fatal(`Could not start spd-say.`)
	}
	defer speakerW.Close()
	if err := speaker.Start(); err != nil {
		log.Fatal(`Could not start spd-say.`)
	}
	style := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color(`#FFFFFF`)).
		Height(7).
		Width(30)

	boards := readBoardsConfig()
	if boards == nil {
		log.Fatal(`Unable to read boards.yml`)
	}

	containers := newContainersFromBoards(boards)
	if len(containers) == 0 {
		log.Fatal(`No containers defined in boards.yml`)
	}

	model := tuispeak.NewModel(
		containers,
		speakerW,
		style,
	)
	if _, err := tea.NewProgram(model).Run(); err != nil {
		log.Fatalf(`Error running program: %s`, err.Error())
		os.Exit(1)
	}
}
