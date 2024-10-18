package main

import (
	"github.com/charmbracelet/lipgloss"
)

var colorMap = map[string]string{
	"White": "#EDEDED",
	"LightBlue": "#AFC0D4",
	"MediumBlue": "#506E91",
	"DarkBlue": "#2D3F62",
	"Black": "#091A34",
}

var (
	normalColor = lipgloss.Color(colorMap["White"])
	greyedOutColor = lipgloss.Color(colorMap["MediumBlue"])
	backgroundColor = lipgloss.Color(colorMap["Black"])
	heading1Color = lipgloss.Color(colorMap["White"])
	heading2Color = lipgloss.Color(colorMap["LightBlue"])

	inputStyle = lipgloss.NewStyle().
		Foreground(normalColor).
		TabWidth(2)

	heading1Style = lipgloss.NewStyle().
		Foreground(heading1Color).
		Padding(1, 2).
		Border(lipgloss.ThickBorder(), true).
		BorderForeground(heading1Color)

	heading2Style = lipgloss.NewStyle().
		Foreground(heading2Color)

	greyedOutStyle = lipgloss.NewStyle().
		Foreground(greyedOutColor).
		TabWidth(2)
)
