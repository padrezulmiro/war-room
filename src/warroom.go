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
	UnderConstructionState
)

type rootModel struct {
	state modelState
	toolSelection toolSelectionModel
	mapSelection mapSelectionModel
	underConstruction underConstructionModel
}

func newRootModel() *rootModel {
	model := rootModel {
		ToolSelectionMenuState,
		toolSelectionModel{},
		mapSelectionModel{},
		underConstructionModel{},
	}

	model.init()
	return &model
}

func (rootModel *rootModel) init() {
	rootModel.toolSelection.Init()
	rootModel.mapSelection.Init()
}

func (model rootModel) Init() tea.Cmd {
	return nil
}

func (model rootModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var newModelState modelState
	var retCmd tea.Cmd

	switch model.state {
	case ToolSelectionMenuState:
		newModelState, retCmd = model.toolSelection.Update(msg)
	case MapSelectionMenuState:
		newModelState, retCmd = model.mapSelection.Update(msg)
	case UnderConstructionState:
		newModelState, retCmd = model.underConstruction.Update(msg)
	}

	model.state = newModelState
	return model, retCmd
}

func (model rootModel) View() string {
	switch model.state {
	case ToolSelectionMenuState:
		return model.toolSelection.View()
	case MapSelectionMenuState:
		return model.mapSelection.View()
	case UnderConstructionState:
		return model.underConstruction.View()
	}
	return "Error: This page shouldn't have been reached!"
}

func main() {
	if len(os.Args) == 2 && os.Args[1] == "-d" {
		fmt.Println("Will log to ../logs/warroom.log")
		file, err := tea.LogToFile("../logs/warroom.log", "Debug")
		if err != nil {
			fmt.Println("fatal: ", err)
			os.Exit(1)
		}
		defer file.Close()
	}

	teaProgram := tea.NewProgram(newRootModel(), tea.WithAltScreen())

	if _, err := teaProgram.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
