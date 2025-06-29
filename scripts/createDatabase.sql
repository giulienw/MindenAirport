/*==============================================================*/
/* DBMS name:      ORACLE Version 19c                           */
/* Created on:     03.12.2024 10:28:13                          */
/*==============================================================*/

/*==============================================================*/
/* Table: AIRLINE                                               */
/*==============================================================*/
create table AIRLINE (
   ID                   VARCHAR2(2)          not null,
   NAME                 VARCHAR2(255)         not null,
   constraint PK_AIRLINE primary key (ID),
   COUNTRY             VARCHAR2(255),
   LOGO_URL            VARCHAR2(255),
   ACTIVE              NUMBER(1) default 1,
   constraint CK_AIRLINE_ACTIVE check (ACTIVE in (0,1))
);

/*==============================================================*/
/* Table: AIRPORT                                               */
/*==============================================================*/
create table AIRPORT (
   ID                   VARCHAR2(3)          not null,
   NAME                 VARCHAR2(255),
   COUNTRY              VARCHAR2(255)         not null,
   CITY                 VARCHAR2(255)         not null,
   constraint PK_AIRPORT primary key (ID),
   TIMEZONE            VARCHAR2(50),
   ELEVATION           NUMBER,
   NUMBER_OF_TERMINALS NUMBER,
   LATITUDE            NUMBER(10,6),
   LONGITUDE           NUMBER(10,6)
);

/*==============================================================*/
/* Table: FLIGHT_STATUS                                          */
/*==============================================================*/
create table FLIGHT_STATUS (
   ID                   NUMBER                not null,
   NAME                 VARCHAR2(50)          not null,
   DESCRIPTION          VARCHAR2(255),
   constraint PK_FLIGHT_STATUS primary key (ID)
);

/*==============================================================*/
/* Table: TRAVEL_CLASS                                           */
/*==============================================================*/
create table TRAVEL_CLASS (
   ID                   NUMBER          not null,
   NAME                 VARCHAR2(50)          not null,
   DESCRIPTION          VARCHAR2(255),
   constraint PK_TRAVEL_CLASS primary key (ID)
);

/*==============================================================*/
/* Table: BAGGAGE                                               */
/*==============================================================*/
create table BAGGAGE (
   ID                   VARCHAR2(36)          not null,
   AIRPORTUSER          VARCHAR2(36)          not null,
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
   "FROM"               VARCHAR2(3)          not null,
   "TO"                 VARCHAR2(3)          not null,
   PILOT                VARCHAR2(36)          not null,
   PLANE                VARCHAR2(36)          not null,
   TERMINAL             VARCHAR2(36),
   STATUS               NUMBER,
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
   constraint PK_HANGAR primary key (ID),
   CAPACITY            NUMBER,
   SIZE_SQFT           NUMBER,
   STATUS              VARCHAR2(20) default 'ACTIVE',
   LAST_INSPECTION     DATE,
   NEXT_INSPECTION     DATE,
   constraint CK_HANGAR_STATUS check (STATUS in ('ACTIVE','MAINTENANCE','CLOSED'))
);

/*==============================================================*/
/* Table: PILOT                                                 */
/*==============================================================*/
create table PILOT (
   ID                   VARCHAR2(36)          not null,
   FIRSTNAME            VARCHAR2(255)         not null,
   LASTNAME             VARCHAR(255)          not null,
   constraint PK_PILOT primary key (ID),
   LICENSE_TYPE        VARCHAR2(50),
   LICENSE_NUMBER      VARCHAR2(50),
   LICENSE_EXPIRY      DATE,
   FLIGHT_HOURS        NUMBER default 0,
   MEDICAL_CHECK_DATE  DATE,
   constraint UQ_PILOT_LICENSE unique (LICENSE_NUMBER)
);

