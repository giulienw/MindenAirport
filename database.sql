/*==============================================================*/
/* DBMS name:      ORACLE Version 19c                           */
/* Created on:     29.10.2024 15:03:56                          */
/*==============================================================*/


alter table FLIGHT
   drop constraint FK_FLIGHT_FROM_AIRPORT;

alter table FLIGHT
   drop constraint FK_FLIGHT_PLANE;

alter table FLIGHT
   drop constraint FK_FLIGHT_TO_AIRPORT;

alter table FLIGHT
   drop constraint FK_FLIGHT_USER;

alter table FLIGHT
   drop constraint FK_FLIGHT_REFERENCE_TERMINAL;

alter table PLANE
   drop constraint FK_PLANE_AIRLINE;

alter table TICKET
   drop constraint FK_TICKET_FLIGHT;

alter table TICKET
   drop constraint FK_TICKET_USER;

alter table USERROLE
   drop constraint FK_USERROLE_ROLE;

alter table USERROLE
   drop constraint FK_USERROLE_USER;

drop table AIRLINE cascade constraints;

drop table AIRPORT cascade constraints;

drop table FLIGHT cascade constraints;

drop table PLANE cascade constraints;

drop table ROLE cascade constraints;

drop table TERMINAL cascade constraints;

drop table TICKET cascade constraints;

drop table "USER" cascade constraints;

drop table USERROLE cascade constraints;

/*==============================================================*/
/* Table: AIRLINE                                               */
/*==============================================================*/
create table AIRLINE (
   ID                   VARCHAR2(36)          not null,
   NAME                 VARCHAR(255)          not null,
   constraint PK_AIRLINE primary key (ID)
);

/*==============================================================*/
/* Table: AIRPORT                                               */
/*==============================================================*/
create table AIRPORT (
   ID                   VARCHAR2(36)          not null,
   NAME                 VARCHAR(255),
   COUNTRY              VARCHAR(255)          not null,
   CITY                 VARCHAR(255)          not null,
   constraint PK_AIRPORT primary key (ID)
);

/*==============================================================*/
/* Table: FLIGHT                                                */
/*==============================================================*/
create table FLIGHT (
   ID                   VARCHAR2(36)          not null,
   "FROM"               VARCHAR2(36)          not null,
   "TO"                 VARCHAR2(36)          not null,
   "DATE"               DATE                  not null,
   PILOT                VARCHAR2(36)          not null,
   PLANE                VARCHAR2(36)          not null,
   TERMINAL             VARCHAR,
   constraint PK_FLIGHT primary key (ID)
);

/*==============================================================*/
/* Table: PLANE                                                 */
/*==============================================================*/
create table PLANE (
   ID                   VARCHAR2(36)          not null,
   NAME                 VARCHAR(255),
   MODEL                VARCHAR(255)          not null,
   SEATS                NUMBER(10)            not null,
   AIRLINE              VARCHAR2(36),
   constraint PK_PLANE primary key (ID)
);

/*==============================================================*/
/* Table: ROLE                                                  */
/*==============================================================*/
create table ROLE (
   ID                   VARCHAR2(36)          not null,
   NAME                 VARCHAR(255)          not null,
   LABEL                VARCHAR(255)          not null,
   constraint PK_ROLE primary key (ID)
);

/*==============================================================*/
/* Table: TERMINAL                                              */
/*==============================================================*/
create table TERMINAL (
   ID                   VARCHAR               not null,
   NAME                 VARCHAR               not null,
   constraint PK_TERMINAL primary key (ID)
);

/*==============================================================*/
/* Table: TICKET                                                */
/*==============================================================*/
create table TICKET (
   ID                   VARCHAR2(36)          not null,
   "USER"               VARCHAR2(36)          not null,
   FLIGHT               VARCHAR2(36)          not null,
   constraint PK_TICKET primary key (ID)
);

/*==============================================================*/
/* Table: "USER"                                                */
/*==============================================================*/
create table "USER" (
   ID                   VARCHAR2(36)          not null,
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
   ID                   VARCHAR2(36)          not null,
   "USER"               VARCHAR2(36)          not null,
   ROLE                 VARCHAR2(36)          not null,
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

alter table FLIGHT
   add constraint FK_FLIGHT_REFERENCE_TERMINAL foreign key (TERMINAL)
      references TERMINAL (ID);

alter table PLANE
   add constraint FK_PLANE_AIRLINE foreign key (AIRLINE)
      references AIRLINE (ID);

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

