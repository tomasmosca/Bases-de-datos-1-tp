package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

func GenerarResumenes() {
	db, err := sql.Open("postgres", "user=postgres host=localhost dbname=gestiontarjetas sslmode=disable")
	if err != nil {
		log.Fatal(err)
	} else {
		_, err := db.Exec(`
			select gen_resumen(30878666,'202111');
			select gen_resumen(94538755,'202111');
			select gen_resumen(17305474,'202111');
		`)
		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Printf("\n------------- RESUMENES GENERADAS -------------\n")
		}

	}
}
