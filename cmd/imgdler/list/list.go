package list

import (
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
	for _, de := range des {
		name := de.Name()
		if name[0] == '.' {
			continue
		}
		log.Println(name)
	}
	return nil
}
