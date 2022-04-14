package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

func CargarTablas() {
	db, err := sql.Open("postgres", "user=postgres host=localhost dbname=gestiontarjetas sslmode=disable")
	if err != nil {
		log.Fatal(err)
	} else {
		_, err := db.Exec(`insert into comercio values(1,'Maxikiosco 365','Int.Becco 458','B1611FDA','11-4741-7580');
						   insert into comercio values(2,'Farmacia 24hs','Av.Angel T. de Alvear 3250','B1611FDA','11-4723-9250');
						   insert into comercio values(3,'Supermercado Lo de Claudio','Formosa 1967','B1619FDB','11-4845-6594');
						   insert into comercio values(4,'El almacen de doña rosa','Acasusso 864','B1622FDA','11-3865-9941');
						   insert into comercio values(5,'Ferreteria Gonzales','Estrada 320','B1611FDB','11-4078-4554');
						   insert into comercio values(6,'Respuestos DIPA','Capilla del Señor 834','B1751FDA','11-4268-4600');
						   insert into comercio values(7,'Suspension Ruben','Av. Eva Peron 1667','B1610FDB','11-4461-1194');
						   insert into comercio values(8,'Montori e Hijos','Rivadavia 1156','B1827FDB','11-3221-4512');
						   insert into comercio values(9,'Neumaticos Fernandez','Av. 9 de Julio 800','B1075FDA','11-5441-0221');
						   insert into comercio values(10,'Joyas Laprida','Laprida 211','B1832FDA','11-5462-3497');
						   insert into comercio values(11,'Papeleria La Feria','Av. Corrientes 834','B1681FDB','11-6262-2747');
						   insert into comercio values(12,'Maxikiosco Martina','Las Violetas 2140','B1748FDA','11-1582-5847');
						   insert into comercio values(13,'Farmacia Crl','Boulevar Caveri 1334','B1736IIO','11-2002-3053');
						   insert into comercio values(14,'Supermercado La Feria','Av Entre Ríos 1072','B1121CEA','11-9212-2577');
						   insert into comercio values(15,'Ferreteria La Feria','Esteban de Luca 2235','B1181FDB','11-1362-2817');
						   insert into comercio values(16,'Pintureria Ambar','Av. Constituyentes 1204','B1617CDE','11-4740-6048');
						   insert into comercio values(17,'Panaderia El Zorzal','Av. Constituyentes 1040','B1691FDB','11-1761-2897');
						   insert into comercio values(18,'Ferreteria Liniers','Av. Liniers 1846','B1648DBL','11-4731-5930');
						   insert into comercio values(19,'Kiosco Candy 24','Av. Sta. María 4711','B1812DCL','11-1821-6282');
						   insert into comercio values(20,'Librería Arturo','Emilio Mitre 351','B1661TDB','11-3827-7217');

						   insert into cliente values(30878666,'Sol','Palermo','Av E Shaw 121','11-6543-5234');
						   insert into cliente values(94538755,'Fonda','Rios','Brasil Oeste 4113','11-7632-7754');
						   insert into cliente values(86734897,'Grazia','Alvarez','Av Rivadavia 7266','11-8964-2314');
						   insert into cliente values(34494968,'Patricia','Perez','Sastre Marcos 33123','11-7696-9980');
						   insert into cliente values(32873620,'Angelino','Martin','San Martin 1996','11-3411-9778');
						   insert into cliente values(37606436,'Florencia','Medina','Chacabuco 1267','11-7590-4221');
						   insert into cliente values(38099826,'Matias','Perez','Av. Mitre 412','11-5543-0098');
						   insert into cliente values(40425321,'Lucas','Armando','Jose M. Gutierrez 1541','11-3533-6766');
						   insert into cliente values(25501234,'Maria','Caceres','Callao 1522','11-4521-2421');
						   insert into cliente values(20341579,'Natalia','Colman','Mario Bravo 350','11-3855-7764');
						   insert into cliente values(32372285,'Dionisa','Bravo','Pte B Mitre 385','11-3127-1121');
						   insert into cliente values(51301792,'Rosario','Pellegrini','Hipolito Irigoyen 3476','11-9911-8676');
						   insert into cliente values(95249665,'Ella','Torres','Riobamba 7110','11-5521-2231');
						   insert into cliente values(21600249,'Adrían','Sánchez','Santa Fe Avda 16102','11-4300-7882');
						   insert into cliente values(86497599,'Micaela','Cruz','Jujuy 1378','11-2932-1010');
						   insert into cliente values(42226235,'Tomas','Arce','Neuquen 498','11-4993-3372');
						   insert into cliente values(42013792,'Nicolas','Arias','Av. Rivadavia 7652','11-2895-4664');
						   insert into cliente values(41768622,'Carla','Cabrera','Nicolas avellaneda 225','11-6537-7397');
						   insert into cliente values(20317620,'Juan','Vaccaro','Mariano Pelliza 7895','11-5719-9847');
						   insert into cliente values(17305474,'Carlos','Durante','Obarrio 122','11-8559-4879');

						   insert into cierre values(2021,01,0,'2020-12-17','2021-01-16','2021-02-05');
						   insert into cierre values(2021,01,1,'2020-12-18','2021-01-17','2021-02-06');
						   insert into cierre values(2021,01,2,'2020-12-19','2021-01-18','2021-02-07');
						   insert into cierre values(2021,01,3,'2020-12-20','2021-01-19','2021-02-08');
						   insert into cierre values(2021,01,4,'2020-12-21','2021-01-20','2021-02-09');
						   insert into cierre values(2021,01,5,'2020-12-22','2021-01-21','2021-02-10');
						   insert into cierre values(2021,01,6,'2020-12-23','2021-01-22','2021-02-11');
						   insert into cierre values(2021,01,7,'2020-12-24','2021-01-23','2021-02-12');
						   insert into cierre values(2021,01,8,'2020-12-25','2021-01-24','2021-02-13');
						   insert into cierre values(2021,01,9,'2020-12-26','2021-01-25','2021-02-14');
						   insert into cierre values(2021,02,0,'2021-01-17','2021-02-16','2021-03-05');
						   insert into cierre values(2021,02,1,'2021-01-18','2021-02-17','2021-03-06');
						   insert into cierre values(2021,02,2,'2021-01-19','2021-02-18','2021-03-07');
						   insert into cierre values(2021,02,3,'2021-01-20','2021-02-19','2021-03-08');
						   insert into cierre values(2021,02,4,'2021-01-21','2021-02-20','2021-03-09');
						   insert into cierre values(2021,02,5,'2021-01-22','2021-02-21','2021-03-10');
						   insert into cierre values(2021,02,6,'2021-01-23','2021-02-22','2021-03-11');
						   insert into cierre values(2021,02,7,'2021-01-24','2021-02-23','2021-03-12');
						   insert into cierre values(2021,02,8,'2021-01-25','2021-02-24','2021-03-13');
						   insert into cierre values(2021,02,9,'2021-01-26','2021-02-25','2021-03-14');
						   insert into cierre values(2021,03,0,'2021-02-17','2021-03-16','2021-04-05');
						   insert into cierre values(2021,03,1,'2021-02-18','2021-03-17','2021-04-06');
						   insert into cierre values(2021,03,2,'2021-02-19','2021-03-18','2021-04-07');
						   insert into cierre values(2021,03,3,'2021-02-20','2021-03-19','2021-04-08');
						   insert into cierre values(2021,03,4,'2021-02-21','2021-03-20','2021-04-09');
						   insert into cierre values(2021,03,5,'2021-02-22','2021-03-21','2021-04-10');
						   insert into cierre values(2021,03,6,'2021-02-23','2021-03-22','2021-04-11');
						   insert into cierre values(2021,03,7,'2021-02-24','2021-03-23','2021-04-12');
						   insert into cierre values(2021,03,8,'2021-02-25','2021-03-24','2021-04-13');
						   insert into cierre values(2021,03,9,'2021-02-26','2021-03-25','2021-04-14');
						   insert into cierre values(2021,04,0,'2021-03-17','2021-04-16','2021-05-05');
						   insert into cierre values(2021,04,1,'2021-03-18','2021-04-17','2021-05-06');
						   insert into cierre values(2021,04,2,'2021-03-19','2021-04-18','2021-05-07');
						   insert into cierre values(2021,04,3,'2021-03-20','2021-04-19','2021-05-08');
						   insert into cierre values(2021,04,4,'2021-03-21','2021-04-20','2021-05-09');
						   insert into cierre values(2021,04,5,'2021-03-22','2021-04-21','2021-05-10');
						   insert into cierre values(2021,04,6,'2021-03-23','2021-04-22','2021-05-11');
						   insert into cierre values(2021,04,7,'2021-03-24','2021-04-23','2021-05-12');
						   insert into cierre values(2021,04,8,'2021-03-25','2021-04-24','2021-05-13');
						   insert into cierre values(2021,04,9,'2021-03-26','2021-04-25','2021-05-14');
						   insert into cierre values(2021,05,0,'2021-04-17','2021-05-16','2021-06-05');
						   insert into cierre values(2021,05,1,'2021-04-18','2021-05-17','2021-06-06');
						   insert into cierre values(2021,05,2,'2021-04-19','2021-05-18','2021-06-07');
						   insert into cierre values(2021,05,3,'2021-04-20','2021-05-19','2021-06-08');
						   insert into cierre values(2021,05,4,'2021-04-21','2021-05-20','2021-06-09');
						   insert into cierre values(2021,05,5,'2021-04-22','2021-05-21','2021-06-10');
						   insert into cierre values(2021,05,6,'2021-04-23','2021-05-22','2021-06-11');
						   insert into cierre values(2021,05,7,'2021-04-24','2021-05-23','2021-06-12');
						   insert into cierre values(2021,05,8,'2021-04-25','2021-05-24','2021-06-13');
						   insert into cierre values(2021,05,9,'2021-04-26','2021-05-25','2021-06-14');
						   insert into cierre values(2021,06,0,'2021-05-17','2021-06-16','2021-07-05');
						   insert into cierre values(2021,06,1,'2021-05-18','2021-06-17','2021-07-06');
						   insert into cierre values(2021,06,2,'2021-05-19','2021-06-18','2021-07-07');
						   insert into cierre values(2021,06,3,'2021-05-20','2021-06-19','2021-07-08');
						   insert into cierre values(2021,06,4,'2021-05-21','2021-06-20','2021-07-09');
						   insert into cierre values(2021,06,5,'2021-05-22','2021-06-21','2021-07-10');
						   insert into cierre values(2021,06,6,'2021-05-23','2021-06-22','2021-07-11');
						   insert into cierre values(2021,06,7,'2021-05-24','2021-06-23','2021-07-12');
						   insert into cierre values(2021,06,8,'2021-05-25','2021-06-24','2021-07-13');
						   insert into cierre values(2021,06,9,'2021-05-26','2021-06-25','2021-07-14');
						   insert into cierre values(2021,07,0,'2021-06-17','2021-07-16','2021-08-05');
						   insert into cierre values(2021,07,1,'2021-06-18','2021-07-17','2021-08-06');
						   insert into cierre values(2021,07,2,'2021-06-19','2021-07-18','2021-08-07');
						   insert into cierre values(2021,07,3,'2021-06-20','2021-07-19','2021-08-08');
						   insert into cierre values(2021,07,4,'2021-06-21','2021-07-20','2021-08-09');
						   insert into cierre values(2021,07,5,'2021-06-22','2021-07-21','2021-08-10');
						   insert into cierre values(2021,07,6,'2021-06-23','2021-07-22','2021-08-11');
						   insert into cierre values(2021,07,7,'2021-06-24','2021-07-23','2021-08-12');
						   insert into cierre values(2021,07,8,'2021-06-25','2021-07-24','2021-08-13');
						   insert into cierre values(2021,07,9,'2021-06-26','2021-07-25','2021-08-14');
						   insert into cierre values(2021,08,0,'2021-07-17','2021-08-16','2021-09-05');
						   insert into cierre values(2021,08,1,'2021-07-18','2021-08-17','2021-09-06');
						   insert into cierre values(2021,08,2,'2021-07-19','2021-08-18','2021-09-07');
						   insert into cierre values(2021,08,3,'2021-07-20','2021-08-19','2021-09-08');
						   insert into cierre values(2021,08,4,'2021-07-21','2021-08-20','2021-09-09');
						   insert into cierre values(2021,08,5,'2021-07-22','2021-08-21','2021-09-10');
						   insert into cierre values(2021,08,6,'2021-07-23','2021-08-22','2021-09-11');
						   insert into cierre values(2021,08,7,'2021-07-24','2021-08-23','2021-09-12');
						   insert into cierre values(2021,08,8,'2021-07-25','2021-08-24','2021-09-13');
						   insert into cierre values(2021,08,9,'2021-07-26','2021-08-25','2021-09-14');
						   insert into cierre values(2021,09,0,'2021-08-17','2021-09-16','2021-10-05');
						   insert into cierre values(2021,09,1,'2021-08-18','2021-09-17','2021-10-06');
						   insert into cierre values(2021,09,2,'2021-08-19','2021-09-18','2021-10-07');
						   insert into cierre values(2021,09,3,'2021-08-20','2021-09-19','2021-10-08');
						   insert into cierre values(2021,09,4,'2021-08-21','2021-09-20','2021-10-09');
						   insert into cierre values(2021,09,5,'2021-08-22','2021-09-21','2021-10-10');
						   insert into cierre values(2021,09,6,'2021-08-23','2021-09-22','2021-10-11');
						   insert into cierre values(2021,09,7,'2021-08-24','2021-09-23','2021-10-12');
						   insert into cierre values(2021,09,8,'2021-08-25','2021-09-24','2021-10-13');
						   insert into cierre values(2021,09,9,'2021-08-26','2021-09-25','2021-10-14');
						   insert into cierre values(2021,10,0,'2021-09-17','2021-10-16','2021-11-05');
						   insert into cierre values(2021,10,1,'2021-09-18','2021-10-17','2021-11-06');
						   insert into cierre values(2021,10,2,'2021-09-19','2021-10-18','2021-11-07');
						   insert into cierre values(2021,10,3,'2021-09-20','2021-10-19','2021-11-08');
						   insert into cierre values(2021,10,4,'2021-09-21','2021-10-20','2021-11-09');
						   insert into cierre values(2021,10,5,'2021-09-22','2021-10-21','2021-11-10');
						   insert into cierre values(2021,10,6,'2021-09-23','2021-10-22','2021-11-11');
						   insert into cierre values(2021,10,7,'2021-09-24','2021-10-23','2021-11-12');
						   insert into cierre values(2021,10,8,'2021-09-25','2021-10-24','2021-11-13');
						   insert into cierre values(2021,10,9,'2021-09-26','2021-10-25','2021-11-14');
						   insert into cierre values(2021,11,0,'2021-10-17','2021-11-16','2021-12-05');
						   insert into cierre values(2021,11,1,'2021-10-18','2021-11-17','2021-12-06');
						   insert into cierre values(2021,11,2,'2021-10-19','2021-11-18','2021-12-07');
						   insert into cierre values(2021,11,3,'2021-10-20','2021-11-19','2021-12-08');
						   insert into cierre values(2021,11,4,'2021-10-21','2021-11-20','2021-12-09');
						   insert into cierre values(2021,11,5,'2021-10-22','2021-11-21','2021-12-10');
						   insert into cierre values(2021,11,6,'2021-10-23','2021-11-22','2021-12-11');
						   insert into cierre values(2021,11,7,'2021-10-24','2021-11-23','2021-12-12');
						   insert into cierre values(2021,11,8,'2021-10-25','2021-11-24','2021-12-13');
						   insert into cierre values(2021,11,9,'2021-10-26','2021-11-25','2021-12-14');
						   insert into cierre values(2021,12,0,'2021-11-17','2021-12-16','2022-01-05');
						   insert into cierre values(2021,12,1,'2021-11-18','2021-12-17','2022-01-06');
						   insert into cierre values(2021,12,2,'2021-11-19','2021-12-18','2022-01-07');
						   insert into cierre values(2021,12,3,'2021-11-20','2021-12-19','2022-01-08');
						   insert into cierre values(2021,12,4,'2021-11-21','2021-12-20','2022-01-09');
						   insert into cierre values(2021,12,5,'2021-11-22','2021-12-21','2022-01-10');
						   insert into cierre values(2021,12,6,'2021-11-23','2021-12-22','2022-01-11');
						   insert into cierre values(2021,12,7,'2021-11-24','2021-12-23','2022-01-12');
						   insert into cierre values(2021,12,8,'2021-11-25','2021-12-24','2022-01-13');
						   insert into cierre values(2021,12,9,'2021-11-26','2021-12-25','2022-01-14');

						   insert into tarjeta values('5449981007097362', 30878666, '201812','202312','0810',80000.00, 'vigente');
						   insert into tarjeta values('5215587392715740', 30878666, '201212','202112','0546',8000.00, 'anulada');
						   insert into tarjeta values('5210983186476711', 94538755, '202008','202508','0251',40000.00, 'vigente');
						   insert into tarjeta values('5245646358674806', 94538755, '201508','202408','0133',20000.00, 'anulada');
						   insert into tarjeta values('4927053520951527', 86734897, '202101','202701','0521',10000.00, 'suspendida');
						   insert into tarjeta values('4024813453580367', 34494968, '201710','202010','0251',30000.00, 'vigente');
						   insert into tarjeta values('5288428075974845', 32873620, '201811','201911','0117',35000.00, 'vigente');
						   insert into tarjeta values('4554207750976410', 37606436, '202006','202106','0178',90000.00, 'vigente');
						   insert into tarjeta values('5267469005021351', 38099826, '201910','202310','0378',50000.00, 'suspendida');
						   insert into tarjeta values('5299857355514060', 40425321, '201512','202112','0635',45000.00, 'vigente');
						   insert into tarjeta values('5289102401524975', 25501234, '201803','202203','0635',120000.00, 'vigente');
						   insert into tarjeta values('4495639786262423', 20341579, '202009','202409','0508',100000.00, 'vigente');
						   insert into tarjeta values('4422834927230484', 32372285, '202109','202509','0986',70000.00, 'suspendida');
						   insert into tarjeta values('4403651796135247', 51301792, '201802','202202','0708',75000.00, 'vigente');
						   insert into tarjeta values('3558880738242960', 95249665, '201512','202112','0217',95000.00, 'suspendida');
						   insert into tarjeta values('3538264985639241', 21600249, '201908','202408','0351',180000.00, 'vigente');
						   insert into tarjeta values('3529559193777084', 86497599, '201703','202203','0252',60000.00, 'anulada');
						   insert into tarjeta values('5577163024783704', 42226235, '201605','202205','0872',100000.00, 'vigente');
						   insert into tarjeta values('5201262513928574', 42013792, '201808','202308','0479',90000.00, 'suspendida');
						   insert into tarjeta values('4191297487611553', 41768622, '201910','202310','0479',85000.00, 'vigente');
						   insert into tarjeta values('4567240522998854', 20317620, '201907','202207','0756',120000.00, 'suspendida');
						   insert into tarjeta values('4658368747812582', 17305474, '201706','202406','0755',110000.00, 'vigente');						   				   
					   

						   insert into consumo values('5449981007097362','0810',2,50000);
						   insert into consumo values('5210983186476711','0251',1,3000);
						   insert into consumo values('4658368747812582','0755',5,11500);
						  
						   insert into consumo values('4658358747812582','0755',5,1800);
						   								   
						   insert into consumo values('3529559193777084','0252',17,2600);
					  							   
						   insert into consumo values('3558880738242960','0217',4,9650);
						  
						   insert into consumo values('4658368747812582','8755',5,11500);

						   insert into consumo values('4554207750976410','0178',9,9000);
						   
						   insert into consumo values('5449981007097362','0810',3,40000);
						   insert into consumo values('5449981007097362','0810',1,1000);
						   insert into consumo values('5210983186476711','0251',4,6000);
						   `)

		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Printf("\n------------- TABLAS CARGADAS -------------\n")
		}

	}
}
