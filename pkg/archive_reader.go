package pkg

import (
	"fmt"
	"strings"

	"github.com/gen2brain/go-unarr"
)

type Archive struct {
	archive *unarr.Archive
}

func NewArchive(path string) (*Archive, error) {
	arr, err := unarr.NewArchive(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open archive. Error: %s", err.Error())
	}

	return &Archive{
		archive: arr,
	}, nil
}

func (a *Archive) GetCSSFIles() []string {
	files, err := a.archive.List()

	if err != nil {
		panic(err)
	}

	result := make([]string, 0)
	for _, file := range files {
		if strings.HasSuffix(file, ".css") {
			result = append(result, file)
		}
	}

	return result
}

func (a *Archive) GetFileContents(fileName string) ([]byte, error) {
	err := a.archive.EntryFor(fileName)
	if err != nil {
		return []byte{}, err
	}

	return a.archive.ReadAll()
}

