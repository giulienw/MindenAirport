/*==============================================================*/
/* DBMS name:      ORACLE Version 19c                           */
/* Created on:     03.12.2024 10:28:13                          */
/*==============================================================*/


alter table BAGGAGE
   drop constraint FK_BAGGAGE_FLIGHT;

alter table BAGGAGE
   drop constraint FK_BAGGAGE_USER;

alter table FLIGHT
   drop constraint FK_FLIGHT_FROM_AIRPORT;

alter table FLIGHT
   drop constraint FK_FLIGHT_USER;

alter table FLIGHT
   drop constraint FK_FLIGHT_PLANE;

alter table FLIGHT
   drop constraint FK_FLIGHT_FK_FLIGHT_TERMINAL;

alter table FLIGHT
   drop constraint FK_FLIGHT_TO_AIRPORT;

alter table HANGAR
   drop constraint FK_HANGAR_PLOT;

alter table PLANE
   drop constraint FK_PLANE_AIRLINE;

alter table PLANE
   drop constraint FK_PLANE_HANGAR;

alter table PLOT
   drop constraint FK_PLOTS_PLOTTYPE;

alter table SHOP
   drop constraint FK_SHOPS_PLOTS;

alter table SHOP
   drop constraint FK_SHOPS_SHOPTYPE;

alter table TICKET
   drop constraint FK_TICKET_FLIGHT;

alter table TICKET
   drop constraint FK_TICKET_USER;

drop table AIRLINE cascade constraints;

drop table AIRPORT cascade constraints;

drop table BAGGAGE cascade constraints;

drop table FLIGHT cascade constraints;

drop table HANGAR cascade constraints;

drop table PILOT cascade constraints;

drop table PLANE cascade constraints;

drop table PLOT cascade constraints;

drop table PLOTTYPE cascade constraints;

drop table SHOP cascade constraints;

drop table SHOPTYPE cascade constraints;

drop table TERMINAL cascade constraints;

drop table TICKET cascade constraints;

drop table "USER" cascade constraints;

/*==============================================================*/
/* Table: AIRLINE                                               */
/*==============================================================*/
create table AIRLINE (
   ID                   VARCHAR2(36)          not null,
   NAME                 VARCHAR2(255)         not null,
   constraint PK_AIRLINE primary key (ID)
);

/*==============================================================*/
/* Table: AIRPORT                                               */
/*==============================================================*/
create table AIRPORT (
   ID                   VARCHAR2(36)          not null,
   NAME                 VARCHAR2(255),
   COUNTRY              VARCHAR2(255)         not null,
   CITY                 VARCHAR2(255)         not null,
   constraint PK_AIRPORT primary key (ID)
);

/*==============================================================*/
/* Table: BAGGAGE                                               */
/*==============================================================*/
create table BAGGAGE (
   ID                   VARCHAR2(36)          not null,
   "USER"               VARCHAR2(36)          not null,
   FLIGHT               VARCHAR2(36)          not null,
   "SIZE"               INT                   not null,
   constraint PK_BAGGAGE primary key (ID)
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
   TERMINAL             VARCHAR2(36),
   constraint PK_FLIGHT primary key (ID)
);

/*==============================================================*/
/* Table: HANGAR                                                */
/*==============================================================*/
create table HANGAR (
   ID                   VARCHAR2(36)          not null,
   PLOT                 VARCHAR2(36)          not null,
   constraint PK_HANGAR primary key (ID)
);

/*==============================================================*/
/* Table: PILOT                                                 */
/*==============================================================*/
create table PILOT (
   ID                   VARCHAR2(36)          not null,
   FIRSTNAME            VARCHAR2(255)         not null,
   LASTNAME             VARCHAR(255)          not null,
   constraint PK_PILOT primary key (ID)
);

/*==============================================================*/
/* Table: PLANE                                                 */
/*==============================================================*/
create table PLANE (
   ID                   VARCHAR2(36)          not null,
   NAME                 VARCHAR2(255),
   MODEL                VARCHAR2(255)         not null,
   SEATS                NUMBER(10)            not null,
   AIRLINE              VARCHAR2(36),
   HANGAR               VARCHAR2(36),
   constraint PK_PLANE primary key (ID)
);

