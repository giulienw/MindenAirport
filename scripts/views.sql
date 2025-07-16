-- View for baggage with pagination count
CREATE VIEW GetBaggageCount AS
SELECT COUNT(*) as TOTAL_COUNT FROM BAGGAGE;

-- View for tickets with pagination count  
CREATE VIEW GetTicketsCount AS
SELECT COUNT(*) as TOTAL_COUNT FROM TICKET;

-- View for users with pagination count
CREATE VIEW GetUsersCount AS
SELECT COUNT(*) as TOTAL_COUNT FROM AIRPORTUSER;

-- View for flights with pagination count
CREATE VIEW GetFlightsCount AS
SELECT COUNT(*) as TOTAL_COUNT FROM FLIGHT;

-- View for flight statistics
CREATE VIEW FlightStatistics AS
SELECT 
    COUNT(*) as TOTAL_FLIGHTS,
    COUNT(CASE WHEN STATUS = 1 THEN 1 END) as ACTIVE_FLIGHTS,
    COUNT(CASE WHEN ACTUAL_DEPARTURE > SCHEDULED_DEPARTURE THEN 1 END) as DELAYED_FLIGHTS
FROM FLIGHT;

-- View for revenue calculation
CREATE VIEW RevenueStatistics AS
SELECT 
    COALESCE(SUM(PRICE), 0) as TOTAL_REVENUE,
    COUNT(*) as TOTAL_TICKETS
FROM TICKET;

-- View for baggage statistics
CREATE VIEW BaggageStatistics AS
SELECT 
    COUNT(*) as TOTAL_BAGGAGE,
    COUNT(CASE WHEN STATUS = 'LOST' THEN 1 END) as LOST_BAGGAGE
FROM BAGGAGE;