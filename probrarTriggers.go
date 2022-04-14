package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

func ProbarTriggers() {
	db, err := sql.Open("postgres", "user=postgres host=localhost dbname=gestiontarjetas sslmode=disable")
	if err != nil {
		log.Fatal(err)
	} else {
		_, err := db.Exec(`insert into rechazo values(150, '5449981007097362',1,current_timestamp,100,'');
						   insert into compra values(120, '5449981007097362',1,current_timestamp,100,true);
						   insert into compra values(121, '5449981007097362',2,current_timestamp,100,true);						   
						   insert into compra values(122, '5210983186476711',3,current_timestamp,100,true);
						   insert into compra values(123, '5210983186476711',4,current_timestamp,100,true);
						   insert into rechazo values(160, '5299857355514060',1,current_timestamp,100,'supera limite de tarjeta');
						   insert into rechazo values(161, '5299857355514060',1,current_timestamp,100,'supera limite de tarjeta');
		`)
		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Printf("\n------------- SP PROBADOS -------------\n")
		}

	}
}
