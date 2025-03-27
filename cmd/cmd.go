package cmd

import "github.com/spf13/cobra"

var IsPkUuid bool

func GetCommandExecuter() *cobra.Command {
	GenCmd.Flags().BoolVarP(&IsPkUuid, "pk-uuid", "", false, "Indicates if model has a UUID as pk. If omitted Id will be an int")

	RootCmd.AddCommand(BootstrapCmd)
	RootCmd.AddCommand(VersionCmd)
	RootCmd.AddCommand(GenCmd)

	return RootCmd
}
