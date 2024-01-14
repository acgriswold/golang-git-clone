package commands

import (
	"bytes"
	"compress/zlib"
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
)

var catFile = &cobra.Command{
	Use:   "cat-file",
	Short: "Reads out the tracked git blob based on a given SHA",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		sha := args[0]

		path := fmt.Sprintf("%s/objects/%s", root, sha)

		contents, err := os.ReadFile(path)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Count not read file with give sha (%s)\n%v\n", sha, err)
			return err
		}

		buffer := bytes.NewBuffer(contents)
		reader, err := zlib.NewReader(buffer)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Could not set up reader for decompressing sha\n%v\n", err)
			return err
		}

		defer reader.Close()

		file, err := io.ReadAll(reader)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Could not read contents of decompressed file\n%v\n", err)
		}

		res := bytes.Split(file, []byte{0})
		fmt.Print(string(res[len(res)-1]))

		return nil
	},
}

func init() {
	catFile.Flags().StringP(
		"pretty-print",
		"p",
		"",
		"pretty-print <object> content",
	)
}
