package main

import (
	"github.com/charmbracelet/log"
	"github.com/manicar2093/gomancer/cmd"
)

func main() {
	log.SetLevel(log.DebugLevel)

	cmd.GetCommandExecuter().Execute()
}
