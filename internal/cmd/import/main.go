package main

import (
	"io"
	"os"
	"path"
	"strings"
	"text/template"

	"github.com/go-git/go-billy/v5"
	"github.com/go-git/go-billy/v5/memfs"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/storage/memory"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

const (
	heroiconsRepo = "https://github.com/tailwindlabs/heroicons.git"
	templTemplate = `package {{.Package}}

templ {{.Name}}() {
	@{{.Name}}WithAttrs(templ.Attributes{})
}

templ {{.Name}}WithAttrs(attrs templ.Attributes) {
	{{.Content}}
}
`
)

type Icon struct {
	Name    string
	Content string
	Package string
}

func main() {
	println("Importing heroicons...")

	fs := memfs.New()

	println("Cloning heroicons git repo")
	_, err := git.Clone(memory.NewStorage(), fs, &git.CloneOptions{
		URL: heroiconsRepo,
	})
	if err != nil {
		panic(err)
	}

	println("Discovering icons...")
	fileInfos, err := fs.ReadDir("src")
	if err != nil {
		panic(err)
	}
	walkDir(fs, "src", fileInfos)
}

func walkDir(fs billy.Filesystem, parent string, fileInfos []os.FileInfo) {

	for _, fileInfo := range fileInfos {
		self := path.Join(parent, fileInfo.Name())

		if fileInfo.IsDir() {
			fi, err := fs.ReadDir(self)
			if err != nil {
				panic(err)
			}
			walkDir(fs, self, fi)
		}

		if strings.HasSuffix(self, ".svg") {
			// process the file...
			processFile(fs, self)
		}
	}
}

func processFile(fs billy.Filesystem, p string) {
	println("Processing " + p)

	file, err := fs.Open(p)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	icon := Icon{
		Name:    iconName(p),
		Content: strings.Replace(string(bytes), "<svg ", "<svg {attrs...} ", 1),
	}

	t, err := template.New(icon.Name).Parse(templTemplate)
	if err != nil {
		panic(err)
	}

	dir, _ := path.Split(p)
	dir = strings.ReplaceAll(dir, "src/", "")
	_, pkg := path.Split(trimSuffix(dir, "/"))
	icon.Package = pkg

	err = os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		panic(err)
	}

	f, err := os.Create(path.Join(dir, icon.Name+".templ"))
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if err = t.Execute(f, icon); err != nil {
		panic(err)
	}
}

func iconName(p string) string {
	var out string
	_, name := path.Split(p)
	nameParts := strings.Split(strings.ReplaceAll(name, ".svg", ""), "-")
	caser := cases.Title(language.English, cases.Compact)
	for _, part := range nameParts {
		out = out + caser.String(part)
	}
	return out
}

func trimSuffix(s, suffix string) string {
	if strings.HasSuffix(s, suffix) {
		s = s[:len(s)-len(suffix)]
	}
	return s
}
