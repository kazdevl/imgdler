package open

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"text/template"

	"github.com/pkg/browser"
	"github.com/spf13/cobra"
)

func NewCmd(contentsDir string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "open [author_name]",
		Short: "open reader",
		Long:  `open is for launch reader to read images in author_name dir`,
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			err := proccess(args[0], contentsDir)
			if err != nil {
				log.Println("found error")
			}
		},
	}
	return cmd
}

func proccess(author, contentsDir string) error {
	contentDir := filepath.Join(contentsDir, author)
	des, err := os.ReadDir(contentDir)
	if err != nil {
		return err
	}
	images := make([]string, len(des))
	for i, de := range des {
		filename := de.Name()
		fileInfo := strings.Split(filename, ".")
		if len(fileInfo) <= 1 {
			continue
		}
		if fileInfo[1] != "jpg" {
			continue
		}
		images[i] = fmt.Sprintf("%s/%s", contentDir, de.Name())
	}
	sort.Strings(images)

	// create html
	t, err := template.New("reader").Parse(tpl)
	if err != nil {
		return err
	}
	f, err := os.Create(fmt.Sprintf("%s/index.html", contentDir))
	defer f.Close()
	if err != nil {
		return err
	}
	if err := t.Execute(f, images); err != nil {
		return err
	}

	// launch browser
	if err := browser.OpenFile(fmt.Sprintf("%s/index.html", contentDir)); err != nil {
		return err
	}
	return nil
}

const tpl = `
<!doctype html>
<html lang="ja">

<head>
    <!-- Required meta tags -->
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">

    <!-- Bootstrap CSS -->
    <link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.2.1/css/bootstrap.min.css"
        integrity="sha384-GJzZqFGwb1QTTN6wy59ffF1BuGJpLSa9DkKMp0DgiMDm4iYMj70gZWKYbI706tWS" crossorigin="anonymous">
    <title>imgdler</title>
</head>

<body class="bg-secondary">
    <header class="fixed-top mb-3 bg-dark" style="opacity: 0.5;" onMouseOut="this.style.opacity=0"
        onMouseOver="this.style.opacity=1">
        <nav class="navbar navbar-dark w-75 mr-auto ml-auto">
            <div class="container-fluid">
                <span class="navbar-brand">
                    imgdler
                </span>
            </div>
        </nav>
    </header>
    <main class="w-50 mr-auto ml-auto">
        <div class="bg-light w-100 text-center">
            {{range .}}
            <img src="{{.}}" class="img-fluid mb-1">
            {{end}}
        </div>
    </main>
    <!-- Optional JavaScript -->
    <!-- jQuery first, then Popper.js, then Bootstrap JS -->
    <script src="https://code.jquery.com/jquery-3.3.1.slim.min.js"
        integrity="sha384-q8i/X+965DzO0rT7abK41JStQIAqVgRVzpbzo5smXKp4YfRvH+8abtTE1Pi6jizo"
        crossorigin="anonymous"></script>
    <script src="https://cdnjs.cloudflare.com/ajax/libs/popper.js/1.14.6/umd/popper.min.js"
        integrity="sha384-wHAiFfRlMFy6i5SRaxvfOCifBUQy1xHdJ/yoi7FRNXMRBu5WHdZYu1hA6ZOblgut"
        crossorigin="anonymous"></script>
    <script src="https://stackpath.bootstrapcdn.com/bootstrap/4.2.1/js/bootstrap.min.js"
        integrity="sha384-B0UglyR+jN6CkvvICOB2joaf5I4l3gm9GU6Hc1og6Ls7i6U/mkkaduKaBhlAXv9k"
        crossorigin="anonymous"></script>
</body>

</html>
`
