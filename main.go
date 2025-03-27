package main

import (
	"github.com/charmbracelet/log"
	"github.com/manicar2093/gomancer/cmd"
)

var IsPkUuid bool

func main() {
	log.SetLevel(log.DebugLevel)
	cmd.GenCmd.Flags().BoolVarP(&IsPkUuid, "pk-uuid", "", false, "Indicates if model has a UUID as pk. If omitted Id will be an int")

	cmd.RootCmd.AddCommand(cmd.BootstrapCmd)
	cmd.RootCmd.AddCommand(cmd.VersionCmd)
	cmd.RootCmd.AddCommand(cmd.GenCmd)

	cmd.RootCmd.Execute()
}
