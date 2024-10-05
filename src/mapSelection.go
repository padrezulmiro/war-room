package main

import (
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type mapSelectionModel struct {
	mapInput textinput.Model
	shardInput textinput.Model
	selectedMap string
	selectedShard string
	isFocused bool
	highlightedField mapSelectionField
}

type mapSelectionField int

const (
	Map mapSelectionField = iota
	Shard
	Confirm
)

var mapSelectionFieldStr = map[mapSelectionField]string{
	Map: "Map",
	Shard: "Shard",
	Confirm: "Confirm",
}

func (model *mapSelectionModel) Init() {
	model.mapInput = textinput.New()
	model.mapInput.Placeholder = "N1W1"
	model.mapInput.Prompt = ""

	model.shardInput = textinput.New()
	model.shardInput.Placeholder = "1"
	model.shardInput.Prompt = ""
}

func (model *mapSelectionModel) Update(msg tea.Msg) (modelState, tea.Cmd) {
	retModelState := MapSelectionMenuState
	var retCmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return retModelState, tea.Quit

		case "q":
			if !model.isFocused {
				return retModelState, tea.Quit
			}

		case "up", "k":
			if model.highlightedField > Map && !model.isFocused {
				model.highlightedField--
			}

		case "down", "j":
			if model.highlightedField < Confirm && !model.isFocused {
				model.highlightedField++
			}

		case "enter":
			switch model.highlightedField {
			case Map:
				model.mapInput.Focus()
			case Shard:
				model.shardInput.Focus()
			}

			if model.isFocused {
				model.isFocused = false
			} else if model.highlightedField == Confirm {
				model.selectedMap = model.mapInput.Value()
				model.selectedShard = model.shardInput.Value()
				// TODO(azul) Change model state (as in the page)
			} else {
				model.isFocused = true
			}
		}
	}

	var mapInputCmd, shardInputCmd tea.Cmd
	model.mapInput, mapInputCmd = model.mapInput.Update(msg)
	model.shardInput, shardInputCmd = model.shardInput.Update(msg)

	return retModelState, tea.Batch(retCmd, mapInputCmd, shardInputCmd)
}

func (model *mapSelectionModel) View() string {
	var sb strings.Builder
	sb.WriteString("Choose a map:\n\n")

	sb.WriteString("  MAP\n")
	if model.highlightedField == Map && model.isFocused {
		sb.WriteString("> ")
	} else if model.highlightedField == Map && !model.isFocused {
		sb.WriteString("- ")
	} else {
		sb.WriteString("  ")
	}
	sb.WriteString(model.mapInput.View() + "\n")

	sb.WriteString("  SHARD\n")
	if model.highlightedField == Shard && model.isFocused {
		sb.WriteString("> ")
	} else if model.highlightedField == Shard && !model.isFocused {
		sb.WriteString("- ")
	} else {
		sb.WriteString("  ")
	}
	sb.WriteString(model.shardInput.View() + "\n")

	sb.WriteRune('\n')
	if model.highlightedField == Confirm {
		sb.WriteString("- ")
	} else {
		sb.WriteString("  ")
	}
	sb.WriteString("CONFIRM\n")

	return sb.String()
}
