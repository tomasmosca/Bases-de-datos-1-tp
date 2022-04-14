create or replace function autorizar_compra(nrotarjeta_ char, codseguridad_ char, nrocomercio_ int, monto_ decimal) returns boolean as $$
declare
    datos_tarjeta record;
	ultimo_rechazo int;	
begin
	select * into datos_tarjeta from tarjeta t1 where nrotarjeta_ = t1.nrotarjeta;
	if not found or datos_tarjeta.estado = 'anulada' then
		select coalesce(max(r1.nrorechazo),0) into ultimo_rechazo from rechazo r1;
		insert into rechazo values(ultimo_rechazo+1, nrotarjeta_, nrocomercio_, current_date, monto_, 'tarjeta no válida ó no vigente');
		return false;
	elsif datos_tarjeta.estado = 'suspendida' then
		select coalesce(max(r1.nrorechazo),0) into ultimo_rechazo from rechazo r1;
		insert into rechazo values(ultimo_rechazo+1, nrotarjeta_, nrocomercio_, current_date, monto_, 'la tarjeta se encuentra suspendida');
		return false;
	elsif datos_tarjeta.codseguridad != codseguridad_ then
		select coalesce(max(r1.nrorechazo),0) into ultimo_rechazo from rechazo r1;
		insert into rechazo values(ultimo_rechazo+1, nrotarjeta_, nrocomercio_, current_date, monto_, 'código de seguridad inválido');
		return false;
	elsif ((select sum(monto) from compra c1 where c1.nrotarjeta = nrotarjeta_ and c1.pagado is false)+ monto_) > datos_tarjeta.limitecompra then
		select coalesce(max(r1.nrorechazo),0) into ultimo_rechazo from rechazo r1;
		insert into rechazo values(ultimo_rechazo+1, nrotarjeta_, nrocomercio_, current_date, monto_, 'supera límite de tarjeta');
		return false;
	elsif to_date(datos_tarjeta.validahasta,'yyyymm') < current_date then
		select coalesce(max(r1.nrorechazo),0) into ultimo_rechazo from rechazo r1;
		insert into rechazo values(ultimo_rechazo+1, nrotarjeta_, nrocomercio_, current_date, monto_, 'plazo de vigencia expirado');
		return false;	
	end if;
	return true;
		
end;
$$ language plpgsql;

CREATE OR REPLACE FUNCTION alerta_rechazo()
  RETURNS trigger AS
$$
DECLARE
	ultima_alerta int;
	rechazo_anterior record;
BEGIN
	select * into rechazo_anterior from rechazo r1 where DATE_TRUNC('day',r1.fecha) = DATE_TRUNC('day', NEW.fecha) and r1.nrotarjeta = NEW.nrotarjeta and r1.nrorechazo != NEW.nrorechazo;
	if found and rechazo_anterior.motivo = 'supera límite de tarjeta' and NEW.motivo = 'supera límite de tarjeta' then 
		UPDATE tarjeta SET estado = 'suspendida' WHERE NEW.nrotarjeta;
		select coalesce(max(a1.nroalerta),0) into ultima_alerta from alerta a1;
		INSERT INTO alerta values(ultima_alerta+1, NEW.nrotarjeta, NEW.fecha, NEW.nrorechazo, 32, 'bloqueo de tarjeta');	
	    RETURN NEW;
	else
		select coalesce(max(a1.nroalerta),0) into ultima_alerta from alerta a1;
		INSERT INTO alerta values(ultima_alerta+1, NEW.nrotarjeta, NEW.fecha, NEW.nrorechazo,0, 'rechazo');	
	    RETURN NEW;
	end if;
END;
$$
LANGUAGE 'plpgsql';

CREATE TRIGGER alerta_trigger_rechazo
  AFTER INSERT
  ON rechazo
  FOR EACH ROW
  EXECUTE PROCEDURE alerta_rechazo();

CREATE OR REPLACE FUNCTION alerta_compra()
RETURNS trigger AS
$$
DECLARE
	ultima_alerta int;
	ultimo_rechazo int;
	compra_anterior record;
