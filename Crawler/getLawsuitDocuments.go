package Crawler

import (
	"errors"
	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
	"strings"
)

type Document struct {
	Lawsuit      string
	Title        string
	Link         string
	DocumentType string
}

const (
	xpathInfoDocuments = "//*[@id=\"body\"]/section/section/table[2]/tbody/tr/td"
	xpathDocuments     = "//*[@id=\"body\"]/section/section/table[2]/tbody/tr"
	xpathTitleDocument = "td[1]/a/text()"
	xpathLink          = "td[1]/a"
	xpathDocumentType  = "td[2]/text()"
)

func GetLawsuitDocuments(htmlPgSrc *html.Node, lawsuit string) ([]Document, error) {
	alert := htmlquery.InnerText(htmlquery.FindOne(htmlPgSrc, xpathInfoDocuments))
	if alert == "Não há anexos!" {
		return []Document{{
			Lawsuit:      lawsuit,
			Title:        alert,
			Link:         "-",
			DocumentType: "-",
		}}, nil
	}

	documents := htmlquery.Find(htmlPgSrc, xpathDocuments)

	if len(documents) > 0 {
		var allDocuments []Document
		for _, document := range documents {
			var title string
			tt := htmlquery.Find(document, xpathTitleDocument)
			if len(tt) > 0 {
				title = strings.TrimSpace(strings.Replace(htmlquery.InnerText(htmlquery.FindOne(document, xpathTitleDocument)), Dirt, "", -1))
			}

			var link string
			lk := htmlquery.Find(document, xpathLink)
			if len(lk) > 0 {
				lkPath := htmlquery.FindOne(document, xpathLink)
				link = "https://sismpconsultapublica.mpsp.mp.br" + htmlquery.SelectAttr(lkPath, "href")
			}

			var documentType string
			dt := htmlquery.Find(document, xpathDocumentType)
			if len(dt) > 0 {
				documentType = strings.TrimSpace(strings.Replace(htmlquery.InnerText(htmlquery.FindOne(document, xpathDocumentType)), Dirt, "", -1))
			}

			dc := Document{
				Lawsuit:      lawsuit,
				Title:        title,
				Link:         link,
				DocumentType: documentType,
			}

			allDocuments = append(allDocuments, dc)

		}

		return allDocuments, nil

	}

	return nil, errors.New("could not find documents")
}
