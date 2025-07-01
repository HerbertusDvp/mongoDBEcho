package ruta

import (
	"net/http"

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

func EjemploPost(c echo.Context) error {
	return c.String(http.StatusOK, "Hola mundo con POST")
}
func EjemploPut(c echo.Context) error {
	return c.String(http.StatusOK, "Hola mundo PUT")
}
func EjemploDelete(c echo.Context) error {
	return c.String(http.StatusOK, "Hola mundo DELETE")
}
