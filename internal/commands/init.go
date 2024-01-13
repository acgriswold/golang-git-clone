package commands

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

const root = ".git-golang"

var Init = &cobra.Command{
	Use:   "init .golang-git repository",
	Short: "Initializes needed files and folder structure for versioning",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		for _, dir := range []string{root, fmt.Sprintf("%s/objects", root), fmt.Sprintf("%s/refs", root)} {
			if err := os.MkdirAll(dir, 0755); err != nil {
				fmt.Fprintf(os.Stderr, "Error creating directory: %s\n", err)
			}
		}

		headFileContents := []byte("ref: refs/heads/master\n")

		if err := os.WriteFile(fmt.Sprintf("%s/HEAD", root), headFileContents, 0644); err != nil {
			fmt.Fprintf(os.Stderr, "Error writing file: %s\n", err)
		}

		fmt.Println("Initialized git directory")

		return nil
	},
}
