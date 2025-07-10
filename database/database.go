package database

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ClienteMongo = ConectarDB()
var MongoDB = "mongoBase" // Nombre de la base de datos

// Nombre de la "Tabla"
var CategoriaColeccion = ClienteMongo.Database(MongoDB).Collection("Categorias")
var ProductoColeccion = ClienteMongo.Database(MongoDB).Collection("Producto")
var clientOptions = options.Client().ApplyURI("mongodb://localhost:27017/" + MongoDB)

func ConectarDB() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err.Error(), "Error en mongo.connect")
		return client
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err.Error(), "Error del ping")
		return client
	}

	return client
}

func ComprobarConexion() int {
	err := ClienteMongo.Ping(context.TODO(), nil)

	if err != nil {
		fmt.Println("Conexion no etcitosa")
		return 0
	}
	fmt.Println("Conexion etcitosa")
	return 1

}
