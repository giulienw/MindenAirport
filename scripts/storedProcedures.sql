
-- Fixed CreateUser procedure to include ROLE
CREATE OR REPLACE PROCEDURE CreateUserWithRole(
    id VARCHAR2, 
    firstName VARCHAR2, 
    lastName VARCHAR2, 
    birthdate DATE, 
    password VARCHAR2, 
    active NUMBER, 
    email VARCHAR2, 
    phone VARCHAR2,
    role VARCHAR2
)
AS
BEGIN
    INSERT INTO AIRPORTUSER (ID, FIRSTNAME, LASTNAME, BIRTHDATE, PASSWORD, ACTIVE, EMAIL, PHONE, ROLE) 
    VALUES (id, firstName, lastName, birthdate, password, active, email, phone, role);
END;
/

-- Update user by admin procedure
CREATE OR REPLACE PROCEDURE UpdateUserByAdmin(
    userID VARCHAR2,
    firstName VARCHAR2,
    lastName VARCHAR2,
    email VARCHAR2,
    phone VARCHAR2,
    active NUMBER,
    role VARCHAR2
)
AS
BEGIN
    UPDATE AIRPORTUSER SET 
        FIRSTNAME = firstName,
        LASTNAME = lastName,
        EMAIL = email,
        PHONE = phone,
        ACTIVE = active,
        ROLE = role
    WHERE ID = userID;
END;
/

-- Get all users with pagination
CREATE OR REPLACE PROCEDURE GetAllUsers(
    page_offset NUMBER,
    page_limit NUMBER,
    result_cursor OUT SYS_REFCURSOR
)
AS
BEGIN
    OPEN result_cursor FOR
    SELECT ID, FIRSTNAME, LASTNAME, BIRTHDATE, PASSWORD, ACTIVE, EMAIL, PHONE, ROLE 
    FROM AIRPORTUSER 
    ORDER BY ID 
    OFFSET page_offset ROWS FETCH NEXT page_limit ROWS ONLY;
END;
/

-- Get user count procedure
CREATE OR REPLACE PROCEDURE GetUserCount(
    user_count OUT NUMBER
)
AS
BEGIN
    SELECT COUNT(*) INTO user_count FROM AIRPORTUSER;
END;
/

-- Get user by email procedure
CREATE OR REPLACE PROCEDURE GetUserByEmail(
    p_email VARCHAR2,
    result_cursor OUT SYS_REFCURSOR
)
AS
BEGIN
    OPEN result_cursor FOR
    SELECT ID, FIRSTNAME, LASTNAME, BIRTHDATE, PASSWORD, ACTIVE, EMAIL, PHONE, ROLE 
    FROM AIRPORTUSER 
    WHERE EMAIL = p_email;
END;
/

-- Get user by ID procedure
CREATE OR REPLACE PROCEDURE GetUserByID(
    p_id VARCHAR2,
    result_cursor OUT SYS_REFCURSOR
)
AS
BEGIN
    OPEN result_cursor FOR
    SELECT ID, FIRSTNAME, LASTNAME, BIRTHDATE, PASSWORD, ACTIVE, EMAIL, PHONE, ROLE 
    FROM AIRPORTUSER 
    WHERE ID = p_id;
END;
/

-- Create flight procedure
CREATE OR REPLACE PROCEDURE CreateFlight(
    id VARCHAR2,
    flight_from VARCHAR2,
    flight_to VARCHAR2,
    pilot VARCHAR2,
    plane VARCHAR2,
    terminal VARCHAR2,
    status NUMBER,
    scheduled_departure TIMESTAMP,
    actual_departure TIMESTAMP,
    scheduled_arrival TIMESTAMP,
    actual_arrival TIMESTAMP,
    gate VARCHAR2,
    baggage_claim VARCHAR2
)
AS
BEGIN
    INSERT INTO FLIGHT (ID, "FROM", "TO", PILOT, PLANE, TERMINAL, STATUS, SCHEDULED_DEPARTURE, ACTUAL_DEPARTURE, SCHEDULED_ARRIVAL, ACTUAL_ARRIVAL, GATE, BAGGAGE_CLAIM) 
    VALUES (id, flight_from, flight_to, pilot, plane, terminal, status, scheduled_departure, actual_departure, scheduled_arrival, actual_arrival, gate, baggage_claim);
