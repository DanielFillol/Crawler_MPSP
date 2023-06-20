package CSV

import (
	"Crawler_MPSP/Crawler"
	"encoding/csv"
)

const fileNameB = "Bonds"

func WriteBonds(lawsuits []Crawler.Lawsuit) error {
	var rows [][]string

	rows = tableBondsRows(lawsuits)

	cf, err := createFile(folderName + "/" + fileNameB + ".csv")
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

func tableBondsRows(lawsuits []Crawler.Lawsuit) [][]string {
	var dcts [][]string

	dcts = append(dcts, []string{"Processo", "Titulo", "Link", "Tipo Documento"})

	for i := 0; i < len(lawsuits); i++ {
		for j := 0; j < len(lawsuits[i].Bonds); j++ {
			dcts = append(dcts, []string{lawsuits[i].Bonds[j].Lawsuit, lawsuits[i].Bonds[j].Title, lawsuits[i].Bonds[j].Link, lawsuits[i].Bonds[j].DocumentType})
		}
	}
	return dcts
}
