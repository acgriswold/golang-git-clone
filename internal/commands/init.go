package commands

import "github.com/spf13/cobra"

var Init = &cobra.Command{
	Use:   "init .golang-git repository",
	Short: "Initializes needed files and folder structure for versioning",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		// TODO: Init .golang-git folder
		return nil
	},
}