END;
/

-- Get all flights with pagination
CREATE OR REPLACE PROCEDURE GetAllFlights(
    page_offset NUMBER,
    page_limit NUMBER,
    result_cursor OUT SYS_REFCURSOR
)
AS
BEGIN
    OPEN result_cursor FOR
    SELECT ID, "FROM", "TO", PILOT, PLANE, TERMINAL, STATUS, SCHEDULED_DEPARTURE, ACTUAL_DEPARTURE, SCHEDULED_ARRIVAL, ACTUAL_ARRIVAL, GATE, BAGGAGE_CLAIM 
    FROM FLIGHT 
    ORDER BY ID DESC 
    OFFSET page_offset ROWS FETCH NEXT page_limit ROWS ONLY;
END;
/

-- Get flight by ID procedure
CREATE OR REPLACE PROCEDURE GetFlightByID(
    p_id VARCHAR2,
    result_cursor OUT SYS_REFCURSOR
)
AS
BEGIN
    OPEN result_cursor FOR
    SELECT ID, "FROM", "TO", PILOT, PLANE, TERMINAL, STATUS, SCHEDULED_DEPARTURE, ACTUAL_DEPARTURE, SCHEDULED_ARRIVAL, ACTUAL_ARRIVAL, GATE, BAGGAGE_CLAIM 
    FROM FLIGHT 
    WHERE ID = p_id;
END;
/

-- Get flight count procedure
CREATE OR REPLACE PROCEDURE GetFlightCount(
    flight_count OUT NUMBER
)
AS
BEGIN
    SELECT COUNT(*) INTO flight_count FROM FLIGHT;
END;
/

-- Update flight procedure
CREATE OR REPLACE PROCEDURE UpdateFlight(
    p_id VARCHAR2,
    flight_from VARCHAR2,
    flight_to VARCHAR2,
    pilot VARCHAR2,
    plane VARCHAR2,
    terminal VARCHAR2,
    status NUMBER,
    scheduled_departure TIMESTAMP,
    actual_departure TIMESTAMP,
    scheduled_arrival TIMESTAMP,
    actual_arrival TIMESTAMP,
    gate VARCHAR2,
    baggage_claim VARCHAR2
)
AS
BEGIN
    UPDATE FLIGHT SET 
        "FROM" = flight_from,
        "TO" = flight_to,
        PILOT = pilot,
        PLANE = plane,
        TERMINAL = terminal,
        STATUS = status,
        SCHEDULED_DEPARTURE = scheduled_departure,
        ACTUAL_DEPARTURE = actual_departure,
        SCHEDULED_ARRIVAL = scheduled_arrival,
        ACTUAL_ARRIVAL = actual_arrival,
        GATE = gate,
        BAGGAGE_CLAIM = baggage_claim
    WHERE ID = p_id;
END;
/

-- Delete flight procedure
CREATE OR REPLACE PROCEDURE DeleteFlight(
    p_id VARCHAR2
)
AS
BEGIN
    DELETE FROM FLIGHT WHERE ID = p_id;
END;
/

-- Get all tickets with pagination and details
CREATE OR REPLACE PROCEDURE GetAllTickets(
    page_offset NUMBER,
    page_limit NUMBER,
    result_cursor OUT SYS_REFCURSOR
)
AS
BEGIN
    OPEN result_cursor FOR
    SELECT 
        TICKET.ID,
        TICKET.SEAT_NUMBER,
        FLIGHT."FROM",
        FLIGHT."TO",
        TICKET.BOOKING_DATE,
        CASE 
            WHEN FLIGHT.SCHEDULED_DEPARTURE = FLIGHT.ACTUAL_DEPARTURE 
            THEN TO_CHAR(FLIGHT.SCHEDULED_DEPARTURE, 'dd.mm.yyyy HH24:MI') 
            ELSE TO_CHAR(FLIGHT.ACTUAL_DEPARTURE, 'dd.mm.yyyy HH24:MI') 
        END AS DEPARTURE_TIME,
        TRAVEL_CLASS.NAME AS TRAVEL_CLASS,
        TICKET.PRICE,
        FLIGHT.GATE,
        FLIGHT.BAGGAGE_CLAIM,
        TICKET.STATUS,
        TICKET.AIRPORTUSER,
        TICKET.FLIGHT
    FROM TICKET 
    LEFT JOIN FLIGHT ON TICKET.FLIGHT = FLIGHT.ID 
    LEFT JOIN TRAVEL_CLASS ON TICKET.TRAVEL_CLASS = TRAVEL_CLASS.ID 
    ORDER BY TICKET.BOOKING_DATE DESC
    OFFSET page_offset ROWS FETCH NEXT page_limit ROWS ONLY;
