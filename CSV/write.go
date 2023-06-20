package CSV

import (
	"Crawler_MPSP/Crawler"
	"os"
	"path/filepath"
)

const folderName = "Result"

func WriteCSV(lawsuits []Crawler.Lawsuit) error {
	err := WriteCovers(lawsuits)
	if err != nil {
		return err
	}

	err = WriteMovements(lawsuits)
	if err != nil {
		return err
	}

	err = WriteDocuments(lawsuits)
	if err != nil {
		return err
	}

	err = WriteBonds(lawsuits)
	if err != nil {
		return err
	}

	return nil
}

func createFile(p string) (*os.File, error) {
	if err := os.MkdirAll(filepath.Dir(p), 0770); err != nil {
		return nil, err
	}
	return os.Create(p)
}
