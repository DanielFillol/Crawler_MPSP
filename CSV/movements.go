package CSV

import (
	"Crawler_MPSP/Crawler"
	"encoding/csv"
)

const fileNameM = "Movements"

func WriteMovements(lawsuits []Crawler.Lawsuit) error {
	var rows [][]string

	rows = tableMovementsRows(lawsuits)

	cf, err := createFile(folderName + "/" + fileNameM + ".csv")
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

func tableMovementsRows(lawsuits []Crawler.Lawsuit) [][]string {
	var mvts [][]string

	mvts = append(mvts, []string{"Processo", "Data", "Movimentação", "Detalhe"})

	for i := 0; i < len(lawsuits); i++ {
		for j := 0; j < len(lawsuits[i].Movements); j++ {
			mvts = append(mvts, []string{lawsuits[i].Movements[j].Lawsuit, lawsuits[i].Movements[j].Date, lawsuits[i].Movements[j].Title, lawsuits[i].Movements[j].Text})
		}
	}
	return mvts
}
