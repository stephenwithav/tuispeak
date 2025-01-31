package tuispeak

import (
	"io"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Question struct {
	q string
}

func (q Question) FilterValue() string {
	return q.q
}

func (q Question) Title() string {
	return q.q
}

func (q Question) Description() string {
	return ""
}

type Model struct {
	list     list.Model
	choices  []Question
	cursorAt int
	speaker  io.Writer
	style    lipgloss.Style
}

func NewModel(questions []string, speaker io.Writer, style lipgloss.Style) Model {
	choices := make([]Question, len(questions))
	for i, val := range questions {
		choices[i] = Question{
			q: val,
		}
	}
	l := list.New([]list.Item{}, list.NewDefaultDelegate(), 30, 20)
	l.Title = "Questions"
	l.SetItems([]list.Item{
		choices[0],
		choices[1],
	})

	return Model{
		list:     l,
		choices:  choices,
		speaker:  speaker,
		style:    style,
		cursorAt: 0,
	}
}

// Implement the init method.
func (m Model) Init() tea.Cmd {
	return nil
}

// Update the model only.
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)

	switch msg := msg.(type) {
	// case tea.WindowSizeMsg:
	// 	h, v := m.style.GetFrameSize()
	// 	m.list.SetSize(msg.Width-h, msg.Height-v)
	case tea.KeyMsg:
		switch msg.String() {
		case "s", "enter":
			m.speaker.Write([]byte("Chosen"))
			m.speaker.Write([]byte{4})
		case "j":
			m.cursorAt++
		case "k":
			m.cursorAt--
		case "q", "x":
			return m, tea.Quit
		}
	}

	return m, cmd
}

// Render the model.
func (m Model) View() string {
	return m.style.Render(m.list.View())
}
