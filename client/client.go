package client

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/harik8/gorona/gorona"
	"github.com/jedib0t/go-pretty/table"
)

const url = "https://corona.lmao.ninja/countries/"

// GetCountry : Get by country
func GetCountry(country string) {

	caseState := gorona.CaseState{}

	req, err := http.Get(url + country)
	if err != nil {
		log.Fatalln(err)
	}
	defer req.Body.Close()

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Fatalln(err)
	}

	err = json.Unmarshal(body, &caseState)
	if err != nil {
		log.Fatalln(err)
	}

	printCaseState(caseState)

}

// GetCountries : GetCountries()
func GetCountries() {

	caseStates := gorona.CaseStates{}

	req, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	defer req.Body.Close()

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Fatalln(err)
	}

	err = json.Unmarshal(body, &caseStates)
	if err != nil {
		log.Fatalln(err)
	}

	printCaseStates(caseStates)

}

func printCaseState(caseState gorona.CaseState) {
	printCaseStates(gorona.CaseStates{caseState})
}

func printCaseStates(caseStates gorona.CaseStates) {

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.SetStyle(table.StyleLight)

	t.AppendHeader(
		table.Row{
			"Country", "Cases", "Today Cases", "Death", "Today Deaths", "Recovered", "Active", "Critical", "Cases Per Million", "Deaths Per Million",
		},
	)

	for _, caseState := range caseStates {
		t.AppendRow(table.Row{
			caseState.Country,
			caseState.Cases,
			caseState.TodayCases,
			caseState.Deaths,
			caseState.TodayDeaths,
			caseState.Recovered,
			caseState.Active,
			caseState.Critical,
			caseState.CasesPerOneMillion,
			caseState.DeathsPerOneMillion,
		})
	}

	if len(caseStates) > 1 {
		t.AppendFooter(
			table.Row{
				"Totals",
				caseStates.Cases(),
				caseStates.TodayCases(),
				caseStates.Deaths(),
				caseStates.TodayDeaths(),
				caseStates.Recovered(),
				caseStates.Active(),
				caseStates.Critical(),
				caseStates.CasesPerOneMillion(),
				caseStates.DeathsPerOneMillion(),
			},
		)
	}

	t.Render()

}
