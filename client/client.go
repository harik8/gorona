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

const (
	url = "https://corona.lmao.ninja/countries/"
)

func getCoronaStatus(cases map[string]interface{}) gorona.CaseState {
	cs := gorona.CaseState{}

	_, cou := cases["country"].(string)
	_, cas := cases["cases"].(float64)
	_, tca := cases["todayCases"].(float64)
	_, dea := cases["deaths"].(float64)
	_, tde := cases["todayDeaths"].(float64)
	_, rec := cases["recovered"].(float64)
	_, act := cases["active"].(float64)
	_, cri := cases["critical"].(float64)
	_, cpm := cases["casesPerOneMillion"].(float64)


	if cou {
		cs.Country = cases["country"].(string)
	}
	if cas {
		cs.Cases = cases["cases"].(float64)
	}
	if tca {
		cs.TodayCases = cases["todayCases"].(float64)
	}
	if dea {
		cs.Deaths = cases["deaths"].(float64)	
	}
	if tde {
		cs.TodayDeaths = cases["todayDeaths"].(float64)
	}
	if rec {
		cs.Recovered = cases["recovered"].(float64)
	}
	if act {
		cs.Active = cases["active"].(float64)
	}
	if cri {
		cs.Critical = cases["critical"].(float64)
	}
	if cpm {
		cs.CasesPerOneMillion = cases["casesPerOneMillion"].(float64)
	}
	return cs
}

func getPrinter() *tableprinter.Printer {
	printer := tableprinter.New(os.Stdout)

	printer.BorderTop, printer.BorderBottom, printer.BorderLeft, printer.BorderRight = true, true, true, true
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
	coronaCases := make(map[string]interface{})
	err = json.Unmarshal(cases, &coronaCases)
	if err != nil {
		log.Fatalln(err)
	}
	cs := getCoronaStatus(coronaCases)
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
	var coronaCases interface{}
	err = json.Unmarshal(cases, &coronaCases)
	if err != nil {
		log.Fatalln(err)
	}
	cs := coronaCases.([]interface{})
	for c:=0; c<len(cs); c++ {
		cm := cs[c].(map[string]interface{})
		gs = append(gs, getCoronaStatus(cm))
	}
	getPrinter().Print(gs)
}