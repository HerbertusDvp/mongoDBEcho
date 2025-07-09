package ruta

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"mongoEcho/dto"

	"github.com/labstack/echo/v4"
)

func EjemploGet(c echo.Context) error {

	respuesta := map[string]string{
		"estado":   "ok",
		"mensaje":  "Metodo get",
		"cabecera": c.Request().Header.Get("Authorization"),
	}

	fmt.Println("Ejecuta controlador EjemploGet")

	return c.JSON(200, respuesta)
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

func EjemploUpload(c echo.Context) error {

	file, err := c.FormFile("foto")

	if err != nil {
		fmt.Println("Error: ", err)
		return err
	}

	src, err := file.Open()
	if err != nil {
		return err
	}

	defer src.Close()
	//Nombrado del archivo

	var extension = strings.Split(file.Filename, ".")[1]
	time := strings.Split(time.Now().String(), " ")
	foto := string(time[4][6:14] + "." + extension)
	var archivo string = "public/uploads/fotos/" + foto

	dst, err := os.Create(archivo)

	if err != nil {
		return err
	}

	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	respuesta := map[string]string{
		"estado":  "ok",
		"mensaje": "Todo ok",
		"nombre":  foto,
	}

	//c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	return c.JSON(200, respuesta)

	//c.Response().WriteHeader(http.StatusOK)
	//return json.NewEncoder(c.Response()).Encode(respuesta)

}
