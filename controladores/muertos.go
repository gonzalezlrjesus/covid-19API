package controladores

import (
	"covid-19API/modelos"
	"covid-19API/utils"
	"fmt"
	"net/http"
	"time"
)

var urlCovid19Deaths = "https://raw.githubusercontent.com/CSSEGISandData/COVID-19/master/csse_covid_19_data/csse_covid_19_time_series/time_series_covid19_deaths_global.csv"

var paisesCasosMuertos []modelos.Pais

// GetCasosMuertos depending on query
var GetCasosMuertos = func(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Valor total:", paisesCasosMuertos)
	fmt.Println("len(paisesCasosMuertos):", len(paisesCasosMuertos))

	switch { // missing expression means "true"

	case len(paisesCasosMuertos) == 0 || !(paisesCasosMuertos[0].ActualizacionDia.Before(time.Now().UTC())):
		paisesCasosMuertos, error := utils.UpdateVariable(urlCovid19Deaths, paisesCasosMuertos)
		if error != nil {
			response := utils.Message("Internal Server Error!")
			utils.Respond(w, response, 500)
			return
		}
		response := utils.Message("Casos Muertos")
		response["paises"] = paisesCasosMuertos
		utils.Respond(w, response, 200)
	default:
		response := utils.Message("Casos Default Muertos")
		response["paises"] = paisesCasosMuertos
		utils.Respond(w, response, 200)
	}
}
