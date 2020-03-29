package controladores

import (
	"covid-19API/modelos"
	"covid-19API/utils"
	"net/http"
	"time"
)

var urlCovid19Recovered = "https://raw.githubusercontent.com/CSSEGISandData/COVID-19/master/csse_covid_19_data/csse_covid_19_time_series/time_series_covid19_recovered_global.csv"

var paisesCasosRecuperados []modelos.Pais

// GetCasosRecuperados depending on query
var GetCasosRecuperados = func(w http.ResponseWriter, r *http.Request) {
	switch { // missing expression means "true"
	case len(paisesCasosRecuperados) == 0 || !(paisesCasosRecuperados[0].ActualizacionDia.Before(time.Now().UTC())):
		paisesCasosRecuperados, error := utils.UpdateVariable(urlCovid19Recovered, paisesCasosRecuperados)
		if error != nil {
			response := utils.Message("Internal Server Error!")
			utils.Respond(w, response, 500)
			return
		}
		response := utils.Message("Casos Recuperados")
		response["paises"] = paisesCasosRecuperados
		utils.Respond(w, response, 200)
	default:
		response := utils.Message("Casos Default Recuperados")
		response["paises"] = paisesCasosRecuperados
		utils.Respond(w, response, 200)
	}
}
