package main

import (
	"mongoEcho/database"
	"mongoEcho/ruta"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var prefijo string = "/api/v1/"

func main() {
	e := echo.New()
	e.Static("/public", "public")

	// Conexion a la db
	database.ComprobarConexion()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// CORS default
	// Allows requests from any origin wth GET, HEAD, PUT, POST or DELETE method.
	// e.Use(middleware.CORS())

	// CORS restricted
	// Allows requests from any `https://labstack.com` or `https://labstack.net` origin
	// wth GET, PUT, POST or DELETE method.
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"https://labstack.com", "https://labstack.net"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	e.GET(prefijo+"ejemplo", ruta.EjemploGet)
	e.GET(prefijo+"ejemplo/:id", ruta.GetParametros)
	e.GET(prefijo+"queryString", ruta.GetQueryString)
	e.GET(prefijo+"ejemploJson", ruta.GetJSON)

	e.POST(prefijo+"ejemplo", ruta.EjemploPost)
	e.POST(prefijo+"ejemploDatos", ruta.PostDatos)
	e.PUT(prefijo+"ejemplo", ruta.EjemploPut)
	e.DELETE(prefijo+"ejemplo", ruta.EjemploDelete)

	e.POST(prefijo+"upload", ruta.EjemploUpload) // Carga un archivo

	e.POST(prefijo+"categorias", ruta.CategoriaPost)         // Insert
	e.GET(prefijo+"categorias", ruta.CategoriaGet)           // Listar / consulta
	e.GET(prefijo+"categorias/:id", ruta.CategoriaGetByID)   // coculta por id
	e.PUT(prefijo+"categorias/:id", ruta.CategoriaSetByID)   // modifica por id
	e.DELETE(prefijo+"categorias/:id", ruta.CategoriaDelete) // borra por id

	e.GET(prefijo+"productos", ruta.ProductoGet)
	e.GET(prefijo+"productosJoin", ruta.ProductoGetJoin)
	e.POST(prefijo+"productos", ruta.ProductoPost)

	errorVariables := godotenv.Load()
	if errorVariables != nil {

		panic(errorVariables)

	}

	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))

}
