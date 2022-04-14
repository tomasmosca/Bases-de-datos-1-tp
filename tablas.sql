drop database if exists gestiontarjetas;
create database gestiontarjetas;

\c gestiontarjetas

create table cliente(
	nrocliente    int,
	nombre        text,
	apellido      text,
	domicilio	  text,
	telefono      char(12)
);

create table tarjeta(
	nrotarjeta    char(16),
	nrocliente    int,
	validadesde   char(6),
	validahasta   char(6),
	codseguridad  char(4),
	limitecompra  decimal(8,2),
	estado        char(10)
);

create table comercio(
	nrocomercio   int,
	nombre        text,
	domicilio     text,
	codigopostal  char(8),
	telefono      char(12)
);

create table detalle(
	nroresumen	int,
	nrolinea	int,
	fecha		date,
	nombrecomercio	text,
	monto		decimal(7,2)
);

create table cierre(
	año		int,
	mes		int,
	terminacion	int,
	fechainicio	date,
	fechacierre	date,
	fechavto	date
);

create table compra(
	nrooperacion	int,
	nrotarjeta 		char(16),
	nrocomercio 	int,
	fecha 			timestamp,
	monto 			decimal(7,2),
	pagado 			boolean
);

create table cabecera(
	nroresumen		int,
	nombre 			text,
	apellido 		text,
	domicilio 		text,
	nrotarjeta 		char(16),
	desde 			date,
	hasta 			date,
	vence 			date,
	total 			decimal(8,2)
);

create table consumo(
	nrotarjeta 		char(16),
	codseguridad    char(4),
	nrocomercio     int,
	monto           decimal(7,2)
);

create table rechazo(
	nrorechazo int,
	nrotarjeta char(16),
	nrocomercio int,
	fecha timestamp,
	monto decimal(7,2),
	motivo text
);

create table alerta(
	nroalerta int,
	nrotarjeta char(16),
	fecha timestamp,
	nrorechazo int,
	codalerta int, 
	descripcion text
);
