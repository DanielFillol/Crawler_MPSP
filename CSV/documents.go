package CSV

import (
	"Crawler_MPSP/Crawler"
	"encoding/csv"
)

const fileNameD = "Documents"

func WriteDocuments(lawsuits []Crawler.Lawsuit) error {
	var rows [][]string

	rows = tableDocumentsRows(lawsuits)

	cf, err := createFile(folderName + "/" + fileNameD + ".csv")
	if err != nil {
		return err
	}

	w := csv.NewWriter(cf)

	err = w.WriteAll(rows)
	if err != nil {
		return err
	}

	return nil
}

func tableDocumentsRows(lawsuits []Crawler.Lawsuit) [][]string {
	var dcts [][]string

	dcts = append(dcts, []string{"Processo", "Titulo", "Link", "Tipo Documento"})

	for i := 0; i < len(lawsuits); i++ {
		for j := 0; j < len(lawsuits[i].Documents); j++ {
			dcts = append(dcts, []string{lawsuits[i].Documents[j].Lawsuit, lawsuits[i].Documents[j].Title, lawsuits[i].Documents[j].Link, lawsuits[i].Documents[j].DocumentType})
		}
	}
	return dcts
}
