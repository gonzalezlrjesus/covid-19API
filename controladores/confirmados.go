package controladores

import (
	"covid-19API/modelos"
	"covid-19API/utils"
	"net/http"
	"time"
)

// GetCasosConfirmados depending on query
var GetCasosConfirmados = func(w http.ResponseWriter, r *http.Request) {
	var errorCasos error
	switch {
	case len(modelos.PaisesCasosConfirmados) == 0 || !(modelos.PaisesCasosConfirmados[0].ActualizacionDia.Before(time.Now().UTC())):
		modelos.PaisesCasosConfirmados, errorCasos = utils.UpdateVariable(modelos.URLCovid19Confirmed, modelos.PaisesCasosConfirmados)
		if errorCasos != nil {
			response := utils.Message("Internal Server Error!")
			utils.Respond(w, response, 500)
			return
		}
		response := utils.Message("Casos Confirmados")
		response["paises"] = modelos.PaisesCasosConfirmados
		utils.Respond(w, response, 200)
	default:
		response := utils.Message("Casos Default Confirmados")
		response["paises"] = modelos.PaisesCasosConfirmados
		utils.Respond(w, response, 200)

	}
}
