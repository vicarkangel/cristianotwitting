package main

import (
	"log"

	"github.com/vicarkangel/cristianotwitting/bd"
	"github.com/vicarkangel/cristianotwitting/handlers"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexion a la BD....")
		return
	}
	handlers.Manejadores()
}
