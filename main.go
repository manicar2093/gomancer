package main

import "github.com/manicar2093/gomancer/cmd"

func main() {
	cmd.RootCmd.AddCommand(cmd.BootstrapCmd)
	cmd.RootCmd.Execute()
}
