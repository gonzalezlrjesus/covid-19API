package controladores

import (
	"covid-19API/modelos"
	"covid-19API/utils"
	"net/http"
	"time"
)

// GetCasosRecuperados encargada de devolver en formato json los casos recuperados hasta la fecha.
var GetCasosRecuperados = func(w http.ResponseWriter, r *http.Request) {
	var errorCasos error
	switch {
	case len(modelos.PaisesCasosRecuperados) == 0 || !(modelos.PaisesCasosRecuperados[0].ActualizacionDia.Before(time.Now().UTC())):
		modelos.PaisesCasosRecuperados, errorCasos = utils.UpdateVariable(modelos.URLCovid19Recovered, modelos.PaisesCasosRecuperados)
		if errorCasos != nil {
			response := utils.Message("Internal Server Error!")
			utils.Respond(w, response, 500)
			return
		}
		response := utils.Message("Casos Recuperados")
		response["paises"] = modelos.PaisesCasosRecuperados
		utils.Respond(w, response, 200)
	default:
		response := utils.Message("Casos Default Recuperados")
		response["paises"] = modelos.PaisesCasosRecuperados
		utils.Respond(w, response, 200)
	}
}
