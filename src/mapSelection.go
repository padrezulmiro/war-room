package main

import tea "github.com/charmbracelet/bubbletea"

type mapSelectionModel struct {
	selectedMap string
	selectedShard int
}

var shards = [3]string{
	"shard1",
	"shard2",
	"shard3",
}

func (model *mapSelectionModel) Update(msg tea.Msg) (modelState, tea.Cmd) {
	return 0 , nil
}

func (model *mapSelectionModel) View() string {
	return "This page is under construction!"
}

func validateMapInput(mapInput string) bool {
	return false
}
