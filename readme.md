# Covid-19 API en GO

Es una RESTFUL API por el cual se puede tener acceso a una serie de casos diarios (recuperados, muertos, confirmados) por paises.

## Endpoints Version 1

### /casos-confirmados

**Peticion:** 
```json
GET https://covid-19-api-go.herokuapp.com/v1/casos-confirmados
```

**Curl:**
```sh
curl -X GET "https://covid-19-api-go.herokuapp.com/v1/casos-confirmados"
```
**Navegador:**

[https://covid-19-api-go.herokuapp.com/v1/casos-confirmados](https://covid-19-api-go.herokuapp.com/v1/casos-confirmados)

### /casos-muertos

**Peticion:** 
```json
GET https://covid-19-api-go.herokuapp.com/v1/casos-muertos
```

**Curl:**
```sh
curl -X GET "https://covid-19-api-go.herokuapp.com/v1/casos-muertos"
```

**Navegador:** 

[https://covid-19-api-go.herokuapp.com/v1/casos-muertos](https://covid-19-api-go.herokuapp.com/v1/casos-muertos)

### /casos-recuperados
**Peticion:** 
```json
GET https://covid-19-api-go.herokuapp.com/v1/casos-recuperados
```

**Curl:**
```sh
curl -X GET "https://covid-19-api-go.herokuapp.com/v1/casos-recuperados"
```

**Navegador:** 

[https://covid-19-api-go.herokuapp.com/v1/casos-recuperados](https://covid-19-api-go.herokuapp.com/v1/casos-recuperados)


##### EJEMPLO. Si desea utilizar la API en su proyecto:

```javascript
fetch('https://covid-19-api-go.herokuapp.com/v1/casos-recuperados')
.then(response => response.json())
.then(data => {
    let Venezuela = data.paises.find((pais) => {
    return pais["Country/Region"] === 'Venezuela';
    });
    console.log(Venezuela);
})
```

```json
Province/State: ""
Country/Region: "Venezuela"
Lat: "6.4238"
Long: "-66.5897"
actualizacion_dia: "2020-03-30T12:25:31.630935758Z"
Dias: (68) [{…}, {…}, {…}, {…}, {…}, {…}, ...
```

#### Origen de los datos

https://github.com/CSSEGISandData/COVID-19

#### Licencia

MIT License 2020, gonzalezlrjesus.

Transitivamente desde el sitio de [the Johns Hopkins University Center for Systems Science and Engineering (JHU CSSE)](https://github.com/CSSEGISandData/COVID-19), los datos no pueden ser utilizados con fines comerciales.