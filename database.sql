/*==============================================================*/
/* DBMS name:      ORACLE Version 19c                           */
/* Created on:     03.12.2024 10:28:13                          */
/*==============================================================*/

/* Drop existing foreign key constraints */
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

/* Drop existing tables */
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

/* Drop new tables if they exist */
drop table FLIGHT_STATUS cascade constraints;
drop table TRAVEL_CLASS cascade constraints;
drop table MAINTENANCE_LOG cascade constraints;
drop table CREW_MEMBER cascade constraints;
drop table FLIGHT_CREW cascade constraints;

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
/* Table: FLIGHT_STATUS                                          */
/*==============================================================*/
create table FLIGHT_STATUS (
   ID                   VARCHAR2(36)          not null,
   NAME                 VARCHAR2(50)          not null,
   DESCRIPTION          VARCHAR2(255),
   constraint PK_FLIGHT_STATUS primary key (ID)
);

/*==============================================================*/
/* Table: TRAVEL_CLASS                                           */
/*==============================================================*/
create table TRAVEL_CLASS (
   ID                   VARCHAR2(36)          not null,
   NAME                 VARCHAR2(50)          not null,
   DESCRIPTION          VARCHAR2(255),
   constraint PK_TRAVEL_CLASS primary key (ID)
);

/*==============================================================*/
/* Table: BAGGAGE                                               */
/*==============================================================*/
create table BAGGAGE (
   ID                   VARCHAR2(36)          not null,
   "USER"               VARCHAR2(36)          not null,
   FLIGHT               VARCHAR2(36)          not null,
   "SIZE"               INT                   not null,
   WEIGHT              NUMBER(5,2)           not null,
   TRACKING_NUMBER     VARCHAR2(20)          not null,
   STATUS              VARCHAR2(20)          default 'CHECKED',
   SPECIAL_HANDLING    VARCHAR2(255),
   constraint PK_BAGGAGE primary key (ID),
   constraint CK_BAGGAGE_STATUS check (STATUS in ('CHECKED','IN_TRANSIT','DELIVERED','LOST'))
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
   STATUS               VARCHAR2(36),
   SCHEDULED_DEPARTURE  TIMESTAMP             not null,
   ACTUAL_DEPARTURE     TIMESTAMP,
   SCHEDULED_ARRIVAL    TIMESTAMP             not null,
   ACTUAL_ARRIVAL       TIMESTAMP,
   GATE                 VARCHAR2(10),
   BAGGAGE_CLAIM        VARCHAR2(10),
   constraint PK_FLIGHT primary key (ID)
);

/*==============================================================*/
/* Table: MAINTENANCE_LOG                                        */
/*==============================================================*/
create table MAINTENANCE_LOG (
   ID                   VARCHAR2(36)          not null,
   PLANE                VARCHAR2(36)          not null,
   MAINTENANCE_DATE     DATE                  not null,
   DESCRIPTION          VARCHAR2(1000)        not null,
   TECHNICIAN          VARCHAR2(255)         not null,
   NEXT_MAINTENANCE    DATE,
   constraint PK_MAINTENANCE_LOG primary key (ID)
);

/*==============================================================*/
/* Table: CREW_MEMBER                                           */
/*==============================================================*/
create table CREW_MEMBER (
   ID                   VARCHAR2(36)          not null,
   FIRSTNAME            VARCHAR2(255)         not null,
   LASTNAME             VARCHAR2(255)         not null,
   ROLE                 VARCHAR2(50)          not null,
   LICENSE_NUMBER       VARCHAR2(50),
   LICENSE_EXPIRY       DATE,
   constraint PK_CREW_MEMBER primary key (ID)
);

/*==============================================================*/
/* Table: FLIGHT_CREW                                           */
/*==============================================================*/
create table FLIGHT_CREW (
   ID                   VARCHAR2(36)          not null,
   FLIGHT               VARCHAR2(36)          not null,
   CREW_MEMBER          VARCHAR2(36)          not null,
   ROLE                 VARCHAR2(50)          not null,
   constraint PK_FLIGHT_CREW primary key (ID)
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
   MANUFACTURING_YEAR   NUMBER(4),
   MAX_TAKEOFF_WEIGHT  NUMBER(10,2),
   FUEL_CAPACITY       NUMBER(10,2),
   STATUS              VARCHAR2(20)          default 'ACTIVE',
   constraint PK_PLANE primary key (ID),
   constraint CK_PLANE_STATUS check (STATUS in ('ACTIVE','MAINTENANCE','INACTIVE'))
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
   NAME                 VARCHAR2(255)         not null,
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
   SEAT_NUMBER         VARCHAR2(10),
   TRAVEL_CLASS        VARCHAR2(36),
   PRICE              NUMBER(10,2),
   BOOKING_DATE       TIMESTAMP             default CURRENT_TIMESTAMP,
   STATUS             VARCHAR2(20)          default 'CONFIRMED',
   constraint PK_TICKET primary key (ID),
   constraint CK_TICKET_STATUS check (STATUS in ('CONFIRMED','CANCELLED','CHECKED_IN'))
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
   EMAIL               VARCHAR2(255),
   PHONE               VARCHAR2(50),
   constraint PK_USER primary key (ID)
);

/*==============================================================*/
/* Add Foreign Keys                                             */
/*==============================================================*/

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
   add constraint FK_FLIGHT_STATUS foreign key (STATUS)
      references FLIGHT_STATUS (ID);

alter table FLIGHT
   add constraint FK_FLIGHT_FK_FLIGHT_TERMINAL foreign key (TERMINAL)
      references TERMINAL (ID);

alter table FLIGHT
   add constraint FK_FLIGHT_TO_AIRPORT foreign key ("TO")
      references AIRPORT (ID);

alter table HANGAR
   add constraint FK_HANGAR_PLOT foreign key (PLOT)
      references PLOT (ID);

alter table MAINTENANCE_LOG
   add constraint FK_MAINTENANCE_PLANE foreign key (PLANE)
      references PLANE (ID);

alter table FLIGHT_CREW
   add constraint FK_FLIGHT_CREW_FLIGHT foreign key (FLIGHT)
      references FLIGHT (ID);

alter table FLIGHT_CREW
   add constraint FK_FLIGHT_CREW_MEMBER foreign key (CREW_MEMBER)
      references CREW_MEMBER (ID);

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

alter table TICKET
   add constraint FK_TICKET_TRAVEL_CLASS foreign key (TRAVEL_CLASS)
      references TRAVEL_CLASS (ID);

/*==============================================================*/
/* Create Indexes                                               */
/*==============================================================*/

create index IDX_FLIGHT_DATES on FLIGHT (SCHEDULED_DEPARTURE, SCHEDULED_ARRIVAL);
create index IDX_BAGGAGE_TRACKING on BAGGAGE (TRACKING_NUMBER);
create index IDX_USER_EMAIL on "USER" (EMAIL);
create index IDX_TICKET_BOOKING on TICKET (BOOKING_DATE);