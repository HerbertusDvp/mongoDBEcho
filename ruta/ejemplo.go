package ruta

import (
	"encoding/json"
	"net/http"

	"mongoEcho/dto"

	"github.com/labstack/echo/v4"
)

func EjemploGet(c echo.Context) error {
	return c.String(http.StatusOK, "Hola mundo con GET")
}

func GetParametros(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, "GET | id="+id)
}

func GetQueryString(c echo.Context) error {
	id := c.QueryParam("id")
	name := c.QueryParam("name")
	return c.String(http.StatusOK, "GET | id="+id+" | name="+name)
}

func GetJSON(c echo.Context) error {

	respuesta := map[string]string{
		"estado":  "ok",
		"mensaje": "Metodo Get status",
	}

	//return c.JSON(200, respuesta)
	return c.JSON(http.StatusOK, respuesta)
}

func EjemploPost(c echo.Context) error { // Sin datos
	return c.String(http.StatusOK, "Hola mundo con POST")
}

func PostDatos(c echo.Context) error {
	var body dto.CategoriaDto
	if err := json.NewDecoder(c.Request().Body).Decode(&body); err != nil {
		respuesta := map[string]string{
			"estado":  "error",
			"mensaje": "Ocurrio un error",
		}

		return c.JSON(400, respuesta)
	}

	respuesta := map[string]string{
		"estado":  "ok",
		"mensaje": "Metodo post",
		"nombre":  body.Nombre,
	}

	return c.JSON(200, respuesta)

}
func EjemploPut(c echo.Context) error {
	return c.String(http.StatusOK, "Hola mundo PUT")
}
func EjemploDelete(c echo.Context) error {
	return c.String(http.StatusOK, "Hola mundo DELETE")
}
