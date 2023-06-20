package Crawler

import (
	"errors"
	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
	"strings"
)

type Movement struct {
	Lawsuit string
	Date    string
	Title   string
	Text    string
}

const (
	xpathMovements = "//*[@id=\"body\"]/section/section/table[3]/tbody/tr"
	xpathDate      = "td[1]"
	xpathTitle     = "td[2]/span"
	xpathText      = "td[3]/span"
	Dirt           = "\n"
)

func GetLawsuitMovements(htmlPgSrc *html.Node, lawsuit string) ([]Movement, error) {
	movements := htmlquery.Find(htmlPgSrc, xpathMovements)

	if len(movements) > 0 {
		var allMovements []Movement
		for _, movement := range movements {
			var date string
			dt := htmlquery.Find(movement, xpathDate)
			if len(dt) > 0 {
				date = strings.TrimSpace(strings.Replace(htmlquery.InnerText(htmlquery.FindOne(movement, xpathDate)), Dirt, "", -1))
			}

			var title string
			tt := htmlquery.Find(movement, xpathTitle)
			if len(tt) > 0 {
				title = strings.TrimSpace(strings.Replace(htmlquery.InnerText(htmlquery.FindOne(movement, xpathTitle)), Dirt, "", -1))
			}

			var text string
			txt := htmlquery.Find(movement, xpathText)
			if len(txt) > 0 {
				text = strings.TrimSpace(strings.Replace(htmlquery.InnerText(htmlquery.FindOne(movement, xpathText)), Dirt, "", -1))
			}

			mv := Movement{
				Lawsuit: lawsuit,
				Date:    date,
				Title:   strings.Replace(strings.Replace(title, text, "", -1), Dirt, "", -1),
				Text:    text,
			}

			allMovements = append(allMovements, mv)

		}

		return allMovements, nil

	}

	return nil, errors.New("could not find movements")
}