END;
/

-- Get ticket count procedure
CREATE OR REPLACE PROCEDURE GetTicketCount(
    ticket_count OUT NUMBER
)
AS
BEGIN
    SELECT COUNT(*) INTO ticket_count FROM TICKET;
END;
/

-- Calculate total revenue
CREATE OR REPLACE PROCEDURE CalculateRevenue(
    total_revenue OUT NUMBER
)
AS
BEGIN
    SELECT COALESCE(SUM(PRICE), 0) INTO total_revenue FROM TICKET;
END;
/

-- Get ticket by ID procedure
CREATE OR REPLACE PROCEDURE GetTicketByID(
    p_id VARCHAR2,
    result_cursor OUT SYS_REFCURSOR
)
AS
BEGIN
    OPEN result_cursor FOR
    SELECT 
        TICKET.ID,
        TICKET.SEAT_NUMBER,
        FLIGHT."FROM",
        FLIGHT."TO",
        TICKET.BOOKING_DATE,
        CASE 
            WHEN FLIGHT.SCHEDULED_DEPARTURE = FLIGHT.ACTUAL_DEPARTURE 
            THEN TO_CHAR(FLIGHT.SCHEDULED_DEPARTURE, 'dd.mm.yyyy HH24:MI') 
            ELSE TO_CHAR(FLIGHT.ACTUAL_DEPARTURE, 'dd.mm.yyyy HH24:MI') 
        END AS DEPARTURE_TIME,
        TRAVEL_CLASS.NAME AS TRAVEL_CLASS,
        TICKET.PRICE,
        FLIGHT.GATE,
        FLIGHT.BAGGAGE_CLAIM,
        TICKET.STATUS
    FROM TICKET 
    LEFT JOIN FLIGHT ON TICKET.FLIGHT = FLIGHT.ID 
    LEFT JOIN TRAVEL_CLASS ON TICKET.TRAVEL_CLASS = TRAVEL_CLASS.ID 
    WHERE TICKET.ID = p_id;
END;
/

-- Get tickets by user ID procedure
CREATE OR REPLACE PROCEDURE GetTicketsByUserID(
    p_user_id VARCHAR2,
    result_cursor OUT SYS_REFCURSOR
)
AS
BEGIN
    OPEN result_cursor FOR
    SELECT 
        TICKET.ID,
        TICKET.SEAT_NUMBER,
        FLIGHT."FROM",
        FLIGHT."TO",
        TICKET.BOOKING_DATE,
        CASE 
            WHEN FLIGHT.SCHEDULED_DEPARTURE = FLIGHT.ACTUAL_DEPARTURE 
            THEN TO_CHAR(FLIGHT.SCHEDULED_DEPARTURE, 'dd.mm.yyyy HH24:MI') 
            ELSE TO_CHAR(FLIGHT.ACTUAL_DEPARTURE, 'dd.mm.yyyy HH24:MI') 
        END AS DEPARTURE_TIME,
        TRAVEL_CLASS.NAME AS TRAVEL_CLASS,
        TICKET.PRICE,
        FLIGHT.GATE,
        FLIGHT.BAGGAGE_CLAIM,
        TICKET.STATUS,
        TICKET.AIRPORTUSER,
        TICKET.FLIGHT
    FROM TICKET 
    LEFT JOIN FLIGHT ON TICKET.FLIGHT = FLIGHT.ID 
    LEFT JOIN TRAVEL_CLASS ON TICKET.TRAVEL_CLASS = TRAVEL_CLASS.ID 
    WHERE TICKET.AIRPORTUSER = p_user_id
    ORDER BY TICKET.BOOKING_DATE DESC;
END;
/

-- Delete baggage procedure
CREATE OR REPLACE PROCEDURE DeleteBaggage(
    baggage_id VARCHAR2
)
AS
BEGIN
    DELETE FROM BAGGAGE WHERE ID = baggage_id;
