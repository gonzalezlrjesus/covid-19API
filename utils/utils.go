package utils

import (
	"bufio"
	"covid-19API/modelos"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"time"
)

// Message function wrap message to users
func Message(message string) map[string]interface{} {
	return map[string]interface{}{"message": message}
}

// Respond function to responde request
func Respond(w http.ResponseWriter, data map[string]interface{}, status uint) {
	w.Header().Add("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(data)
}

func downloadFile(filePath string) (string, error) {
	response, error := http.Get(filePath)
	if error != nil {
		return "", error
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		fmt.Println("bad status: ", response.Status)
	}

	file, err := os.Create(path.Base(response.Request.URL.String()))
	if err != nil {
		return "", err
	}
	defer file.Close()

	// Use io.Copy to just dump the response body to the file. This supports huge files
	_, err = io.Copy(file, response.Body)
	if err != nil {
		return "", err
	}
	return path.Base(response.Request.URL.String()), nil
}

func readFile(filePath string) (*[][]string, error) {
	csvFile, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	reader := csv.NewReader(bufio.NewReader(csvFile))

	records, err := reader.ReadAll()
	if err == io.EOF {
		return nil, err
	} else if err != nil {
		return nil, err
	}
	return &records, err
}

func parseCSVFileToJSON(paises *[]modelos.Pais, records *[][]string) {
	var dias []modelos.Dia
	for indexFilas := 1; indexFilas < len(*records); indexFilas++ {
		dias = nil
		for indexColumnas := 4; indexColumnas < len((*records)[0]); indexColumnas++ {

			dias = append(dias, modelos.Dia{
				Fecha:    (*records)[0][indexColumnas],
				Cantidad: (*records)[indexFilas][indexColumnas],
			})

		}
		*paises = append(*paises, modelos.Pais{
			ProvinceState:    (*records)[indexFilas][0],
			CountryRegion:    (*records)[indexFilas][1],
			Latitud:          (*records)[indexFilas][2],
			Longitud:         (*records)[indexFilas][3],
			ActualizacionDia: time.Now().UTC(),
			Dias:             dias,
		})

	}

}

// UpdateVariable actualiza la variable donde se almacenan los casos en memoria
func UpdateVariable(filePATH string, paisesCasos []modelos.Pais) ([]modelos.Pais, error) {
	fileName, err := downloadFile(filePATH)
	if err != nil {
		return nil, err
	}
	CSVFile, err := readFile(fileName)
	if err != nil {
		return nil, err
	}
	parseCSVFileToJSON(&paisesCasos, CSVFile)
	return paisesCasos, nil
}
