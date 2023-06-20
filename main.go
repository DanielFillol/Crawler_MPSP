package main

import (
	"Crawler_MPSP/CSV"
	"Crawler_MPSP/Crawler"
	"fmt"
	"log"
	"strconv"
	"time"
)

type SearchLawsuits struct {
	LawsuitNumberMP string
	NameLawsuitMP   string
	DocumentNumber  []string
}

func main() {
	start1 := time.Now()

	driver, err := Crawler.SeleniumWebDriver()
	if err != nil {
		fmt.Println(err)
	}

	defer driver.Close()

	var suits []Crawler.Lawsuit
	for i, lawsuitNumber := range lawsuitNumbers {
		start2 := time.Now()

		lawsuit, err := Crawler.Craw(driver, lawsuitNumber.LawsuitNumberMP, lawsuitNumber.NameLawsuitMP, lawsuitNumber.DocumentNumber[0])
		if err != nil {
			log.Println(err)
		}
		suits = append(suits, lawsuit)

		t1 := time.Since(start1).String()
		t2 := time.Since(start2).String()
		m := int(time.Since(start1).Seconds()) / (i + 1)
		r := int(time.Since(start1).Seconds()) % (i + 1)
		md := strconv.Itoa(m) + "." + strconv.Itoa(r)
		fmt.Printf("processado %v | tempo: %v%v | total: %v%v | média: %vs \n", i+1, t2[0:4], t2[len(t2)-1:], t1[0:4], t1[len(t1)-1:], md)
	}

	err = CSV.WriteCSV(suits)
	if err != nil {
		fmt.Println(err)
	}

}

var lawsuitNumbers = []SearchLawsuits{
	{LawsuitNumberMP: "13.0004.0001104/2012-4", NameLawsuitMP: "EDSON DE OLIVEIRA PEIXOTO", DocumentNumber: []string{"28522778833", "33086689"}},
}