/*==============================================================*/
/* Table: PLANE                                                 */
/*==============================================================*/
create table PLANE (
   ID                   VARCHAR2(36)          not null,
   NAME                 VARCHAR2(255),
   MODEL                VARCHAR2(255)         not null,
   SEATS                NUMBER(10)            not null,
   AIRLINE              VARCHAR2(2),
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
   POSITION            INT                   not null,
   TYPE                 NUMBER          not null,
   constraint PK_PLOT primary key (ID),
   AREA_SQFT           NUMBER,
   STATUS              VARCHAR2(20) default 'AVAILABLE',
   LAST_MAINTENANCE    DATE,
   MAX_WEIGHT_CAPACITY NUMBER,
   UTILITIES_AVAILABLE VARCHAR2(255),
   constraint CK_PLOT_STATUS check (STATUS in ('AVAILABLE','OCCUPIED','MAINTENANCE'))
);

/*==============================================================*/
/* Table: PLOTTYPE                                              */
/*==============================================================*/
create table PLOTTYPE (
   ID                   NUMBER          not null,
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
   constraint PK_SHOP primary key (ID),
   OPENING_TIME        VARCHAR2(5),
   CLOSING_TIME        VARCHAR2(5),
   DESCRIPTION         VARCHAR2(1000),
   IS_DUTY_FREE        NUMBER(1) default 0,
   constraint CK_SHOP_DUTY_FREE check (IS_DUTY_FREE in (0,1))
);

/*==============================================================*/
/* Table: SHOPTYPE                                              */
/*==============================================================*/
create table SHOPTYPE (
   ID                   VARCHAR2(36)          not null,
   NAME                 VARCHAR2(255)         not null,
   LABEL                VARCHAR2(255)         not null,
   constraint PK_SHOPTYPE primary key (ID),
   CATEGORY            VARCHAR2(50),
   SECURITY_LEVEL      VARCHAR2(20),
   DESCRIPTION         VARCHAR2(1000),
   TYPICAL_HOURS       VARCHAR2(100),
   constraint CK_SHOPTYPE_SECURITY check (SECURITY_LEVEL in ('PRE_SECURITY','POST_SECURITY'))
);

/*==============================================================*/
/* Table: TERMINAL                                              */
/*==============================================================*/
create table TERMINAL (
   ID                   VARCHAR2(36)          not null,
   NAME                 VARCHAR2(255)         not null,
   constraint PK_TERMINAL primary key (ID),
   CAPACITY            NUMBER,
   STATUS              VARCHAR2(20) default 'ACTIVE',
   FLOOR_COUNT         NUMBER,
   SERVICES            VARCHAR2(1000),
   OPENING_HOURS       VARCHAR2(255),
   constraint CK_TERMINAL_STATUS check (STATUS in ('ACTIVE','MAINTENANCE','CLOSED'))
);

/*==============================================================*/
/* Table: TICKET                                                */
/*==============================================================*/
create table TICKET (
   ID                   VARCHAR2(36)          not null,
   AIRPORTUSER               VARCHAR2(36)          not null,
   FLIGHT               VARCHAR2(36)          not null,
   SEAT_NUMBER         VARCHAR2(10),
   TRAVEL_CLASS        NUMBER,
   PRICE              NUMBER(10,2),
   BOOKING_DATE       TIMESTAMP             default CURRENT_TIMESTAMP,
   STATUS             VARCHAR2(20)          default 'CONFIRMED',
   constraint PK_TICKET primary key (ID),
   constraint CK_TICKET_STATUS check (STATUS in ('CONFIRMED','CANCELLED','CHECKED_IN'))
);

/*==============================================================*/
/* Table: AIRPORTUSER                                                */
/*==============================================================*/
create table AIRPORTUSER (
   ID                   VARCHAR2(36)          not null,
   FIRSTNAME            VARCHAR2(255)         not null,
   LASTNAME             VARCHAR2(255)         not null,
   BIRTHDATE            DATE                  not null,
   PASSWORD             VARCHAR2(255)         not null,
   ACTIVE               NUMBER(1)             not null,
   EMAIL               VARCHAR2(255),
   PHONE               VARCHAR2(50),
   constraint PK_AIRPORTUSER primary key (ID),
   constraint CK_AIRPORTUSER_ACTIVE check (ACTIVE in (0,1))
);


