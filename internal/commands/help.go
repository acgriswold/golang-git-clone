package commands

import "github.com/spf13/cobra"

var Help = &cobra.Command{
	Use:   "help",
	Short: "Version control system (git clone) tracking changes in a set of computer files.",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

func init() {
	/**
	Used within main.go to setup entry point.
	Initialize and add all valid cobra commands to be used
	within the cli.
	*/

	Help.AddCommand(Init)
}
