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

// PaisesCasosMuertos variable
var PaisesCasosMuertos []Pais
