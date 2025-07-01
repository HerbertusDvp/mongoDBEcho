package main

import (
	"mongoEcho/ruta"
	"os"

	"github.com/joho/godotenv"
	echo "github.com/labstack/echo/v4"
)

var prefijo string = "/api/v1/"

func main() {
	e := echo.New()

	e.GET(prefijo+"ejemplo", ruta.EjemploGet)
	e.GET(prefijo+"ejemplo/:id", ruta.GetParametros)
	e.GET(prefijo+"queryString", ruta.GetQueryString)
	e.GET(prefijo+"ejemploJson", ruta.GetJSON)

	e.POST(prefijo+"ejemplo", ruta.EjemploPost)
	e.PUT(prefijo+"ejemplo", ruta.EjemploPut)
	e.DELETE(prefijo+"ejemplo", ruta.EjemploDelete)

	errorVariables := godotenv.Load()
	if errorVariables != nil {

		panic(errorVariables)

	}

	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))

}