END;
/

-- Get all baggage with pagination
CREATE OR REPLACE PROCEDURE GetAllBaggage(
    page_offset NUMBER,
    page_limit NUMBER,
    result_cursor OUT SYS_REFCURSOR
)
AS
BEGIN
    OPEN result_cursor FOR
    SELECT ID, AIRPORTUSER, FLIGHT, SIZE, WEIGHT, TRACKING_NUMBER, STATUS, SPECIAL_HANDLING 
    FROM BAGGAGE 
    ORDER BY ID DESC
    OFFSET page_offset ROWS FETCH NEXT page_limit ROWS ONLY;
END;
/

-- Get baggage count procedure
CREATE OR REPLACE PROCEDURE GetBaggageCount(
    baggage_count OUT NUMBER
)
AS
BEGIN
    SELECT COUNT(*) INTO baggage_count FROM BAGGAGE;
END;
/

-- Get baggage by ID procedure
CREATE OR REPLACE PROCEDURE GetBaggageByID(
    p_id VARCHAR2,
    result_cursor OUT SYS_REFCURSOR
)
AS
BEGIN
    OPEN result_cursor FOR
    SELECT ID, AIRPORTUSER, FLIGHT, SIZE, WEIGHT, TRACKING_NUMBER, STATUS, SPECIAL_HANDLING 
    FROM BAGGAGE 
    WHERE ID = p_id;
END;
/

-- Get baggage by user ID procedure
CREATE OR REPLACE PROCEDURE GetBaggageByUserID(
    p_user_id VARCHAR2,
    result_cursor OUT SYS_REFCURSOR
)
AS
BEGIN
    OPEN result_cursor FOR
    SELECT ID, AIRPORTUSER, FLIGHT, SIZE, WEIGHT, TRACKING_NUMBER, STATUS, SPECIAL_HANDLING 
    FROM BAGGAGE 
    WHERE AIRPORTUSER = p_user_id 
    ORDER BY ID DESC;
END;
/

-- Get baggage by flight ID procedure
CREATE OR REPLACE PROCEDURE GetBaggageByFlightID(
    p_flight_id VARCHAR2,
    result_cursor OUT SYS_REFCURSOR
)
AS
BEGIN
    OPEN result_cursor FOR
    SELECT ID, AIRPORTUSER, FLIGHT, SIZE, WEIGHT, TRACKING_NUMBER, STATUS, SPECIAL_HANDLING 
    FROM BAGGAGE 
    WHERE FLIGHT = p_flight_id 
    ORDER BY ID DESC;
END;
/

-- Get baggage by tracking number procedure
CREATE OR REPLACE PROCEDURE GetBaggageByTrackingNumber(
    p_tracking_number VARCHAR2,
    result_cursor OUT SYS_REFCURSOR
)
AS
BEGIN
    OPEN result_cursor FOR
    SELECT ID, AIRPORTUSER, FLIGHT, SIZE, WEIGHT, TRACKING_NUMBER, STATUS, SPECIAL_HANDLING 
    FROM BAGGAGE 
    WHERE TRACKING_NUMBER = p_tracking_number;
END;
/

-- Create baggage procedure
CREATE OR REPLACE PROCEDURE CreateBaggage(
    p_id VARCHAR2,
    p_airportuser VARCHAR2,
    p_flight VARCHAR2,
    p_size NUMBER,
    p_weight NUMBER,
    p_tracking_number VARCHAR2,
    p_status VARCHAR2,
    p_special_handling VARCHAR2
)
AS
BEGIN
    INSERT INTO BAGGAGE (ID, AIRPORTUSER, FLIGHT, SIZE, WEIGHT, TRACKING_NUMBER, STATUS, SPECIAL_HANDLING)
    VALUES (p_id, p_airportuser, p_flight, p_size, p_weight, p_tracking_number, p_status, p_special_handling);
END;
/

-- Update baggage procedure
CREATE OR REPLACE PROCEDURE UpdateBaggage(
    p_id VARCHAR2,
    p_airportuser VARCHAR2,
    p_flight VARCHAR2,
    p_size NUMBER,
    p_weight NUMBER,
    p_tracking_number VARCHAR2,
    p_status VARCHAR2,
    p_special_handling VARCHAR2
)
AS
BEGIN
    UPDATE BAGGAGE SET 
        AIRPORTUSER = p_airportuser,
        FLIGHT = p_flight,
        SIZE = p_size,
        WEIGHT = p_weight,
        TRACKING_NUMBER = p_tracking_number,
        STATUS = p_status,
        SPECIAL_HANDLING = p_special_handling
    WHERE ID = p_id;
