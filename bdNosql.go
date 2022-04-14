package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"time"

	bolt "go.etcd.io/bbolt"
)

type Cliente struct {
	Nrocliente int
	Nombre     string
	Apellido   string
	Domicilio  string
	Telefono   string
}

type Tarjeta struct {
	Nrotarjeta   string
	Nrocliente   int
	Validadesde  string
	Validahasta  string
	Codseguridad string
	Limitecompra int
	Estado       string
}

type Comercio struct {
	Nrocomercio  int
	Nombre       string
	Domicilio    string
	Codigopostal string
	Telefono     string
}

type Compra struct {
	Nrooperacion int
	Nrotarjeta   string
	Nrocomercio  int
	Fecha        string
	Monto        int
	Pagado       bool
}

func CreateUpdate(db *bolt.DB, bucketName string, key []byte, val []byte) error {
	// abre transacción de escritura
	tx, err := db.Begin(true)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	b, _ := tx.CreateBucketIfNotExists([]byte(bucketName))

	err = b.Put(key, val)
	if err != nil {
		return err
	}

	// cierra transacción
	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}


func ReadAll(db *bolt.DB, bucketname string) {
	fmt.Printf("---------------%s--------------------\n\n", bucketname)
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketname))

		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			fmt.Printf("key=%s, value=%s\n", k, v)
		}
		return nil
	})
}

func CrearDB() {
	db, err := bolt.Open("bdNosql.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	tiempo := time.Now()
	current_timestamp := tiempo.Format("01-02-2006 15:04:05")

	cliente := Cliente{30878666, "Juan", "Perez", "Av. Siempreviva 123", "11-6543-5234"}
	cliente2 := Cliente{20638426, "Sol", "Perez", "Av. Siempreviva 123", "11-6543-5234"}
	cliente3 := Cliente{10673266, "Marcos", "Perez", "Av. Siempreviva 123", "11-6543-5234"}
	dataCliente, err := json.Marshal(cliente)
	dataCliente2, err := json.Marshal(cliente2)
	dataCliente3, err := json.Marshal(cliente3)

	tarjeta := Tarjeta{"5449981007097362", 30878666, "201812", "202312", "0810", 80000.00, "vigente"}
	tarjeta2 := Tarjeta{"5215587392715740", 30878666, "201212", "202112", "0546", 8000.00, "anulada"}
	tarjeta3 := Tarjeta{"4927053520951527", 86734897, "202101", "202701", "0521", 10000.00, "suspendida"}
	dataTarjeta, err := json.Marshal(tarjeta)
	dataTarjeta2, err := json.Marshal(tarjeta2)
	dataTarjeta3, err := json.Marshal(tarjeta3)

	comercio := Comercio{1, "Maxikiosco 365", "Int.Becco 458", "B1611FDA", "11-4741-7580"}
	comercio2 := Comercio{2, "Farmacia 24hs", "Av.Angel T. de Alvear 3250", "B1611FDA", "11-4723-9250"}
	comercio3 := Comercio{3, "Supermercado Lo de Claudio", "Formosa 1967", "B1619FDB", "11-4845-6594"}
	dataComercio, err := json.Marshal(comercio)
	dataComercio2, err := json.Marshal(comercio2)
	dataComercio3, err := json.Marshal(comercio3)

	compra := Compra{1, "5449981007097362", 1, current_timestamp, 6000, false}
	compra2 := Compra{2, "5215587392715740", 2, current_timestamp, 10000, false}
	compra3 := Compra{3, "4927053520951527", 3, current_timestamp, 20, false}
	dataCompra, err := json.Marshal(compra)
	dataCompra2, err := json.Marshal(compra2)
	dataCompra3, err := json.Marshal(compra3)

	if err != nil {
		log.Fatal(err)
	}

	CreateUpdate(db, "Cliente", []byte(strconv.Itoa(cliente.Nrocliente)), dataCliente)
	CreateUpdate(db, "Cliente", []byte(strconv.Itoa(cliente2.Nrocliente)), dataCliente2)
	CreateUpdate(db, "Cliente", []byte(strconv.Itoa(cliente3.Nrocliente)), dataCliente3)
	CreateUpdate(db, "Tarjeta", []byte(tarjeta.Nrotarjeta), dataTarjeta)
	CreateUpdate(db, "Tarjeta", []byte(tarjeta2.Nrotarjeta), dataTarjeta2)
	CreateUpdate(db, "Tarjeta", []byte(tarjeta3.Nrotarjeta), dataTarjeta3)
	CreateUpdate(db, "Comercio", []byte(strconv.Itoa(comercio.Nrocomercio)), dataComercio)
	CreateUpdate(db, "Comercio", []byte(strconv.Itoa(comercio2.Nrocomercio)), dataComercio2)
	CreateUpdate(db, "Comercio", []byte(strconv.Itoa(comercio3.Nrocomercio)), dataComercio3)
	CreateUpdate(db, "Compra", []byte(strconv.Itoa(compra.Nrooperacion)), dataCompra)
	CreateUpdate(db, "Compra", []byte(strconv.Itoa(compra2.Nrooperacion)), dataCompra2)
	CreateUpdate(db, "Compra", []byte(strconv.Itoa(compra3.Nrooperacion)), dataCompra3)

	ReadAll(db, "Cliente")
	ReadAll(db, "Tarjeta")
	ReadAll(db, "Comercio")
	ReadAll(db, "Compra")

}
