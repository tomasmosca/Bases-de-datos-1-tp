
alter table cliente   add constraint cliente_pk    primary key (nrocliente);
alter table tarjeta   add constraint tarjeta_pk    primary key (nrotarjeta);
alter table comercio  add constraint comercio_pk    primary key (nrocomercio);
alter table detalle   add constraint detalle_pk    primary key (nroresumen, nrolinea);
alter table cierre   add constraint cierre_pk      primary key (año, mes, terminacion);
alter table rechazo add constraint rechazo_pk      primary key (nrorechazo);
alter table alerta add constraint alerta_pk      primary key (nroalerta);
alter table compra add constraint compra_pk      primary key (nrooperacion);
alter table cabecera add constraint cabecera_pk      primary key (nroresumen);

alter table tarjeta add constraint tarjeta_nrocliente_fk foreign key (nrocliente) references cliente(nrocliente);
alter table compra add constraint compra_nrocomercio_fk foreign key (nrocomercio) references comercio(nrocomercio);
alter table compra add constraint compra_nrotarjeta_fk foreign key (nrotarjeta) references tarjeta(nrotarjeta);
alter table rechazo add constraint rechazo_nrocomercio_fk foreign key (nrocomercio) references comercio(nrocomercio);
alter table rechazo add constraint rechazo_nrotarjeta_fk foreign key (nrotarjeta) references tarjeta(nrotarjeta);
alter table cabecera add constraint cabecera_nrotarjeta_fk foreign key (nrotarjeta) references tarjeta(nrotarjeta);
alter table detalle add constraint detalle_nroresumen_fk foreign key (nroresumen) references cabecera(nroresumen);
alter table alerta add constraint alerta_nrotarjeta_fk foreign key (nrotarjeta) references tarjeta(nrotarjeta);
alter table alerta add constraint alerta_nrorechazo_fk foreign key (nrorechazo) references rechazo(nrorechazo);

