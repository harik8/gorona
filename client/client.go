package client

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/kataras/tablewriter"
	"github.com/landoop/tableprinter"

	gorona "github.com/harik8/gorona/gorona"
)

const url = "https://corona.lmao.ninja/countries/"

func getPrinter() *tableprinter.Printer {
	printer := tableprinter.New(os.Stdout)
	printer.BorderTop = true
	printer.BorderBottom = true
	printer.BorderLeft = true
	printer.BorderRight = true
	printer.CenterSeparator = "│"
	printer.ColumnSeparator = "│"
	printer.RowSeparator = "─"
	printer.HeaderFgColor = tablewriter.FgGreenColor
	return printer
}

// GetCountry : Get by country
func GetCountry(country string) {

	req, err := http.Get(url + country)
	if err != nil {
		log.Fatalln(err)
	}
	defer req.Body.Close()

	cases, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Fatalln(err)
	}

	cs := gorona.CaseState{}
	err = json.Unmarshal(cases, &cs)
	if err != nil {
		log.Fatalln(err)
	}

	getPrinter().Print(cs)

}

// GetCountries : GetCountries()
func GetCountries() {

	gs := []gorona.CaseState{}

	req, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	defer req.Body.Close()

	cases, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Fatalln(err)
	}

	err = json.Unmarshal(cases, &gs)
	if err != nil {
		log.Fatalln(err)
	}

	getPrinter().Print(gs)

}
