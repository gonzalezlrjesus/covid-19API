package controladores

import (
	"covid-19API/modelos"
	"covid-19API/utils"
	"net/http"
	"time"
	"fmt"
)

var urlCovid19Deaths = "https://raw.githubusercontent.com/CSSEGISandData/COVID-19/master/csse_covid_19_data/csse_covid_19_time_series/time_series_covid19_deaths_global.csv"

// GetCasosMuertos depending on query
var GetCasosMuertos = func(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Valor total:" , modelos.PaisesCasosMuertos)
	fmt.Println("len(modelos.PaisesCasosMuertos):" , len(PaisesCasosMuertos))
	fmt.Println("modelos.PaisesCasosMuertos[0].ActualizacionDia.Before(time.Now().UTC())" , modelos.PaisesCasosMuertos[0].ActualizacionDia.Before(time.Now().UTC()))
	switch { // missing expression means "true"

	case len(modelos.PaisesCasosMuertos) == 0 || !(modelos.PaisesCasosMuertos[0].ActualizacionDia.Before(time.Now().UTC())):
		modelos.PaisesCasosMuertos, error := utils.UpdateVariable(urlCovid19Deaths, modelos.PaisesCasosMuertos)
		if error != nil {
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
