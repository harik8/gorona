package gorona

// CaseState : Struct Case State
type CaseState struct {
	Country string `header:"Country"`
	Cases float64 `header:"Cases"`
	TodayCases float64 `header:"Today Cases"`
	Deaths float64 `header:"Deaths"`
	TodayDeaths float64 `header:"Today Deaths"`
	Recovered float64 `header:"Recovered"`
	Active float64 `header:"Active"`
	Critical float64 `header:"Critical"`
	CasesPerOneMillion float64 `header:"Cases Per One Million"`
}