BEGIN
	select * into compra_anterior from compra c1 where c1.nrotarjeta = NEW.nrotarjeta and NEW.nrocomercio != c1.nrocomercio  order by c1.fecha desc limit 1; 
	if found then
		if EXTRACT(MINUTE from (NEW.fecha - compra_anterior.fecha)) = 0 and EXTRACT(SECOND from (NEW.fecha - compra_anterior.fecha)) < 60  and 
		    (select codigopostal from comercio where nrocomercio = compra_anterior.nrocomercio) = (select codigopostal from comercio where nrocomercio = NEW.nrocomercio) then
		    select coalesce(max(r1.nrorechazo),0) into ultimo_rechazo from rechazo r1;
			insert into rechazo values(ultimo_rechazo+1, NEW.nrotarjeta, NEW.nrocomercio, current_date, NEW.monto, 'supera límite de compras en un minuto');
			select coalesce(max(a1.nroalerta),0) into ultima_alerta from alerta a1; 
			INSERT INTO alerta values(ultima_alerta+1, NEW.nrotarjeta, NEW.fecha, ultimo_rechazo+1, 1, 'compra lapso menor a 1min');	
			RETURN NEW;
		elsif EXTRACT(MINUTE from (NEW.fecha - compra_anterior.fecha)) < 5 and
		(select codigopostal from comercio where nrocomercio = compra_anterior.nrocomercio) != (select codigopostal from comercio where nrocomercio = NEW.nrocomercio) then
		    select coalesce(max(r1.nrorechazo),0) into ultimo_rechazo from rechazo r1;
			insert into rechazo values(ultimo_rechazo+1, NEW.nrotarjeta, NEW.nrocomercio, current_date, NEW.monto, 'supera límite de compras en 5 minuto');
			select coalesce(max(a1.nroalerta),0) into ultima_alerta from alerta a1; 
			INSERT INTO alerta values(ultima_alerta+1, NEW.nrotarjeta, NEW.fecha, ultimo_rechazo+1, 5, 'compra lapso menor a 5min');	
			RETURN NEW;
		end if;
	end if;
END;
$$
LANGUAGE 'plpgsql';

CREATE TRIGGER alerta_trigger_compra
  BEFORE INSERT
  ON compra
  FOR EACH ROW
  EXECUTE PROCEDURE alerta_compra();

create or replace function gen_resumen(nrocliente int, periodo char) returns boolean as $$
declare
	v record;
	compras_cliente record;
	periodo_cierre record;
	ultimo_cabecera int;
	ultimo_detalle int;
	monto_total decimal(7,2);
begin
	monto_total := 0.00;

	select coalesce(max(nroresumen),0) into ultimo_cabecera from cabecera;
	ultimo_cabecera := ultimo_cabecera + 1;
	
	for v in select *,com.nombre as nombrecomercio  from cliente c, tarjeta t, compra co, comercio com  where c.nrocliente = t.nrocliente and t.nrotarjeta = co.nrotarjeta and co.nrocomercio = com.nrocomercio loop

		select * into periodo_cierre from cierre c where c.terminacion = cast(substring(v.nrotarjeta,16,1) as int) and cast(substring(periodo,1,4) as int) = c.año and cast(substring(periodo,5,2) as int) = c.mes;
		if v.fecha >= periodo_cierre.fechainicio and v.fecha <= periodo_cierre.fechacierre then
			select coalesce(max(nrolinea),0) into ultimo_detalle from detalle;
			insert into detalle values (ultimo_cabecera,ultimo_detalle+1,v.fecha,v.nombrecomercio,v.monto);
			monto_total := monto_total + v.monto;
		end if;

	end loop;

	insert into cabecera values (ultimo_cabecera,v.nombre,v.apellido,v.domicilio,v.nrotarjeta,periodo_cierre.fechainicio,periodo_cierre.fechacierre,periodo_cierre.fechavto,monto_total);	

	return true;

end;

$$ language plpgsql;