/*==============================================================*/
/* Add Foreign Keys                                             */
/*==============================================================*/

alter table BAGGAGE
   add constraint FK_BAGGAGE_FLIGHT foreign key (FLIGHT)
      references FLIGHT (ID);

alter table BAGGAGE
   add constraint FK_BAGGAGE_AIRPORTUSER foreign key (AIRPORTUSER)
      references AIRPORTUSER (ID);

alter table FLIGHT
   add constraint FK_FLIGHT_FROM_AIRPORT foreign key ("FROM")
      references AIRPORT (ID);

alter table FLIGHT
   add constraint FK_FLIGHT_PILOT foreign key (PILOT)
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
   add constraint FK_TICKET_AIRPORTUSER foreign key (AIRPORTUSER)
      references AIRPORTUSER (ID);

alter table TICKET
   add constraint FK_TICKET_TRAVEL_CLASS foreign key (TRAVEL_CLASS)
      references TRAVEL_CLASS (ID);

/*==============================================================*/
/* Create Indexes                                               */
/*==============================================================*/

create index IDX_FLIGHT_DATES on FLIGHT (SCHEDULED_DEPARTURE, SCHEDULED_ARRIVAL);
create index IDX_BAGGAGE_TRACKING on BAGGAGE (TRACKING_NUMBER);
create index IDX_AIRPORTUSER_EMAIL on AIRPORTUSER (EMAIL);
create index IDX_TICKET_BOOKING on TICKET (BOOKING_DATE);

CREATE SEQUENCE travel_class_seq START WITH 1;

CREATE OR REPLACE TRIGGER travel_class_bir 
BEFORE INSERT ON TRAVEL_CLASS 
FOR EACH ROW

BEGIN
  SELECT travel_class_seq.NEXTVAL
  INTO   :new.id
  FROM   dual;
END;
/

CREATE SEQUENCE flight_status_seq START WITH 1;

CREATE OR REPLACE TRIGGER flight_status_bir 
BEFORE INSERT ON FLIGHT_STATUS 
FOR EACH ROW

BEGIN
  SELECT flight_status_seq.NEXTVAL
  INTO   :new.id
  FROM   dual;
END;
/

CREATE SEQUENCE plot_type_seq START WITH 1;

CREATE OR REPLACE TRIGGER plot_type_bir 
BEFORE INSERT ON PLOTTYPE 
FOR EACH ROW

BEGIN
  SELECT plot_type_seq.NEXTVAL
  INTO   :new.id
  FROM   dual;
END;
/

/*==============================================================*/
/* Create Views                                                 */
/*==============================================================*/

CREATE VIEW GetAirlines AS
SELECT ID,NAME,COUNTRY,LOGO_URL,ACTIVE
FROM airline;

CREATE VIEW GetAirports AS
SELECT * FROM Airport;

CREATE VIEW GetFlights AS
SELECT * FROM flight;

CREATE VIEW GetFlightStatuses AS
SELECT ID,NAME,DESCRIPTION
FROM flight_status;

CREATE VIEW GetMaintenanceLog AS
SELECT * FROM maintenance_log;

/*==============================================================*/
/* Create Views                                                 */
/*==============================================================*/

CREATE PROCEDURE getAirlineByID (id VARCHAR2)
AS
BEGIN
   SELECT * FROM airline WHERE ID =id;
END;
/

CREATE PROCEDURE getAirportByID (id VARCHAR2)
AS
BEGIN
   SELECT * FROM AIRPORT WHERE ID = id;
END;
/

CREATE PROCEDURE getUserByEmail (email VARCHAR2)
AS   
BEGIN
   SELECT ID, FIRSTNAME, LASTNAME, BIRTHDATE, PASSWORD, ACTIVE, EMAIL, PHONE 
			  FROM AIRPORTUSER WHERE EMAIL = email;
END;
/

