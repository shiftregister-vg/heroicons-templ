package main

import "github.com/stevegood/heroicons-templ/internal/git"

const (
	heroiconsRepo = "https://github.com/heroicons/heroicons.git"
)

func main() {
	println("Importing heroicons...")
	r := git.Clone(heroiconsRepo)
	println(r.Head)
}
