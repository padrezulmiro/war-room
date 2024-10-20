package main

import (
	"strings"
	tea "github.com/charmbracelet/bubbletea"
)

type toolSelectionModel struct {
	currentSelection toolSelectionType
}

type toolSelectionType int
const (
	BuildingPlannerTool toolSelectionType = iota
	InvasionPlannerTool
)

var toolSelectionTypeStr = map[toolSelectionType]string{
	BuildingPlannerTool: "Building planner",
	InvasionPlannerTool: "Invasion planner",
}

var toolTypeToModelStateMap = map[toolSelectionType]modelState{
	BuildingPlannerTool: MapSelectionMenuState,
	InvasionPlannerTool: UnderConstructionState,
}

func (model *toolSelectionModel) Init() {
	model.currentSelection = BuildingPlannerTool
}

func (model *toolSelectionModel) Update(msg tea.Msg) (modelState, tea.Cmd) {
	retModelState := ToolSelectionMenuState
	var retCmd tea.Cmd = nil

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return retModelState, tea.Quit

		case "up", "k":
			if model.currentSelection > BuildingPlannerTool {
				model.currentSelection--
			}

		case "down", "j":
			if model.currentSelection < InvasionPlannerTool {
				model.currentSelection++
			}

		case "enter":
			retModelState = toolTypeToModelStateMap[model.currentSelection]
		}
	}

	return retModelState, retCmd
}

func (model *toolSelectionModel) View() string {
	var sb strings.Builder
	sb.WriteString(heading1Style.Render("Screeps War Room"))
	sb.WriteRune('\n')
	sb.WriteRune('\n')
	sb.WriteString(heading2Style.Render("Select a tool:"))
	sb.WriteRune('\n')

	for i := BuildingPlannerTool; i <= InvasionPlannerTool; i++ {
		if model.currentSelection == i {
			sb.WriteString(inputStyle.Render("\t" + toolSelectionTypeStr[i]))
		} else {
			sb.WriteString(greyedOutStyle.Render("\t" + toolSelectionTypeStr[i]))
		}
		sb.WriteRune('\n')
	}

	return sb.String()
}
