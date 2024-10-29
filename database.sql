/*==============================================================*/
/* DBMS name:      ORACLE Version 19c                           */
/* Created on:     22.10.2024 16:52:38                          */
/*==============================================================*/


alter table FLIGHT
   drop constraint FK_FLIGHT_FROM_AIRPORT;

alter table FLIGHT
   drop constraint FK_FLIGHT_PLANE;

alter table FLIGHT
   drop constraint FK_FLIGHT_TO_AIRPORT;

alter table FLIGHT
   drop constraint FK_FLIGHT_USER;

alter table TICKET
   drop constraint FK_TICKET_FLIGHT;

alter table TICKET
   drop constraint FK_TICKET_USER;

alter table USERROLE
   drop constraint FK_USERROLE_ROLE;

alter table USERROLE
   drop constraint FK_USERROLE_USER;

drop table AIRPORT cascade constraints;

drop table FLIGHT cascade constraints;

drop table PLANE cascade constraints;

drop table ROLE cascade constraints;

drop table TICKET cascade constraints;

drop table "USER" cascade constraints;

drop table USERROLE cascade constraints;

/*==============================================================*/
/* Table: AIRPORT                                               */
/*==============================================================*/
create table AIRPORT (
   ID                   LONG                  not null,
   NAME                 VARCHAR(255),
   COUNTRY              VARCHAR(255)          not null,
   CITY                 VARCHAR(255)          not null,
   constraint PK_AIRPORT primary key (ID)
);

/*==============================================================*/
/* Table: FLIGHT                                                */
/*==============================================================*/
create table FLIGHT (
   ID                   LONG                  not null,
   "FROM"               LONG                  not null,
   "TO"                 LONG                  not null,
   "DATE"               DATE                  not null,
   PILOT                LONG                  not null,
   PLANE                LONG                  not null,
   constraint PK_FLIGHT primary key (ID)
);

/*==============================================================*/
/* Table: PLANE                                                 */
/*==============================================================*/
create table PLANE (
   ID                   LONG                  not null,
   NAME                 VARCHAR(255),
   MODEL                VARCHAR(255)          not null,
   SEATS                NUMBER(10)            not null,
   constraint PK_PLANE primary key (ID)
);

/*==============================================================*/
/* Table: ROLE                                                  */
/*==============================================================*/
create table ROLE (
   ID                   LONG                  not null,
   NAME                 VARCHAR(255)          not null,
   LABEL                VARCHAR(255)          not null,
   constraint PK_ROLE primary key (ID)
);

/*==============================================================*/
/* Table: TICKET                                                */
/*==============================================================*/
create table TICKET (
   ID                   LONG                  not null,
   "USER"               LONG                  not null,
   FLIGHT               LONG                  not null,
   constraint PK_TICKET primary key (ID)
);

/*==============================================================*/
/* Table: "USER"                                                */
/*==============================================================*/
create table "USER" (
   ID                   LONG                  not null,
   FIRSTNAME            VARCHAR(255)          not null,
   LASTNAME             VARCHAR(255)          not null,
   BIRTHDATE            DATE                  not null,
   PASSWORD             VARCHAR(255)          not null,
   ACTIVE               BINARY(1)             not null,
   constraint PK_USER primary key (ID)
);

/*==============================================================*/
/* Table: USERROLE                                              */
/*==============================================================*/
create table USERROLE (
   ID                   LONG                  not null,
   "USER"               LONG                  not null,
   ROLE                 LONG                  not null,
   constraint PK_USERROLE primary key (ID)
);

alter table FLIGHT
   add constraint FK_FLIGHT_FROM_AIRPORT foreign key ("FROM")
      references AIRPORT (ID);

alter table FLIGHT
   add constraint FK_FLIGHT_PLANE foreign key (PLANE)
      references PLANE (ID);

alter table FLIGHT
   add constraint FK_FLIGHT_TO_AIRPORT foreign key ("TO")
      references AIRPORT (ID);

alter table FLIGHT
   add constraint FK_FLIGHT_USER foreign key (PILOT)
      references "USER" (ID);

alter table TICKET
   add constraint FK_TICKET_FLIGHT foreign key (FLIGHT)
      references FLIGHT (ID);

alter table TICKET
   add constraint FK_TICKET_USER foreign key ("USER")
      references "USER" (ID);

alter table USERROLE
   add constraint FK_USERROLE_ROLE foreign key (ROLE)
      references ROLE (ID);

alter table USERROLE
   add constraint FK_USERROLE_USER foreign key ("USER")
      references "USER" (ID);

