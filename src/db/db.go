package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (

	// Mongo Connection
	URI     = "mongodb://localhost:27017"
	TIMEOUT = 15

	//Mongo Parameters
	DB_NAME = "practica"
)

// Genera una conexion con la API de Mongo
func connect(uri string) (*mongo.Client, context.Context, context.CancelFunc, error) {

	ctx, cancel := context.WithTimeout(
		context.Background(),
		TIMEOUT*time.Second,
	)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))

	return client, ctx, cancel, err

}

func ObtenerColeccion(nameCollection string) (out [][]byte, err error) {

	client, ctx, cancel, err := connect(URI)

	if err != nil {
		cancel()
		return nil, err
	}

	defer client.Disconnect(ctx)

	database := client.Database(DB_NAME)
	collection := database.Collection(nameCollection)

	cursor, err := collection.Find(ctx, bson.D{})

	if err != nil {
		return nil, err
	}

	// Obtenemos todos los documentos de la coleccion y los convertimos en JSON's
	// para almacenarlos en una matriz de bytes y enviarla como respuesta hacia el cliente
	for cursor.Next(ctx) {

		document, err := bson.MarshalExtJSON(cursor.Current, false, false)

		if err != nil {
			return nil, err
		}

		out = append(out, document)

	}

	return out, nil

}

func InsertarDocumento(data interface{}, nameCollection string) error {

	client, ctx, cancel, err := connect(URI)

	if err != nil {
		cancel()
		return err
	}

	defer client.Disconnect(ctx)

	database := client.Database(DB_NAME)

	_, err = database.Collection(nameCollection).InsertOne(ctx, data)

	if err != nil {
		return err
	}

	return nil
}