END;
/

-- Get all airports procedure
CREATE OR REPLACE PROCEDURE GetAllAirports(
    result_cursor OUT SYS_REFCURSOR
)
AS
BEGIN
    OPEN result_cursor FOR
    SELECT ID, NAME, COUNTRY, CITY, TIMEZONE, ELEVATION, NUMBER_OF_TERMINALS, LATITUDE, LONGITUDE 
    FROM AIRPORT 
    ORDER BY NAME;
END;
/

-- Get airport by ID procedure
CREATE OR REPLACE PROCEDURE GetAirportByID(
    p_id VARCHAR2,
    result_cursor OUT SYS_REFCURSOR
)
AS
BEGIN
    OPEN result_cursor FOR
    SELECT ID, NAME, COUNTRY, CITY, TIMEZONE, ELEVATION, NUMBER_OF_TERMINALS, LATITUDE, LONGITUDE 
    FROM AIRPORT 
    WHERE ID = p_id;
END;
/

-- Get all airlines procedure
CREATE OR REPLACE PROCEDURE GetAllAirlines(
    result_cursor OUT SYS_REFCURSOR
)
AS
BEGIN
    OPEN result_cursor FOR
    SELECT ID, NAME, COUNTRY, LOGO_URL, ACTIVE 
    FROM AIRLINE 
    ORDER BY NAME;
END;
/

-- Get airline by ID procedure
CREATE OR REPLACE PROCEDURE GetAirlineByID(
    p_id VARCHAR2,
    result_cursor OUT SYS_REFCURSOR
)
AS
BEGIN
    OPEN result_cursor FOR
    SELECT ID, NAME, COUNTRY, LOGO_URL, ACTIVE 
    FROM AIRLINE 
    WHERE ID = p_id;
END;
/

-- Get all maintenance logs
CREATE OR REPLACE PROCEDURE GetMaintenanceLogs(
    result_cursor OUT SYS_REFCURSOR
)
AS
BEGIN
    OPEN result_cursor FOR
    SELECT * FROM MAINTENANCE_LOG;
END;
/

-- Get maintenance log by ID procedure
CREATE OR REPLACE PROCEDURE GetMaintenanceLogByID(
    p_id VARCHAR2,
    result_cursor OUT SYS_REFCURSOR
)
AS
BEGIN
    OPEN result_cursor FOR
    SELECT * FROM MAINTENANCE_LOG 
    WHERE ID = p_id;
END;
/

-- Get all flight statuses procedure
CREATE OR REPLACE PROCEDURE GetFlightStatuses(
    result_cursor OUT SYS_REFCURSOR
)
AS
BEGIN
    OPEN result_cursor FOR
    SELECT ID, NAME, DESCRIPTION FROM FLIGHT_STATUS;
END;
/

-- Get flight status by ID procedure
CREATE OR REPLACE PROCEDURE GetFlightStatusByID(
    p_id NUMBER,
    result_cursor OUT SYS_REFCURSOR
)
AS
BEGIN
    OPEN result_cursor FOR
    SELECT ID, NAME, DESCRIPTION FROM FLIGHT_STATUS 
    WHERE ID = p_id;
END;
/

-- Fix UpdateFlight procedure to match new schema
CREATE OR REPLACE PROCEDURE UpdateFlightFixed (
   flight_from VARCHAR2,
   flight_to VARCHAR2,
   pilotID VARCHAR2,
   planeID VARCHAR2,
   terminalID VARCHAR2,
   statusID NUMBER,
   scheduledDeparture TIMESTAMP,
   actualDeparture TIMESTAMP,
   scheduledArrival TIMESTAMP,
   actualArrival TIMESTAMP,
   gate VARCHAR2,
   baggageClaim VARCHAR2,
   flight_id VARCHAR2
)
AS
BEGIN
   UPDATE FLIGHT SET
      "FROM" = flight_from,
      "TO" = flight_to,
      PILOT = pilotID,
      PLANE = planeID,
      TERMINAL = terminalID,
      STATUS = statusID,
      SCHEDULED_DEPARTURE = scheduledDeparture,
      ACTUAL_DEPARTURE = actualDeparture,
      SCHEDULED_ARRIVAL = scheduledArrival,
      ACTUAL_ARRIVAL = actualArrival,
      GATE = gate,
      BAGGAGE_CLAIM = baggageClaim
   WHERE ID = flight_id;
