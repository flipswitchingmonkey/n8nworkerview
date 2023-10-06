package main

import (
	"slices"
	"time"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd
	switch msg := msg.(type) {
		case RedisWorkerResponse:
			m.lastUpdated[msg.WorkerId] = time.Now()
			switch msg.Command {
			case "getId":
				if !slices.Contains(m.workerNames, msg.WorkerId) {
					m.workerNames = append(m.workerNames, msg.WorkerId)
				}
			case "getStatus":
				if !slices.Contains(m.workerNames, msg.WorkerId) {
					m.workerNames = append(m.workerNames, msg.WorkerId)
				}
				m.workers[msg.WorkerId] = msg.Payload
			}
		case tea.WindowSizeMsg:
			verticalMarginHeight := 2
			if !m.ready {
				// Since this program is using the full size of the viewport we
				// need to wait until we've received the window dimensions before
				// we can initialize the viewport. The initial dimensions come in
				// quickly, though asynchronously, which is why we wait for them
				// here.
				m.viewport = viewport.New(msg.Width, msg.Height-verticalMarginHeight)
				m.viewport.YPosition = 0
				m.viewport.HighPerformanceRendering = false
				m.viewport.SetContent(RenderWorkerDetails(m))
				m.ready = true
			} else {
				m.viewport.Width = msg.Width
				m.viewport.Height = msg.Height - verticalMarginHeight
			}
		case tea.KeyMsg:
			selectedWorkerName, _ := m.GetSelectedRowName()
			m.workerTable, cmd = m.workerTable.Update(msg)
			cmds = append(cmds, cmd)
			switch msg.String() {
				case "ctrl+c", "q", tea.KeyEsc.String():
					return m, tea.Quit
				case "+":
					IncreaseInterval()
				case "-":
					DecreaseInterval()
				case "s", "up", "k", "down", "j":
					if selectedWorkerName != "" {
						SendGetStatusByName(selectedWorkerName)
					}
				case tea.KeyTab.String():
					if m.workerTable.Focused() {
						m.workerTable.Blur()
					} else {
						m.workerTable.Focus()
					}
				case "enter", " ":
					m.showDetails = !m.showDetails
			}
			
	}
	m.CheckStaleWorkers()
	m.workerTable.SetRows(m.GetWorkerNamesAsRows())
	m.jobsTable.SetRows(m.GetSelectedWorkerJobSummariesAsRows())
	// Handle keyboard and mouse events in the viewport
	m.viewport.SetContent(RenderWorkerDetails(m))
	if m.workerTable.Focused() {
		m.viewport, cmd = m.viewport.Update("")
	} else {
		m.viewport, cmd = m.viewport.Update(msg)
	}
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}


func (m model) View() string {
	
	if !m.ready {
		return "\n  Initializing..."
	}

	var leftStyle lipgloss.Style
	var rightStyle lipgloss.Style

	if m.workerTable.Focused() {
		leftStyle = FocusStyle
		rightStyle = BaseStyle
	} else {
		leftStyle = BaseStyle
		rightStyle = FocusStyle
	}

	var rightSide string
	if m.showDetails {
		rightSide = rightStyle.Render(m.viewport.View())
		} else {
	  rightSide = rightStyle.Render(m.jobsTable.View())
	}

	return lipgloss.JoinHorizontal(
	  lipgloss.Left,
	  leftStyle.Render(m.workerTable.View()),
		rightSide,
	 )
}
