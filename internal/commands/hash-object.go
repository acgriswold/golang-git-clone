package commands

import (
	"bytes"
	"compress/zlib"
	"crypto/sha1"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var hashObject = &cobra.Command{
	Use:   "hash-object",
	Short: "compute the object ID value for an object",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
        objectsDirectory := fmt.Sprintf("%s/objects", root)
        _, err := os.ReadDir(objectsDirectory)
        if err != nil {
            fmt.Println("Must initialize git repository before running commands")
            return err
        }

		flag := cmd.Flag("write")
		fileToHash := flag.Value.String()

		data, err := os.ReadFile(fileToHash)
		if err != nil {
			fmt.Printf("Error reading contents of %s", fileToHash)
		}
		filesize := string(rune(len(data)))
		header := []byte("blob " + filesize + "\x00")
		content := []byte(append(header, data...))

		hashWriter := sha1.New()
		hashWriter.Write(content)
		sha := hashWriter.Sum(nil)
		hash := fmt.Sprintf("%x", sha)

		/**
		  Generate path in ".golang-git/objects/".
		  - subDirectory = first two chars of hash
		  - fileName = rest of hash
		*/
		directoryToCreate := hash[:2]
		fileName := hash[2:]

		fileDirectory := fmt.Sprintf("%s/%s", objectsDirectory, directoryToCreate)
		if err := os.MkdirAll(fileDirectory, 0755); err != nil {
			fmt.Fprintf(os.Stderr, "Error creating directory: %s\n", err)
			return err
		}

		var buffer bytes.Buffer
		writer := zlib.NewWriter(&buffer)
		writer.Write(content)
		writer.Close()

		path := fmt.Sprintf("%s/%s", fileDirectory, fileName)
        err = os.WriteFile(path, buffer.Bytes(), 0755)
        if err != nil {
            fmt.Fprintf(os.Stderr, "Error writing file: %s\n", err)
            return err
        }

        fmt.Println(hash)

		return nil
	},
}

func init() {
	hashObject.Flags().StringP(
		"write",
		"w",
		"",
		"actually write object from standard input instead of from a file.",
	)
}
