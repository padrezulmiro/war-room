package main

import (
	tea "github.com/charmbracelet/bubbletea"
)

type underConstructionModel struct {
	content string
}

func (model *underConstructionModel) Init(content string) {
	model.content = content
}

func (model *underConstructionModel) Update(msg tea.Msg) (modelState, tea.Cmd) {
	retModelState := UnderConstructionState
	var retCmd tea.Cmd = nil

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return retModelState, tea.Quit
		}
	}

	return retModelState, retCmd
}

func (model *underConstructionModel) View() string {
	if model.content != "" {
		return model.content
	}

	return `404 - Apologies, this page is still under construction!
Please press q or Ctrl-c to exit the app.\n\n`
}