CREATE PROCEDURE getUserById (id VARCHAR2)
AS
BEGIN 
   SELECT ID, FIRSTNAME, LASTNAME, BIRTHDATE, PASSWORD, ACTIVE, EMAIL, PHONE 
			  FROM AIRPORTUSER WHERE ID = id;
END;
/

CREATE PROCEDURE createUser(id VARCHAR2, firstName VARCHAR2, lastName VARCHAR2, birthdate DATE, password VARCHAR2, active NUMBER, email VARCHAR2, phone VARCHAR2)
AS
BEGIN
   INSERT INTO AIRPORTUSER (ID, FIRSTNAME, LASTNAME, BIRTHDATE, PASSWORD, ACTIVE, EMAIL, PHONE) 
			  VALUES (id, firstName, lastName, birthdate, password, active, email, phone);
END;
/

/* Node: In der Funktion UpdateLastLogin in auth.go wird LAST_LOGIN gesetzt. In AirportUser gibt es aber kein LAST_LOGIN */
CREATE PROCEDURE updateLastLogin(last_login VARCHAR2 , id VARCHAR2)
AS
BEGIN
   UPDATE AIRPORTUSER SET LAST_LOGIN = last_login WHERE ID = id;
END;
/

CREATE PROCEDURE deactivateUser(id VARCHAR2)
AS 
BEGIN
   UPDATE AIRPORTUSER SET ACTIVE = 0 WHERE ID = id;
END;
/

CREATE PROCEDURE checkEmailExists(email VARCHAR2)
AS
BEGIN
   SELECT COUNT(*) FROM AIRPORTUSER WHERE EMAIL = email;
END;
/

CREATE PROCEDURE getBaggabeByID(id VARCHAR2)
AS
BEGIN
   SELECT ID, AIRPORTUSER, FLIGHT, SIZE, WEIGHT, TRACKING_NUMBER, STATUS, SPECIAL_HANDLING 
			  FROM BAGGAGE WHERE ID = id;
END;
/

CREATE PROCEDURE getBaggabeByUserID(userID VARCHAR2)
AS
BEGIN
   SELECT * FROM BAGGAGE WHERE AIRPORTUSER = userID ORDER BY ID DESC;
END;
/

CREATE PROCEDURE getBaggaeByFlightID(flightID VARCHAR2)
AS
BEGIN
   SELECT ID, AIRPORTUSER, FLIGHT, SIZE, WEIGHT, TRACKING_NUMBER, STATUS, SPECIAL_HANDLING 
			  FROM BAGGAGE WHERE FLIGHT = flightID ORDER BY ID DESC;
END;
/

CREATE PROCEDURE createBaggage(id VARCHAR2, airportUser VARCHAR2, flight VARCHAR2, size INT, weight NUMBER, tracking_Number VARCHAR2, status VARCHAR2, special_Handling VARCHAR2)
AS
BEGIN
   INSERT INTO BAGGAGE (ID, AIRPORTUSER, FLIGHT, SIZE, WEIGHT, TRACKING_NUMBER, STATUS, SPECIAL_HANDLING)
			  VALUES (id, airportUser, flight, size, weight, tracking_Number, status, special_Handling);
END;
/

CREATE PROCEDURE updateBaggae(airportUser VARCHAR2, flight VARCHAR2, size INT, weight Number, tracking_Number VARCHAR2, status VARCHAR2, special_Handling VARCHAR2)
AS
BEGIN
   UPDATE BAGGAGE SET 
			  AIRPORTUSER = airportUser, FLIGHT = flight, SIZE = size, WEIGHT = weight, 
			  TRACKING_NUMBER = tracking_Number, STATUS = status, SPECIAL_HANDLING = special_Handling
			  WHERE ID = id;
END;
/

CREATE PROCEDURE getBaggageByTrackingNumber(tracking_Number VARCHAR2)
AS
BEGIN
   SELECT ID, AIRPORTUSER, FLIGHT, SIZE, WEIGHT, TRACKING_NUMBER, STATUS, SPECIAL_HANDLING 
			  FROM BAGGAGE WHERE TRACKING_NUMBER = tracking_Number;
END;
/