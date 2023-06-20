package Crawler

import (
	"github.com/tebeka/selenium"
	"log"
)

const InitWebSite = "https://sismpconsultapublica.mpsp.mp.br/ConsultarProcedimentos/ObterProcedimentos"

type Lawsuit struct {
	Cover     LawsuitCover
	Bonds     []Bond
	Documents []Document
	Movements []Movement
}

func Craw(driver selenium.WebDriver, lawsuit string, poleName string, lawsuitDocument string) (Lawsuit, error) {
	searchLink := InitWebSite
	htmlPgSrc, err := SearchLawsuit(driver, searchLink, lawsuit, poleName, lawsuitDocument)
	if err != nil {
		log.Println(err)
		return Lawsuit{}, nil
	}

	cover, err := GetLawsuitCover(htmlPgSrc)
	if err != nil {
		log.Println(err)
		return Lawsuit{}, err
	}

	movements, err := GetLawsuitMovements(htmlPgSrc, lawsuit)
	if err != nil {
		log.Println(err)
		return Lawsuit{}, err
	}

	documents, err := GetLawsuitDocuments(htmlPgSrc, lawsuit)
	if err != nil {
		log.Println(err)
		return Lawsuit{}, err
	}

	bonds, err := GetLawsuitBonds(htmlPgSrc, lawsuit)
	if err != nil {
		log.Println(err)
		return Lawsuit{}, err
	}

	return Lawsuit{
		Cover:     cover,
		Bonds:     bonds,
		Documents: documents,
		Movements: movements,
	}, nil
}
