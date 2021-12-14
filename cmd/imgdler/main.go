package main

import (
	"github.com/kazdevl/imgdler/cmd/imgdler/start"
	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{Use: "imgdler"}
	rootCmd.AddCommand(
		start.NewCmd(),
	)
	cobra.CheckErr(rootCmd.Execute())
}
