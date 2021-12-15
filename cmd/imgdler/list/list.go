package list

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

func NewCmd(contentsDir string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "list author_names",
		Long:  `list is for listing author_names`,
		Run: func(cmd *cobra.Command, args []string) {
			err := proccess(contentsDir)
			if err != nil {
				log.Println("found error")
			}
		},
	}
	return cmd
}

func proccess(contentsDir string) error {
	des, err := os.ReadDir(contentsDir)
	if err != nil {
		return err
	}
	fmt.Println("The list of author names that you can read")
	for i, de := range des {
		name := de.Name()
		if name[0] == '.' {
			continue
		}
		fmt.Printf("[%d]: %s\n", i, name)
	}
	fmt.Println("You cna read with `imgdler open [author name]`")
	return nil
}
