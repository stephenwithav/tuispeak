package tuispeak

import (
	"fmt"
	"io"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type QueryModel struct {
	input    textinput.Model
	speaker  io.Writer
	returnTo tea.Model
}

func NewQueryModel(query string, speaker io.Writer, returnTo tea.Model) *QueryModel {
	ti := textinput.New()
	ti.Placeholder = `query`
	ti.Focus()

	return &QueryModel{
		input:    ti,
		speaker:  speaker,
		returnTo: returnTo,
	}
}

func (m QueryModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m QueryModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			m.speaker.Write([]byte(m.input.Value()))
			m.speaker.Write([]byte("\n"))
			return m.returnTo, nil
		}
	}

	m.input, cmd = m.input.Update(msg)
	return m, cmd
}

func (m QueryModel) View() string {
	return fmt.Sprintf("Say what?\n%s", m.input.View())
}
