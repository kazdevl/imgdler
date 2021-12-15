package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/kazdevl/imgdler/cmd/imgdler/list"
	"github.com/kazdevl/imgdler/cmd/imgdler/open"
	"github.com/kazdevl/imgdler/cmd/imgdler/start"
	"github.com/spf13/cobra"
)

func main() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	contentsDir := filepath.Join(homeDir, "imgdler", "contents")
	if err := os.MkdirAll(contentsDir, 0755); err != nil {
		log.Fatal(err)
	}

	var rootCmd = &cobra.Command{Use: "imgdler"}
	rootCmd.AddCommand(
		start.NewCmd(contentsDir),
		open.NewCmd(contentsDir),
		list.NewCmd(contentsDir),
	)
	cobra.CheckErr(rootCmd.Execute())
}
