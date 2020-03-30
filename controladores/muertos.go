package controladores

import (
	"covid-19API/modelos"
	"covid-19API/utils"
	"net/http"
	"time"
)

// GetCasosMuertos encargada de devolver en formato json los casos muertos hasta la fecha.
var GetCasosMuertos = func(w http.ResponseWriter, r *http.Request) {
	var errorCasos error

	switch {
	case len(modelos.PaisesCasosMuertos) == 0 || !(modelos.PaisesCasosMuertos[0].ActualizacionDia.Before(time.Now().UTC())):

		modelos.PaisesCasosMuertos, errorCasos = utils.UpdateVariable(modelos.URLCovid19Deaths, modelos.PaisesCasosMuertos)
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
