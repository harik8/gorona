package gorona

import "math"

// CaseState : Struct Case State
type CaseState struct {
	Country             string  `json:"country" header:"Country"`
	Cases               int64   `json:"cases" header:"Cases"`
	TodayCases          int64   `json:"todayCases" header:"Today Cases"`
	Deaths              int64   `json:"deaths" header:"Deaths"`
	TodayDeaths         int64   `json:"todayDeaths" header:"Today Deaths"`
	Recovered           int64   `json:"recovered" header:"Recovered"`
	Active              int64   `json:"active" header:"Active"`
	Critical            int64   `json:"critical" header:"Critical"`
	CasesPerOneMillion  float64 `json:"casesPerOneMillion" header:"Cases Per One Million"`
	DeathsPerOneMillion float64 `json:"deathsPerOneMillion" header:"Deaths Per One Million"`
}

// CaseStates is a list of cases
type CaseStates []CaseState

// Cases is the total number of cases
func (cs CaseStates) Cases() int64 {
	var total int64
	for _, s := range cs {
		total += s.Cases
	}
	return total
}

// TodayCases is the total number of today cases
func (cs CaseStates) TodayCases() int64 {
	var total int64
	for _, s := range cs {
		total += s.TodayCases
	}
	return total
}

// Deaths is the total number of deaths
func (cs CaseStates) Deaths() int64 {
	var total int64
	for _, s := range cs {
		total += s.Deaths
	}
	return total
}

// TodayDeaths is the total number of today deaths
func (cs CaseStates) TodayDeaths() int64 {
	var total int64
	for _, s := range cs {
		total += s.TodayDeaths
	}
	return total
}

// Recovered is the total number of recovered
func (cs CaseStates) Recovered() int64 {
	var total int64
	for _, s := range cs {
		total += s.Recovered
	}
	return total
}

// Active is the total number of active
func (cs CaseStates) Active() int64 {
	var total int64
	for _, s := range cs {
		total += s.Active
	}
	return total
}

// Critical is the total number of active
func (cs CaseStates) Critical() int64 {
	var total int64
	for _, s := range cs {
		total += s.Critical
	}
	return total
}

// CasesPerOneMillion is the total number of cases per million
func (cs CaseStates) CasesPerOneMillion() float64 {
	var total float64
	for _, s := range cs {
		total += s.CasesPerOneMillion
	}
	return math.Round(total*10)/10
}

// DeathsPerOneMillion is the total number of deaths per million
func (cs CaseStates) DeathsPerOneMillion() float64 {
	var total float64
	for _, s := range cs {
		total += s.DeathsPerOneMillion
	}
	return math.Round(total*10)/10
}
