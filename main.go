package main

import (
	"github.com/Enthys/ebook_formatter/pkg"
)

type Example struct {
	ThisIsTheFirstName string
	ThisIsLast int
}

func main() {
	archive, err := pkg.NewArchive("./Project Hail Mary (Andy Weir).epub")
	if err != nil {
		panic(err)
	}

	files := archive.GetCSSFIles()
	firstFileContents, err := archive.GetFileContents(files[0])
	if err != nil {
		panic(err)
	}

	cssReader := pkg.NewCssReader(firstFileContents)

	cssReader.GetCssSelectors()
}
