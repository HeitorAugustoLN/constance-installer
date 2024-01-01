package styles

import "github.com/charmbracelet/lipgloss"

var GreenText = lipgloss.NewStyle().Foreground(lipgloss.Color("#88f97e"))
var CyanText = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#80fde9"))
var YellowText = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#e5e577"))

var PurpleBoldText = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("#9580fd"))
