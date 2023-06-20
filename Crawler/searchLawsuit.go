package Crawler

import (
	"errors"
	"github.com/antchfx/htmlquery"
	"github.com/tebeka/selenium"
	"golang.org/x/net/html"
	"log"
	"strings"
	"time"
)

const (
	xpathName      = "//*[@id=\"NomeParte\"]"
	xpathDocument  = "//*[@id=\"DocParte\"]"
	xpathSearchBtt = "//*[@id=\"btnConsultar\"]"
	xpathNumber1   = "//*[@id=\"NumeroMPTipo\"]"
	xpathNumber2   = "//*[@id=\"NumeroMPUA\"]"
	xpathNumber3   = "//*[@id=\"NumeroMPSequencial\"]"
	xpathNumber4   = "//*[@id=\"NumeroMPAno\"]"
	xpathReturn    = "//*[@id=\"resultado\"]/p"
)

const textReturn = "Não existem resultados para a pesquisa informada."

type lawsuitNumber struct {
	number1 string
	number2 string
	number3 string
	number4 string
}

func SearchLawsuit(driver selenium.WebDriver, searchLink string, lawsuit string, poleName string, poleDocument string) (*html.Node, error) {
	err := driver.Get(searchLink)
	if err != nil {
		return nil, errors.New("url unavailable")
	}

	n1, err := driver.FindElement(selenium.ByXPATH, xpathNumber1)
	if err != nil {
		return nil, errors.New("could not find xpathNumber1")
	}

	n2, err := driver.FindElement(selenium.ByXPATH, xpathNumber2)
	if err != nil {
		return nil, errors.New("could not find xpathNumber2")
	}

	n3, err := driver.FindElement(selenium.ByXPATH, xpathNumber3)
	if err != nil {
		return nil, errors.New("could not find xpathNumber3")
	}

	n4, err := driver.FindElement(selenium.ByXPATH, xpathNumber4)
	if err != nil {
		return nil, errors.New("could not find xpathNumber4")
	}

	number, err := formatNumber(lawsuit)
	if err != nil {
		log.Println(err, number)
		return nil, err
	}

	err = n1.SendKeys(number.number1)
	err = n2.SendKeys(number.number2)
	err = n3.SendKeys(number.number3)
	err = n4.SendKeys(number.number4)
	if err != nil {
		return nil, errors.New("could not send lawsuit number as search parameter")
	}

	name, err := driver.FindElement(selenium.ByXPATH, xpathName)
	if err != nil {
		return nil, errors.New("could not find xpathName")
	}

	err = name.SendKeys(poleName)
	if err != nil {
		return nil, errors.New("could not input name as search parameter")
	}

	document, err := driver.FindElement(selenium.ByXPATH, xpathDocument)
	if err != nil {
		return nil, errors.New("could not find xpathDocument")
	}

	err = document.SendKeys(poleDocument)
	if err != nil {
		return nil, errors.New("could not input document as search parameter")
	}

	btt, err := driver.FindElement(selenium.ByXPATH, xpathSearchBtt)
	if err != nil {
		return nil, errors.New("could not find xpathSearchBtt")
	}

	err = btt.Click()
	if err != nil {
		return nil, errors.New("could not click on search button")
	}

	//there is a delay of 0.5s to load the page
	time.Sleep(1500 * time.Millisecond)

	pageSource, err := driver.PageSource()
	if err != nil {
		return nil, errors.New("could not get page source")
	}

	htmlPgSrc, err := htmlquery.Parse(strings.NewReader(pageSource))
	if err != nil {
		return nil, errors.New("could not convert string to node html")
	}

	exist := existLawsuit(htmlPgSrc)
	if exist != true {
		log.Println("could not find lawsuit")
		return nil, errors.New("could not find lawsuit")
	}

	return htmlPgSrc, nil
}

func formatNumber(lawsuit string) (lawsuitNumber, error) {
	if len(lawsuit) != 22 {
		return lawsuitNumber{}, errors.New("lawsuit does not respect format: 14.0732.0000013/2021-1")
	}

	return lawsuitNumber{
		number1: lawsuit[:2],
		number2: lawsuit[3:7],
		number3: lawsuit[8:15],
		number4: lawsuit[16:20],
	}, nil

}

func existLawsuit(htmlPgSrc *html.Node) bool {
	noReturn := htmlquery.Find(htmlPgSrc, xpathReturn)

	if len(noReturn) > 0 {
		tReturn := htmlquery.InnerText(noReturn[0])
		if tReturn == textReturn {
			return false
		}
	}

	return true
}
