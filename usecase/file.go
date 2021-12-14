package usecase

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/kazdevl/imgdler/entity"
)

type FileUsecase struct {
	appDir string
}

func NewFileUsecase(dir string) *FileUsecase {
	return &FileUsecase{appDir: dir}
}

func (f *FileUsecase) AppDirName() string {
	return f.appDir
}

func (f *FileUsecase) CreateJpegs(auhtor string, pages entity.Pages) error {
	for i, link := range pages.Links {
		res, err := http.Get(link)
		if err != nil {
			log.Fatal(err)
		}
		defer res.Body.Close()

		f, err := os.Create(fmt.Sprintf("%s/%s/%s_%d.jpg", f.appDir, auhtor, pages.Datetime.Format("20060102030405"), i))
		if err != nil {
			return err
		}
		defer f.Close()

		io.Copy(f, res.Body)
	}
	return nil
}
