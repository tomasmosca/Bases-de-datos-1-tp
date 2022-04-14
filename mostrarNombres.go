package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type comercio struct {
	nrocomercio                               int
	nombre, domicilio, codigopostal, telefono string
}

func MostrarNombres() {
	fmt.Printf("\n ------------- NOMBRES DE COMERCIOS -------------\n\n")
	db, err := sql.Open("postgres", "user=postgres host=localhost dbname=gestiontarjetas sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query(`select * from comercio`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var c comercio
	for rows.Next() {
		if err := rows.Scan(&c.nrocomercio, &c.nombre, &c.domicilio, &c.codigopostal, &c.telefono); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Nombre : %v \n", c.nombre)
	}
}
