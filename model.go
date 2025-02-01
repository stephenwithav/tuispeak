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
	lists                   []list.Model
	choices                 []Container
	currentPositionsInLists []int
	speaker                 io.Writer
	style                   lipgloss.Style
	currentlyFocusedList    int
}

type Container struct {
	Title     string
	Questions []string
}

func NewModel(containers []Container, speaker io.Writer, style lipgloss.Style) Model {
	lists := make([]list.Model, len(containers))
	defaultListPositions := make([]int, len(containers))
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
		defaultListPositions[0] = 0
	}

	return Model{
		lists:                   lists,
		choices:                 containers,
		speaker:                 speaker,
		style:                   style,
		currentPositionsInLists: defaultListPositions,
		currentlyFocusedList:    0,
	}
}

// Implement the init method.
func (m Model) Init() tea.Cmd {
	return nil
}

// Update the model only.
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		h, v := m.style.GetFrameSize()
		m.lists[m.currentlyFocusedList].SetSize(msg.Width-h, msg.Height-v)
	case tea.KeyMsg:
		switch msg.String() {
		case "s", "enter":
			m.speaker.Write([]byte(m.choices[m.currentlyFocusedList].Questions[m.currentPositionsInLists[m.currentlyFocusedList]]))
			m.speaker.Write([]byte("\n"))
		case "j", "up":
			if m.currentPositionsInLists[m.currentlyFocusedList] < len(m.choices[0].Questions)-1 {
				m.currentPositionsInLists[m.currentlyFocusedList]++
			}
		case "k", "down":
			if m.currentPositionsInLists[m.currentlyFocusedList] > 0 {
				m.currentPositionsInLists[m.currentlyFocusedList]--
			}
		case "h", "left":
			if m.currentlyFocusedList > 0 {
				m.currentlyFocusedList--
			}
		case "l", "right":
			if m.currentlyFocusedList < len(m.lists)-1 {
				m.currentlyFocusedList++
			}
		case "q", "x":
			return m, tea.Quit
		}
	}
	m.lists[m.currentlyFocusedList], cmd = m.lists[m.currentlyFocusedList].Update(msg)

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
