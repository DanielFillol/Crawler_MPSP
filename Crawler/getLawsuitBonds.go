package Crawler

import (
	"errors"
	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
	"strings"
)

type Bond struct {
	Lawsuit      string
	Title        string
	Link         string
	DocumentType string
}

const (
	xpathInfoBonds = "//*[@id=\"body\"]/section/section/table[1]/tbody/tr/td"
	xpathBonds     = "//*[@id=\"body\"]/section/section/table[1]/tbody/tr"
	xpathTitleBond = "td[1]/span/a"
	xpathBondLink  = "td[1]/span/a"
	xpathBondType  = "td[3]"
)

func GetLawsuitBonds(htmlPgSrc *html.Node, lawsuit string) ([]Bond, error) {
	alert := htmlquery.InnerText(htmlquery.FindOne(htmlPgSrc, xpathInfoBonds))
	if alert == "Não há vínculos!" {
		return []Bond{{
			Lawsuit:      lawsuit,
			Title:        alert,
			Link:         "-",
			DocumentType: "-",
		}}, nil
	}

	bonds := htmlquery.Find(htmlPgSrc, xpathBonds)

	if len(bonds) > 0 {
		var allBonds []Bond
		for _, bond := range bonds {
			var title string
			tt := htmlquery.Find(bond, xpathTitleBond)
			if len(tt) > 0 {
				title = strings.TrimSpace(strings.Replace(htmlquery.InnerText(htmlquery.FindOne(bond, xpathTitleBond)), Dirt, "", -1))
			}

			var link string
			lk := htmlquery.Find(bond, xpathBondLink)
			if len(lk) > 0 {
				lkPath := htmlquery.FindOne(bond, xpathBondLink)
				link = "https://sismpconsultapublica.mpsp.mp.br" + htmlquery.SelectAttr(lkPath, "href")
			}

			var documentType string
			dt := htmlquery.Find(bond, xpathBondType)
			if len(dt) > 0 {
				documentType = strings.TrimSpace(strings.Replace(htmlquery.InnerText(htmlquery.FindOne(bond, xpathBondType)), Dirt, "", -1))
			}

			dc := Bond{
				Lawsuit:      lawsuit,
				Title:        title,
				Link:         link,
				DocumentType: documentType,
			}

			allBonds = append(allBonds, dc)

		}

		return allBonds, nil

	}

	return nil, errors.New("could not find documents")
}
