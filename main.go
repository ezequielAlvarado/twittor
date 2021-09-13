package main

// "fmt"
import (
	"log"
	"twittor/bd"
	"twittor/handlers"
)

func main(){
	if(bd.ChequeoConnection()==0) {
		log.Fatal("Sin conexion a la BD")
		return
	}
	handlers.Manejadores()
	
	
}