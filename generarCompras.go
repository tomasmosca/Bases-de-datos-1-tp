package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

func GenerarCompras() {
	db, err := sql.Open("postgres", "user=postgres host=localhost dbname=gestiontarjetas sslmode=disable")
	if err != nil {
		log.Fatal(err)
	} else {
		_, err := db.Exec(`
		create or replace function generarCompras() returns void as $$ 
		declare
			c record;
		begin
			for c in select * from consumo
				loop
					if (select autorizar_compra(c.nrotarjeta,c.codseguridad,c.nrocomercio,c.monto)) = true then
						insert into compra values((select coalesce(max(nrooperacion),0)+1 from compra),c.nrotarjeta,c.nrocomercio,current_timestamp,c.monto,false);
					end if;	
				end loop;						
				return;
		end;
		$$language plpgsql;

		select generarCompras();
		`)
		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Printf("\n------------- COMPRAS GENERADAS -------------\n")
		}

	}
}
