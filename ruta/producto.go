package ruta

import (
	"context"
	"encoding/json"
	"fmt"
	"mongoEcho/database"
	"mongoEcho/dto"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ProductoPost(c echo.Context) error {

	var body dto.ProductoDto

	if err := json.NewDecoder(c.Request().Body).Decode(&body); err != nil {
		respuesta := map[string]string{
			"estado":  "error",
			"mensaje": "Error en NewDecoder de ProductosPost",
		}

		return c.JSON(400, respuesta)
	}

	if len(body.Nombre) == 0 {
		respuesta := map[string]string{
			"estado":  "error",
			"mensaje": "Error: Valor vacío",
		}

		return c.JSON(400, respuesta)

	}

	//Sobre la base de datos

	CategoriaID, err := primitive.ObjectIDFromHex(body.CategoriaID)

	if err != nil {
		fmt.Println("Error con Asignacion de CategoariaID en ProductoPOST")
	}

	registro := bson.D{
		{"nombre", body.Nombre},
		{"precio", body.Precio},
		{"stock", body.Stock},
		{"descripcion", body.Descripcion},
		{"categoria_id", CategoriaID},
	}

	database.ProductoColeccion.InsertOne(context.TODO(), registro)

	respuesta := map[string]string{
		"estado":  "ok",
		"mensaje": "Accion correcta",
	}

	return c.JSON(200, respuesta)

}

func ProductoGet(c echo.Context) error {

	findOptions := options.Find()
	cursor, err := database.ProductoColeccion.Find(context.TODO(), bson.D{}, findOptions.SetSort(bson.D{{"_id", -1}}))

	if err != nil {
		fmt.Println("error del cursor en productoGet")
	}

	var results []bson.M

	if err = cursor.All(context.TODO(), &results); err != nil {
		fmt.Println("Error de cursor.All en ProductoGet")
	}

	return c.JSON(200, results)
}

func ProductoGetJoin(c echo.Context) error {

	pipeline := []bson.M{
		{"$match": bson.M{}}, //Coleccion ext.  //campo de col princ                                  //alias
		{"$lookup": bson.M{"from": "Categorias", "localField": "categoria_id", "foreignField": "_id", "as": "categorias"}},
		{"$sort": bson.M{"_id": -1}},
	}

	//findOptions := options.Find()
	cursor, err := database.ProductoColeccion.Aggregate(context.TODO(), pipeline)

	if err != nil {
		fmt.Println("error del cursor en productoGetJoin")
	}

	var results []bson.M

	if err = cursor.All(context.TODO(), &results); err != nil {
		fmt.Println("Error de cursor.All en ProductoGet")
	}

	return c.JSON(200, results)
}
