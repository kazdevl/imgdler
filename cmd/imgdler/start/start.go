package start

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/kazdevl/imgdler/usecase"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

type contentOfFlags struct {
	AuthorName string
	Keyword    string
	Token      string
	Max        int
}

type creator struct {
	fu *usecase.FileUsecase
	tu *usecase.TwitterUsecase
	c  contentOfFlags
}

func (cr *creator) fetchAndCreate() error {
	pagesList, err := cr.tu.FetchContent(cr.c.AuthorName, cr.c.Keyword, cr.c.Max)
	if err != nil {
		return err
	}
	if err := os.MkdirAll(fmt.Sprintf("%s/%s", cr.fu.ContentsDirName(), cr.c.AuthorName), 0755); err != nil {
		return err
	}
	for _, pages := range pagesList {
		if err := cr.fu.CreateJpegs(cr.c.AuthorName, pages); err != nil {
			return err
		}
	}
	return nil
}

func NewCmd(contentDir string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "start",
		Short: "download tweets images with auhtor, keyword, token, max",
		Long:  `download tweets images with auhtor, keyword, token, max`,
		Run: func(cmd *cobra.Command, args []string) {
			var c contentOfFlags
			c.AuthorName, c.Keyword, c.Token, c.Max = getFlagValues(cmd.Flags())
			err := proccess(c, contentDir)
			if err != nil {
				log.Println("found error")
			}
		},
	}
	cmd.Flags().StringP("author", "a", "", "set author_name")
	cmd.Flags().StringP("keyword", "k", "", "set keyword")
	cmd.Flags().StringP("token", "t", "", "set token")
	cmd.Flags().IntP("max", "m", 10, "set max tweets")
	cmd.MarkFlagRequired("author")
	cmd.MarkFlagRequired("keyword")
	cmd.MarkFlagRequired("token")
	return cmd
}

func getFlagValues(fSet *pflag.FlagSet) (a, k, t string, m int) {
	a, _ = fSet.GetString("author")
	k, _ = fSet.GetString("keyword")
	t, _ = fSet.GetString("token")
	m, _ = fSet.GetInt("max")
	return
}

func proccess(c contentOfFlags, contentsDir string) error {
	tu := usecase.NewTwitterUsecase(c.Token)
	fu := usecase.NewFileUsecase(contentsDir)
	cr := &creator{fu: fu, tu: tu, c: c}

	if err := cr.fetchAndCreate(); err != nil {
		return err
	}

	s := gocron.NewScheduler(time.Local)
	s.Every(1).Day().At("21:00").Do(cr.fetchAndCreate)
	s.StartAsync()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("stop downloader")
	return nil
}
