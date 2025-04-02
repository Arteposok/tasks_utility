package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
    tasks []string
	chosen int
	add_mode bool
	input textinput.Model

}
func initialModel() model {
	ti := textinput.New()
	ti.Width=30
	ti.CharLimit=30
	ti.Placeholder="code?"
	ti.Focus()
	return model{
		tasks: []string{"task"},
		chosen: 0,
		add_mode: false,
		input: ti,
	}
}
func (m model) Init() tea.Cmd {
    return textinput.Blink
}
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
    switch msg := msg.(type) {
    case tea.KeyMsg:
        switch msg.String() {
			case "ctrl+c":
				return m, tea.Quit
			case "up", "k":
				if m.chosen > 0 {
					m.chosen--
				}
			case "down", "j":
				if m.chosen < len(m.tasks)-1 {
					m.chosen++
				}
			case "enter":
				m.tasks = append(m.tasks, m.input.Value())
				m.input.Reset()
			case "ctrl+a":
				m.add_mode=true
			case "tab":
				m.tasks=append(m.tasks[:m.chosen], m.tasks[m.chosen+1:]...)
				if m.chosen>len(m.tasks)-1{
					m.chosen-=1
				}
        }
		m.input, cmd = m.input.Update(msg)
    }
	return m, cmd
}
func (m model) View() string {
    s := "\n\tTASKS\n"
	s += "________________________________________________\n\n"
	s +=  "]|"+m.input.View()+"\n\n"
    for i, choice := range m.tasks {
        cursor := " "
        if m.chosen==i{
			cursor="x"
		}
        s += fmt.Sprintf("%v [%s] - - - - - %s [tab X]\n", i, cursor,choice)
    }
    s += "\n\tCTRL+C to exit\n"
    return s
}
func main() {
    p := tea.NewProgram(initialModel())
    if _, err := p.Run(); err != nil {
        fmt.Printf("Alas, there's been an error: %v", err)
        os.Exit(1)
    }
}