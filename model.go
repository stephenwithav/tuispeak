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
	lists    []list.Model
	choices  []Container
	cursorAt int
	speaker  io.Writer
	style    lipgloss.Style
	focused  int
}

type Container struct {
	Title     string
	Questions []string
}

func NewModel(containers []Container, speaker io.Writer, style lipgloss.Style) Model {
	choices := make([]Container, len(containers))
	lists := make([]list.Model, len(containers))
	for i, container := range containers {
		// TODO foreach container, ccreate a list of questions.
		lists[i] = list.New([]list.Item{}, list.NewDefaultDelegate(), 30, 20)
		lists[i].SetShowHelp(false)
		lists[i].Title = container.Title

		// TODO create slice of list.Item, set them.
		listOfQuestions := make([]list.Item, len(container.Questions))
		for j, question := range container.Questions {
			listOfQuestions[j] = Question{
				q: question,
			}
		}
		lists[i].SetItems(listOfQuestions)
	}

	return Model{
		lists:    lists,
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
	m.lists[m.focused], cmd = m.lists[m.focused].Update(msg)

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		h, v := m.style.GetFrameSize()
		m.lists[m.focused].SetSize(msg.Width-h, msg.Height-v)
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
	views := make([]string, len(m.choices))
	for i, list := range m.lists {
		views[i] = list.View()
	}
	return lipgloss.JoinHorizontal(lipgloss.Left, views...)
}
