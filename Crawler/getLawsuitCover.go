package Crawler

import (
	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
	"strings"
)

type LawsuitCover struct {
	NumberMP            string
	NumberTJ            string
	ProcedureType       string
	Unit                string
	Situation           string
	Subject             string
	Poles               string
	InitialDate         string
	NumberInvestigation string
	NamePrecinct        string
	FirstCourt          string
	Position            string
	SecondCourt         string
}

const (
	titleNumberMP            = "Número MP: "
	titleNumberTJ            = "Número TJ: "
	titleProcedureType       = "Tipo de Procedimento: "
	titleUnit                = "Unidade: "
	titleSituation           = "Situação: "
	titleSubject             = "Assunto: "
	titlePoles               = "Partes: "
	titleInitialDate         = "Instauração: "
	titleNumberInvestigation = "N° do Inquérito na Delegacia: "
	titleNamePrecinct        = "Nome da Delegacia: "
	titleFirstCourt          = "Vara de Origem: "
	titlePosition            = "Cargo: "
	titleSecondCourt         = "Remetido para: "
)

func GetLawsuitCover(htmlPgSrc *html.Node) (LawsuitCover, error) {
	var numberMP string
	var numberTJ string
	var procedureType string
	var unit string
	var situation string
	var subject string
	var poles string
	var initialDate string
	var numberInvestigation string
	var namePrecinct string
	var firstCourt string
	var position string
	var secondCourt string

	totalNodes := htmlquery.Find(htmlPgSrc, "//*[@id=\"body\"]/section/section/div")
	for _, node := range totalNodes {
		title := htmlquery.InnerText(htmlquery.FindOne(node, "/label"))
		information := strings.ReplaceAll(strings.TrimSpace(htmlquery.InnerText(htmlquery.FindOne(node, "/span"))), "\n", "")

		if title == titleNumberMP {
			numberMP = information
		} else if title == titleNumberTJ {
			numberTJ = information
		} else if title == titleProcedureType {
			procedureType = information
		} else if title == titleUnit {
			unit = information
		} else if title == titleSituation {
			situation = information
		} else if title == titleSubject {
			subject = information
		} else if title == titlePoles {
			poles = information
		} else if title == titleInitialDate {
			initialDate = information
		} else if title == titleNumberInvestigation {
			numberInvestigation = information
		} else if title == titleNamePrecinct {
			namePrecinct = information
		} else if title == titleFirstCourt {
			firstCourt = information
		} else if title == titlePosition {
			position = information
		} else if title == titleSecondCourt {
			secondCourt = information
		}
	}

	return LawsuitCover{
		NumberMP:            numberMP,
		NumberTJ:            numberTJ,
		ProcedureType:       procedureType,
		Unit:                unit,
		Situation:           situation,
		Subject:             subject,
		Poles:               poles,
		InitialDate:         initialDate,
		NumberInvestigation: numberInvestigation,
		NamePrecinct:        namePrecinct,
		FirstCourt:          firstCourt,
		Position:            position,
		SecondCourt:         secondCourt,
	}, nil

}
