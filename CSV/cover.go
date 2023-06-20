package CSV

import (
	"Crawler_MPSP/Crawler"
	"encoding/csv"
)

const fileNameC = "Covers"

func WriteCovers(lawsuits []Crawler.Lawsuit) error {
	var rows [][]string

	rows = append(rows, generateCoverHeaders())

	for _, lawsuit := range lawsuits {
		rows = append(rows, tableCoverRows(lawsuit))
	}

	cf, err := createFile(folderName + "/" + fileNameC + ".csv")
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
func generateCoverHeaders() []string {
	return []string{
		"Número MP",
		"Número TJ",
		"Tipo de Procedimento",
		"Unidade",
		"Situação",
		"Assunto",
		"Partes",
		"Instauração",
		"N° do Inquérito na Delegacia",
		"Nome da Delegacia",
		"Vara de Origem",
		"Cargo",
		"Remetido para",
	}
}

func tableCoverRows(results Crawler.Lawsuit) []string {
	return []string{
		results.Cover.NumberMP,
		results.Cover.NumberTJ,
		results.Cover.ProcedureType,
		results.Cover.Unit,
		results.Cover.Situation,
		results.Cover.Subject,
		results.Cover.Poles,
		results.Cover.InitialDate,
		results.Cover.NumberInvestigation,
		results.Cover.NamePrecinct,
		results.Cover.FirstCourt,
		results.Cover.Position,
		results.Cover.SecondCourt,
	}
}
