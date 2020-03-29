package controladores

import (
	"covid-19API/modelos"
	"covid-19API/utils"
	"net/http"
	"time"
)

var urlCovid19Confirmed = "https://raw.githubusercontent.com/CSSEGISandData/COVID-19/master/csse_covid_19_data/csse_covid_19_time_series/time_series_covid19_confirmed_global.csv"

var paisesCasosConfirmados []modelos.Pais

// GetCasosConfirmados depending on query
var GetCasosConfirmados = func(w http.ResponseWriter, r *http.Request) {
	switch { // missing expression means "true"
	case len(paisesCasosConfirmados) == 0 || !(paisesCasosConfirmados[0].ActualizacionDia.Before(time.Now().UTC())):
		paisesCasosConfirmados, error := utils.UpdateVariable(urlCovid19Confirmed, paisesCasosConfirmados)
		if error != nil {
			response := utils.Message("Internal Server Error!")
			utils.Respond(w, response, 500)
			return
		}
		response := utils.Message("Casos Confirmados")
		response["paises"] = paisesCasosConfirmados
		utils.Respond(w, response, 200)
	default:
		response := utils.Message("Casos Default Confirmados")
		response["paises"] = paisesCasosConfirmados
		utils.Respond(w, response, 200)

	}
}