END;
/

-- Fix GetAirlineByID to use correct column mapping
CREATE OR REPLACE PROCEDURE GetAirlineByIDFixed (
   airline_id VARCHAR2,
   result_cursor OUT SYS_REFCURSOR
)
AS
BEGIN
   OPEN result_cursor FOR
   SELECT ID, NAME, COUNTRY, LOGO_URL, ACTIVE FROM AIRLINE WHERE ID = airline_id;
END;
/

-- Removed duplicate procedure GetAirportByIDFixed. Ensure the corrected logic is applied to GetAirportByID.

-- Fix GetFlightByID to use correct column mapping
CREATE OR REPLACE PROCEDURE GetFlightByIDFixed (
   flight_id VARCHAR2,
   result_cursor OUT SYS_REFCURSOR
)
AS
BEGIN
   OPEN result_cursor FOR
   SELECT ID, "FROM", "TO", PILOT, PLANE, TERMINAL, STATUS, SCHEDULED_DEPARTURE, ACTUAL_DEPARTURE, SCHEDULED_ARRIVAL, ACTUAL_ARRIVAL, GATE, BAGGAGE_CLAIM 
   FROM FLIGHT WHERE ID = flight_id;
END;
/

-- Fix GetMaintenanceLogByID to use correct table name
CREATE OR REPLACE PROCEDURE GetMaintenanceLogByIDFixed (
   log_id VARCHAR2,
   result_cursor OUT SYS_REFCURSOR
)
AS
BEGIN
   OPEN result_cursor FOR
   SELECT * FROM MAINTENANCE_LOG WHERE ID = log_id;
END;
/

/*==============================================================*/
/* Admin Dashboard Procedures                                   */
/*==============================================================*/

-- Get admin dashboard statistics
CREATE OR REPLACE PROCEDURE GetAdminDashboardStats(
    total_flights OUT NUMBER,
    active_flights OUT NUMBER,
    total_passengers OUT NUMBER,
    total_baggage OUT NUMBER,
    delayed_flights OUT NUMBER,
    lost_baggage OUT NUMBER,
    total_revenue OUT NUMBER
)
AS
BEGIN
    -- Get total flights
    SELECT COUNT(*) INTO total_flights FROM FLIGHT;
    
    -- Get active flights (assuming status 1 is active)
    SELECT COUNT(*) INTO active_flights FROM FLIGHT WHERE STATUS = 1;
    
    -- Get total passengers (unique users with tickets)
    SELECT COUNT(DISTINCT AIRPORTUSER) INTO total_passengers FROM TICKET;
    
    -- Get total baggage
    SELECT COUNT(*) INTO total_baggage FROM BAGGAGE;
    
    -- Get delayed flights
    SELECT COUNT(*) INTO delayed_flights 
    FROM FLIGHT 
    WHERE ACTUAL_DEPARTURE > SCHEDULED_DEPARTURE;
    
    -- Get lost baggage
    SELECT COUNT(*) INTO lost_baggage FROM BAGGAGE WHERE STATUS = 'LOST';
    
    -- Get total revenue
    SELECT COALESCE(SUM(PRICE), 0) INTO total_revenue FROM TICKET;
END;
/

/*==============================================================*/
/* Utility Procedures                                           */
/*==============================================================*/

-- Check if user exists by email
CREATE OR REPLACE PROCEDURE UserExistsByEmail(
    user_email VARCHAR2,
    user_exists OUT NUMBER
)
AS
BEGIN
    SELECT COUNT(*) INTO user_exists FROM AIRPORTUSER WHERE EMAIL = user_email;
END;
/

-- Activate/Deactivate user
CREATE OR REPLACE PROCEDURE SetUserActiveStatus(
    user_id VARCHAR2,
    active_status NUMBER
)
AS
BEGIN
    UPDATE AIRPORTUSER SET ACTIVE = active_status WHERE ID = user_id;
END;
/