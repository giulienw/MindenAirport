/* Drop existing foreign key constraints */
alter table BAGGAGE
   drop constraint FK_BAGGAGE_FLIGHT;

alter table BAGGAGE
   drop constraint FK_BAGGAGE_AIRPORTUSER;

alter table FLIGHT
   drop constraint FK_FLIGHT_FROM_AIRPORT;

alter table FLIGHT
   drop constraint FK_FLIGHT_PILOT;

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
   drop constraint FK_TICKET_AIRPORTUSER;

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
drop table AIRPORTUSER cascade constraints;

/* Drop new tables if they exist */
drop table FLIGHT_STATUS cascade constraints;
drop table TRAVEL_CLASS cascade constraints;
drop table MAINTENANCE_LOG cascade constraints;
drop table CREW_MEMBER cascade constraints;
drop table FLIGHT_CREW cascade constraints;

drop sequence travel_class_seq;
drop sequence flight_status_seq;
drop sequence plot_type_seq;