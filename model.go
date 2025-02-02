package tuispeak

import (
	"io"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
)

const (
	showBoards = iota
	showInput
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
	lists                     []list.Model
	choices                   []Container
	currentPositionsInLists   []int
	currentFocusedItemInLists []int
	speaker                   io.Writer
	focusedStyle              lipgloss.Style
	unfocusedStyle            lipgloss.Style
	currentlyFocusedList      int
	form                      *huh.Form
	currentlyFocusedBoard     int
}

type Container struct {
	Title     string
	Questions []string
}

func NewModel(containers []Container, speaker io.Writer, focusedStyle lipgloss.Style, unfocusedStyle lipgloss.Style) Model {
	lists := make([]list.Model, len(containers))
	defaultListPositions := make([]int, len(containers))
	for i, container := range containers {
		// TODO foreach container, ccreate a list of questions.
		lists[i] = list.New([]list.Item{}, list.NewDefaultDelegate(), 50, 20)
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
		lists:                     lists,
		choices:                   containers,
		speaker:                   speaker,
		focusedStyle:              focusedStyle,
		unfocusedStyle:            unfocusedStyle,
		currentPositionsInLists:   defaultListPositions,
		currentFocusedItemInLists: defaultListPositions,
		currentlyFocusedList:      0,
		form: huh.NewForm(
			huh.NewGroup(
				huh.NewInput().
					Title(`Say what?`).
					Key(`saythis`).
					Prompt(`> `)),
		),
	}
}

// Implement the init method.
func (m Model) Init() tea.Cmd {
	return m.form.Init()
}

// Update the model only.
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		h, v := m.focusedStyle.GetFrameSize()
		m.lists[m.currentlyFocusedList].SetSize(msg.Width-h, msg.Height-v)
	case tea.KeyMsg:
		switch msg.String() {
		case "tab", "ctrl+i":
			switch m.currentlyFocusedBoard {
			case showBoards:
				m.currentlyFocusedBoard = showInput
			case showInput:
				m.currentlyFocusedBoard = showBoards
			}
		case "s", "enter":
			switch m.currentlyFocusedBoard {
			case showBoards:
				m.speaker.Write([]byte(m.choices[m.currentlyFocusedList].Questions[m.currentPositionsInLists[m.currentlyFocusedList]]))
				m.speaker.Write([]byte("\n"))
			case showInput:
				s := m.form.GetString(`saythis`)
				m.speaker.Write([]byte(s))
				m.speaker.Write([]byte("\n"))
				m.currentlyFocusedBoard = showInput
			}
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
				m.currentFocusedItemInLists[m.currentlyFocusedList] = m.lists[m.currentlyFocusedList].Index()
				m.currentlyFocusedList--
				m.lists[m.currentlyFocusedList].Select(m.currentFocusedItemInLists[m.currentlyFocusedList])
			}
		case "l", "right":
			if m.currentlyFocusedList < len(m.lists)-1 {
				m.currentFocusedItemInLists[m.currentlyFocusedList] = m.lists[m.currentlyFocusedList].Index()
				m.currentlyFocusedList++
				m.lists[m.currentlyFocusedList].Select(m.currentFocusedItemInLists[m.currentlyFocusedList])
			}
		case "q", "x":
			return m, tea.Quit
		}
	}
	m.lists[m.currentlyFocusedList], cmd = m.lists[m.currentlyFocusedList].Update(msg)
	if m.currentlyFocusedBoard == showInput {
		return m.form.Update(msg)
	}

	return m, cmd
}

// Render the model.
func (m Model) View() string {
	switch m.currentlyFocusedBoard {
	case showBoards:
		views := make([]string, len(m.choices))
		for i, list := range m.lists {
			if i == m.currentlyFocusedList {
				views[i] = m.focusedStyle.Render(list.View())
				continue
			}

			views[i] = m.unfocusedStyle.Render(list.View())
		}
		return lipgloss.JoinHorizontal(lipgloss.Left, views...)
	case showInput:
		return m.form.View()
	}

	return ""
}