/*==============================================================*/
/* Table: PLOT                                                  */
/*==============================================================*/
create table PLOT (
   ID                   VARCHAR2(36)          not null,
   POSTITION            INT                   not null,
   TYPE                 VARCHAR2(36)          not null,
   constraint PK_PLOT primary key (ID)
);

/*==============================================================*/
/* Table: PLOTTYPE                                              */
/*==============================================================*/
create table PLOTTYPE (
   ID                   VARCHAR2(36)          not null,
   NAME                 VARCHAR2(255)         not null,
   LABEL                VARCHAR2(255)         not null,
   constraint PK_PLOTTYPE primary key (ID)
);

/*==============================================================*/
/* Table: SHOP                                                  */
/*==============================================================*/
create table SHOP (
   ID                   VARCHAR2(36)          not null,
   NAME                 VARCHAR2(255          not null,
   TYPE                 VARCHAR2(36)          not null,
   PLOT                 VARCHAR2(36)          not null,
   constraint PK_SHOP primary key (ID)
);

/*==============================================================*/
/* Table: SHOPTYPE                                              */
/*==============================================================*/
create table SHOPTYPE (
   ID                   VARCHAR2(36)          not null,
   NAME                 VARCHAR2(255)         not null,
   LABEL                VARCHAR2(255)         not null,
   constraint PK_SHOPTYPE primary key (ID)
);

/*==============================================================*/
/* Table: TERMINAL                                              */
/*==============================================================*/
create table TERMINAL (
   ID                   VARCHAR2(36)          not null,
   NAME                 VARCHAR2(255)         not null,
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
   FIRSTNAME            VARCHAR2(255)         not null,
   LASTNAME             VARCHAR2(255)         not null,
   BIRTHDATE            DATE                  not null,
   PASSWORD             VARCHAR2(255)         not null,
   ACTIVE               BINARY(1)             not null,
   constraint PK_USER primary key (ID)
);

alter table BAGGAGE
   add constraint FK_BAGGAGE_FLIGHT foreign key (FLIGHT)
      references FLIGHT (ID);

alter table BAGGAGE
   add constraint FK_BAGGAGE_USER foreign key ("USER")
      references "USER" (ID);

alter table FLIGHT
   add constraint FK_FLIGHT_FROM_AIRPORT foreign key ("FROM")
      references AIRPORT (ID);

alter table FLIGHT
   add constraint FK_FLIGHT_USER foreign key (PILOT)
      references PILOT (ID);

alter table FLIGHT
   add constraint FK_FLIGHT_PLANE foreign key (PLANE)
      references PLANE (ID);

alter table FLIGHT
   add constraint FK_FLIGHT_FK_FLIGHT_TERMINAL foreign key (TERMINAL)
      references TERMINAL (ID);

alter table FLIGHT
   add constraint FK_FLIGHT_TO_AIRPORT foreign key ("TO")
      references AIRPORT (ID);

alter table HANGAR
   add constraint FK_HANGAR_PLOT foreign key (PLOT)
      references PLOT (ID);

alter table PLANE
   add constraint FK_PLANE_AIRLINE foreign key (AIRLINE)
      references AIRLINE (ID);

alter table PLANE
   add constraint FK_PLANE_HANGAR foreign key (HANGAR)
      references HANGAR (ID);

alter table PLOT
   add constraint FK_PLOTS_PLOTTYPE foreign key (TYPE)
      references PLOTTYPE (ID);

alter table SHOP
   add constraint FK_SHOPS_PLOTS foreign key (PLOT)
      references PLOT (ID);

alter table SHOP
   add constraint FK_SHOPS_SHOPTYPE foreign key (TYPE)
      references SHOPTYPE (ID);

alter table TICKET
   add constraint FK_TICKET_FLIGHT foreign key (FLIGHT)
      references FLIGHT (ID);

alter table TICKET
   add constraint FK_TICKET_USER foreign key ("USER")
      references "USER" (ID);

