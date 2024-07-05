package main

import (
	tea "github.com/charmbracelet/bubbletea"
)

type toolSelectionModel struct {
	currentSelection int
	chosenTool string
}

var tools = [2]string{
	"Building planner",
	"Invasion planner",
}

func (model *toolSelectionModel) Update(msg tea.Msg) (modelState, tea.Cmd) {
	retModelState := ToolSelectionMenuState
	var retCmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			retCmd = tea.Quit

		case "up", "k":
			if model.currentSelection > 0 {
				model.currentSelection--
			}

		case "down", "j":
			if model.currentSelection < len(tools) - 1 {
				model.currentSelection++
			}

		case "enter":
			model.chosenTool = tools[model.currentSelection]
		}
	}

	return retModelState, retCmd
}

func (model *toolSelectionModel) View() string {
	vStr := "SCREEPS WAR ROOM\n"
	vStr += "----------------\n\n"

	for i, tool := range tools {
		if model.currentSelection == i {
			vStr += "> "
		} else {
			vStr += "- "
		}

		vStr += tool + "\n"
	}

	return vStr
}
