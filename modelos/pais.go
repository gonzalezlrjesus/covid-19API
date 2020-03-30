package modelos

import (
	"time"
)

// Pais estructura de los paises afectados
type Pais struct {
	ProvinceState    string    `json:"Province/State"`
	CountryRegion    string    `json:"Country/Region"`
	Latitud          string    `json:"Lat"`
	Longitud         string    `json:"Long"`
	ActualizacionDia time.Time `json:"actualizacion_dia"`
	Dias             []Dia
}

var (
	// PaisesCasosMuertos variable
	PaisesCasosMuertos []Pais
	// PaisesCasosConfirmados variable
	PaisesCasosConfirmados []Pais
	// PaisesCasosRecuperados variable
	PaisesCasosRecuperados []Pais
)

var (
	// URLCovid19Deaths variable
	URLCovid19Deaths = "https://raw.githubusercontent.com/CSSEGISandData/COVID-19/master/csse_covid_19_data/csse_covid_19_time_series/time_series_covid19_deaths_global.csv"
	// URLCovid19Recovered variable
	URLCovid19Recovered = "https://raw.githubusercontent.com/CSSEGISandData/COVID-19/master/csse_covid_19_data/csse_covid_19_time_series/time_series_covid19_recovered_global.csv"
	// URLCovid19Confirmed variable
	URLCovid19Confirmed = "https://raw.githubusercontent.com/CSSEGISandData/COVID-19/master/csse_covid_19_data/csse_covid_19_time_series/time_series_covid19_confirmed_global.csv"
)
