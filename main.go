package main

import (
	"fmt"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type timerType int

const (
	workSession timerType = iota
	shortBreak
	longBreak
)

// Model for managing the timer state
type model struct {
	timerDuration int
	remaining     int
	running       bool
	timerCycle    int // Tracks the current cycle in the loop
	timerType     timerType
	width         int // Terminal width
	height        int // Terminal height
}

// tickMsg is the message sent after every tick
type tickMsg struct{}

// Init initializes the Bubble Tea app.
func (m model) Init() tea.Cmd {
	m.running = true // Auto-start the first timer
	return tick()
}

// Update processes messages and updates the model
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		// Update terminal dimensions
		m.width = msg.Width
		m.height = msg.Height

	case tea.KeyMsg:
		// Handle user input
		switch msg.String() {
		case "s": // Start the timer
			if !m.running {
				m.running = true
				return m, tick()
			}
		case "p": // Pause the timer
			m.running = false
		case "q": // Quit the program
			return m, tea.Quit
		}

	case tickMsg:
		// Timer countdown logic
		if m.running && m.remaining > 0 {
			m.remaining--
			if m.remaining == 0 {
				m.running = false
				switch m.timerType {
				case workSession:
					if m.timerCycle < 3 { // First three work sessions followed by short breaks
						m.timerType = shortBreak
						m.timerDuration = 5 // 5 seconds for testing
						m.timerCycle++
					} else { // Fourth work session followed by a long break
						m.timerCycle = 0 // Reset the cycle count
						m.timerType = longBreak
						m.timerDuration = 30 // 30 seconds for testing
					}
				case shortBreak:
					m.timerType = workSession
					m.timerDuration = 25 // 25 seconds for testing
				case longBreak:
					m.timerType = workSession
					m.timerDuration = 25 // 25 seconds for testing
				}
				m.remaining = m.timerDuration
				m.running = true // Auto-start the next timer
				return m, tick()
			}
			return m, tick()
		}
	}

	return m, nil
}

// tick returns a command to send a `tickMsg` after one second
func tick() tea.Cmd {
	return tea.Tick(time.Second, func(time.Time) tea.Msg {
		return tickMsg{}
	})
}

// Lipgloss styles for colored text (removed bold)
var (
	redStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("#e78284"))                                           // Red
	greenStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("#a6d189"))                                           // Green
	blueStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("#8caaee"))                                           // Blue
	statusBar   = lipgloss.NewStyle().Foreground(lipgloss.Color("7")).Background(lipgloss.Color("236")).Padding(0, 1) // Status bar style
	grayStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("#838ba7"))
	yellowStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#e5c890")) // Green
)

// View renders the timer UI.
func (m model) View() string {
	var status, coloredTimer string

	if m.running {
		status = "Running"
	} else {
		status = "Stopped"
	}

	// Format remaining time as MM:SS
	minutes := m.remaining / 60
	seconds := m.remaining % 60
	timerDisplay := fmt.Sprintf("%02d:%02d", minutes, seconds)

	// Color the timer based on the session type
	switch m.timerType {
	case workSession:
		coloredTimer = redStyle.Render(timerDisplay)
	case shortBreak:
		coloredTimer = greenStyle.Render(timerDisplay)
	case longBreak:
		coloredTimer = blueStyle.Render(timerDisplay)
	}

	// Status Bar (Cycle info)
	cycleInfo := fmt.Sprintf("Cycle: %d/4", m.timerCycle+1)
	title := fmt.Sprintf("Timer: %s", status)
	statusBarContent := statusBar.Render(cycleInfo, title)

	// Center the content within the terminal (excluding the status bar)
	body := fmt.Sprintf("Time Left: %s", coloredTimer)
	controls := grayStyle.Render("Press 's' to start, 'p' to pause, 'q' to quit.")
	content := fmt.Sprintf("%s\n\n%s", body, controls)

	// Calculate vertical centering for the content
	padding := (m.height - 7) / 2
	emptyLine := "\n"
	for i := 0; i < padding; i++ {
		emptyLine += "\n"
	}

	// Combine everything: status bar at the top and centered main content
	return statusBarContent + "\n" + emptyLine + lipgloss.NewStyle().
		Width(m.width).
		Align(lipgloss.Center).
		Render(content)
}

func main() {
	// Define the initial state
	initialModel := model{
		timerDuration: 25, // 25 seconds for testing
		remaining:     25, // 25 seconds for testing
		timerCycle:    0,
		timerType:     workSession,
	}

	// Create a new Bubble Tea program and start it
	p := tea.NewProgram(initialModel, tea.WithAltScreen())

	// Start the program and handle events
	if err := p.Start(); err != nil {
		fmt.Println("Error starting Bubble Tea program:", err)
		os.Exit(1)
	}
}
