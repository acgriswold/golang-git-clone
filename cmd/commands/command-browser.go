package commands

import (
	"github.com/spf13/cobra"
)

func NewCommandEventBrowser() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "init",
		Short:   "Scafold a new function",
		Example: "inngest init",
		Run:     runBrowser,
	}
	return cmd
}

func runBrowser(cmd *cobra.Command, args []string) {

}
