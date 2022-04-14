package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

func CrearBase() {
	db, err := sql.Open("postgres", "user=postgres host=localhost dbname=postgres sslmode=disable")
	if err != nil {
		log.Fatal(err)
	} else {
		_, err := db.Exec(`drop database if exists gestiontarjetas;`)
		if err != nil {
			log.Fatal(err)

		} else {
			_, err := db.Exec(`create database gestiontarjetas;`)
			if err != nil {
				log.Fatal(err)
			} else {
				fmt.Printf("\n ------------- BASE CREADA -------------\n\n")
			}
		}
	}
}
