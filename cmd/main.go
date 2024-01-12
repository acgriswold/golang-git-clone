package main

import (
	"fmt"
	"os"

	"github.com/acgriswold/golang-git-example/internal/commands"
)

func main() {
	if err := commands.Help.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
