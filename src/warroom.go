package main

import (
	"fmt"
	"os"
	tea "github.com/charmbracelet/bubbletea"
)

type modelState int
const (
	ToolSelectionMenuState modelState = iota
	MapSelectionMenuState
	BuildingPlannerState
)

type rootModel struct {
	state modelState
	toolSelection toolSelectionModel
	mapSelection mapSelectionModel
	buildingPlanner buildingPlannerModel
}

type mapSelectionModel struct {
	selectedMap string
	selectedShard int
}

type buildingPlannerModel struct {}

func (model rootModel) Init() tea.Cmd {
	return nil
}

func (model rootModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch model.state {
	case ToolSelectionMenuState:
		return model, model.toolSelection.Update(msg)
	}
	return model, nil
}

func (model rootModel) View() string {
	switch model.state {
	case ToolSelectionMenuState:
		return model.toolSelection.View()
	}
	return "Error: This page shouldn't have been reached!"
}

func main() {
	file, err := tea.LogToFile("../logs/warroom.log", "Debug")
	if err != nil {
		fmt.Println("fatal: ", err)
		os.Exit(1)
	}
	defer file.Close()

	model := rootModel{
		ToolSelectionMenuState,
		toolSelectionModel{},
		mapSelectionModel{},
		buildingPlannerModel{},
	}

	teaProgram := tea.NewProgram(model)

	if _, err := teaProgram.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
