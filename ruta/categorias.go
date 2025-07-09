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
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func CategoriaGetByID(c echo.Context) error {
	objectID, err := primitive.ObjectIDFromHex(c.Param("id"))

	if err != nil {
		fmt.Println("Error con objectID: ", err)
	}

	var result bson.M

	if err := database.CategoriaColeccion.FindOne(context.TODO(), bson.M{"_id": objectID}).Decode(&result); err != nil {
		respuesta := map[string]string{
			"estado":  "error",
			"mensaje": "Error en el FindOne: ",
		}

		return c.JSON(400, respuesta)
	}

	return c.JSON(200, result)
}

func CategoriaSetByID(c echo.Context) error {

	var body dto.CategoriaDto

	err := json.NewDecoder(c.Request().Body).Decode(&body)

	if err != nil {
		respuesta := map[string]string{
			"estado":  "error",
			"mensaje": "Error con el NewDecoder",
		}

		return c.JSON(400, respuesta)
	}

	if len(body.Nombre) == 0 {
		respuesta := map[string]string{
			"estado":  "error",
			"mensaje": "Error: Nombre vacio",
		}

		return c.JSON(400, respuesta)
	}

	var result bson.M
	objectID, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err := database.CategoriaColeccion.FindOne(context.TODO(), bson.M{"_id": objectID}).Decode(&result); err != nil {
		respuesta := map[string]string{
			"estado":  "error",
			"mensaje": "Error en finOne",
		}

		return c.JSON(400, respuesta)
	}

	registro := make(map[string]interface{})

	registro["nombre"] = body.Nombre
	registro["ice"] = "Good time"

	updateString := bson.M{
		"$set": registro,
	}

	database.CategoriaColeccion.UpdateOne(context.TODO(), bson.M{"_id": bson.M{"$eq": objectID}}, updateString)

	respuesta := map[string]string{
		"estado":  "ok",
		"mensaje": "Registro modificado",
	}

	return c.JSON(200, respuesta)
}

func CategoriaDelete(c echo.Context) error {

	var result bson.M

	objectID, err := primitive.ObjectIDFromHex(c.Param("id"))

	if err != nil {
		fmt.Println("Error en object id")
	}

	if err := database.CategoriaColeccion.FindOne(context.TODO(), bson.M{"_id": objectID}).Decode(&result); err != nil {
		respuesta := map[string]string{
			"estado":  "ok",
			"mensaje": "Error con FindOne de CategoriaDelete",
		}

		c.JSON(400, respuesta)
	}
	// Where id = objectID
	database.CategoriaColeccion.DeleteOne(context.TODO(), bson.M{"_id": objectID})

	//Respuesta exitosa
	respuesta := map[string]string{
		"estado":  "ok",
		"mensaje": "Elemento eliminado",
	}

	return c.JSON(400, respuesta)
}
