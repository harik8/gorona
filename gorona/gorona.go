package gorona

// CaseState : Struct Case State
type CaseState struct {
	Country            string `json:"country" header:"Country"`
	Cases              int64  `json:"cases" header:"Cases"`
	TodayCases         int64  `json:"todayCases" header:"Today Cases"`
	Deaths             int64  `json:"deaths" header:"Deaths"`
	TodayDeaths        int64  `json:"todayDeaths" header:"Today Deaths"`
	Recovered          int64  `json:"recovered" header:"Recovered"`
	Active             int64  `json:"active" header:"Active"`
	Critical           int64  `json:"critical" header:"Critical"`
	CasesPerOneMillion int64  `json:"casesPerOneMillion" header:"Cases Per One Million"`
}
