package controladores

import (
	"covid-19API/modelos"
	"covid-19API/utils"
	"fmt"
	"net/http"
	"time"
)

var urlCovid19Deaths = "https://raw.githubusercontent.com/CSSEGISandData/COVID-19/master/csse_covid_19_data/csse_covid_19_time_series/time_series_covid19_deaths_global.csv"

// GetCasosMuertos depending on query
var GetCasosMuertos = func(w http.ResponseWriter, r *http.Request) {
	var errorCasos error
	fmt.Println("Valor total:", modelos.PaisesCasosMuertos)
	fmt.Println("len(modelos.PaisesCasosMuertos):", len(modelos.PaisesCasosMuertos))

	switch { // missing expression means "true"

	case len(modelos.PaisesCasosMuertos) == 0 || !(modelos.PaisesCasosMuertos[0].ActualizacionDia.Before(time.Now().UTC())):

		modelos.PaisesCasosMuertos, errorCasos = utils.UpdateVariable(urlCovid19Deaths, modelos.PaisesCasosMuertos)
		if errorCasos != nil {
			response := utils.Message("Internal Server Error!")
			utils.Respond(w, response, 500)
			return
		}
		response := utils.Message("Casos Muertos")
		response["paises"] = modelos.PaisesCasosMuertos
		utils.Respond(w, response, 200)
	default:
		response := utils.Message("Casos Default Muertos")
		response["paises"] = modelos.PaisesCasosMuertos
		utils.Respond(w, response, 200)
	}
}
