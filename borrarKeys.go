package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func BorrarKeys() {
	db, err := sql.Open("postgres", "user=postgres host=localhost dbname=gestiontarjetas sslmode=disable")
	_, err = db.Exec(`alter table tarjeta drop constraint tarjeta_nrocliente_fk;
						alter table compra drop constraint compra_nrocomercio_fk;
						alter table compra drop constraint compra_nrotarjeta_fk;
						alter table rechazo drop constraint rechazo_nrocomercio_fk;
						alter table cabecera drop constraint cabecera_nrotarjeta_fk;
						alter table detalle drop constraint detalle_nroresumen_fk;
						alter table alerta drop constraint alerta_nrorechazo_fk;
						alter table cliente   drop constraint cliente_pk;
						alter table tarjeta   drop constraint tarjeta_pk;
						alter table comercio  drop constraint comercio_pk;
						alter table detalle   drop constraint detalle_pk;
						alter table cierre   drop constraint cierre_pk;
						alter table rechazo drop constraint rechazo_pk;
						alter table alerta drop constraint alerta_pk;
						alter table compra drop constraint compra_pk;
						alter table cabecera drop constraint cabecera_pk;`)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("------------- SE BORRARON PKs y FKs -------------\n")
	}
}
