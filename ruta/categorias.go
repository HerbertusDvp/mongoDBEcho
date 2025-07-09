package ruta

import (
	"context"
	"encoding/json"
	"fmt"
	"mongoEcho/database"
	"mongoEcho/dto"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CategoriaPost(c echo.Context) error {
	var body dto.CategoriaDto

	if err := json.NewDecoder(c.Request().Body).Decode(&body); err != nil {
		respuesta := map[string]string{
			"estado":  "error",
			"mensaje": "ocurrio un error",
		}

		//c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		//c.Response().WriteHeader(http.StatusBadRequest)
		//return json.NewEncoder(c.Response()).Encode(respuesta)

		return c.JSON(http.StatusBadRequest, respuesta)
	}

	if len(body.Nombre) == 0 {
		respuesta := map[string]string{
			"estado":  "error",
			"mensaje": "No hay nombre",
		}

		return c.JSON(http.StatusBadRequest, respuesta)
	}

	// Guardar en la bd

	registro := bson.D{
		{"nombre", body.Nombre},
		{"ice", "Luking to mai ais"},
	}

	database.CategoriaColeccion.InsertOne(context.TODO(), registro)

	// Respuesta

	respuesta := map[string]string{
		"estatus": "ok",
		"mensaje": "Se creó correctamente",
	}

	return c.JSON(http.StatusCreated, respuesta)
}

func CategoriaGet(c echo.Context) error {
	findOptions := options.Find()
	//Consulta normal
	//cursor, err := database.CategoriaColeccion.Find(context.TODO(), bson.D{})

	//Consulta ordenada                                                       //Para ordenar            //-1 -> Desc
	cursor, err := database.CategoriaColeccion.Find(context.TODO(), bson.D{}, findOptions.SetSort(bson.D{{"_id", -1}}))

	if err != nil {
		fmt.Println("Error de cursor")
		panic(err)
	}

	var results []bson.M

	if err = cursor.All(context.TODO(), &results); err != nil {
		fmt.Println("Error: 2do if cursor.All de cursor")
		panic(err)

	}

	return c.JSON(http.StatusOK, results)
}
