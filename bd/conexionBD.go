package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*MongoCN*/
var MongoCN=ConectarBD()
var clientOptions = options.Client().ApplyURI("mongodb+srv://lemeze:3011gake@twitter.x3krj.mongodb.net/test?retryWrites=true&w=majority")

/*conectar BD mongo*/
func ConectarBD() *mongo.Client  {
    client, err := mongo.Connect(context.TODO(), clientOptions)
    if err != nil {
        log.Fatal(err.Error())
        return client
    }
    err = client.Ping(context.TODO(), nil)
    if err != nil {
        log.Fatal(err.Error())
        return client
    }
    log.Println("Conexion exitosa con la BD");
    return client
}

func ChequeoConnection() int {
    err := MongoCN.Ping(context.TODO(), nil)
    if err != nil {
        return 0
    }
    return 1
}