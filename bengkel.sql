create database bengkel;
use bengkel;
select database();
create table services (
serviceID varchar(100) primary key,
serviceType varchar(100),
vehicleType varchar(100),
servicePrice int);
select * from services;

create table customers (
customerID varchar(100)primary key,
customerName varchar(100),
bookingDate date not null,
bookingTime time not null,
numberPlat Varchar(10),
serviceID varchar(100));
select * from customers;

create table progress(
progressID varchar(10) primary key,
serviceID varchar(10),
customerID varchar(10),
serviceStatus varchar(10));
select * from progress;


 insert into services(serviceID, serviceType, vehicleType, servicePrice) values
("SRV1", "Ganti Oli", "Mobil", 708000),
("SRV2", "Ganti Oli", "Mobil", 590000),
("SRV3", "Ganti Oli", "Motor", 120000),
("SRV4", "Tune up Komplit", "Mobil", 780000),
("SRV5", "Tune up", "Motor", 100000);

insert into customers(customerID, customerName, bookingDate, bookingTime, numberPlat, serviceID) values
("CUS1", "Ahmadin Winarko","2020-07-07", "10:30:11", "L1234AN", "SRV1"),
("CUS2", "Sugeng Harianto","2020-07-07", "11:00:11", "L1345MO", "SRV2"),
("CUS3", "Tiara Putri","2020-07-08", "10:00:00", "L3670SO", "SRV3"),
("CUS4", "Risya Amarta","2020-07-08", "12:00:11", "L3676MA", "SRV4"),
("CUS5", "Ajeng Lazuardi","2020-07-08","14:00:59", "L36790OP", "SRV5");

insert into progress(progressID, serviceID, customerID, serviceStatus) values
("PROG1","SRV1", "CUS1","Done"),
("PROG2","SRV2", "CUS2","Waiting"),
("PROG3","SRV3", "CUS3","Waiting"),
("PROG4","SRV4", "CUS4","Done"),
("PROG5","SRV5", "CUS5","Done");
