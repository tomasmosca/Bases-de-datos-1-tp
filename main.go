package main

import (
	"fmt"
)

func main() {
	menu()
}

func menu() {
	fmt.Print("\033[H\033[2J")  //limpia la consola y muestra las opciones en pantalla
	var opcionElegida int
	fmt.Printf("------------- ELIGA UNA OPCION ------------- \n\n")
	fmt.Printf("1- Crear base de datos \n")
	fmt.Printf("2- Crear tablas \n")
	fmt.Printf("3- Crear PKs y FKs \n")
	fmt.Printf("4- Cargar tablas \n")
	fmt.Printf("5- Crear Stored Procedures y Triggers \n")
	fmt.Printf("6- Generar compras a partir de tabla consumo\n")
	fmt.Printf("7- Generar resumenes\n")
	fmt.Printf("8- Crear BD NoSQL y cargar datos\n")
	fmt.Printf("9- Borrar PKs y FKs\n")
	fmt.Printf("0- Salir \n\n")
	fmt.Printf("Eliga una opcion: ")
	fmt.Scanf("%d", &opcionElegida)
	elegirOpcion(opcionElegida)
}

func elegirOpcion(opcionElegida int) { //se ejecuta una funcion dependiendo la opcion elegida
	fmt.Print("\033[H\033[2J")
	if opcionElegida == 1 {
		CrearBase()
	}
	if opcionElegida == 2 {
		CrearTablas()
	}
	if opcionElegida == 3 {
		CrearKeys()
	}
	if opcionElegida == 4 {
		CargarTablas()
	}
	if opcionElegida == 5 {
		CrearSpTriggers()
	}
	if opcionElegida == 6 {
		GenerarCompras()
	}
	if opcionElegida == 7 {
		GenerarResumenes()
	}
	if opcionElegida == 8 {
		CrearDB()
	}
	if opcionElegida == 9 {
		BorrarKeys()
	}

	if opcionElegida != 0 {
		fmt.Printf("\nPresione 0 para volver al menu \n")
		salir()
	}
}

func salir() {  //se queda esperando que se ingrese el numero 0 para volver al menu
	var opcionElegida int
	fmt.Scanf("%d", &opcionElegida)
	if opcionElegida == 0 {
		menu()
	} else {
		salir()
	}
}
