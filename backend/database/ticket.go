package database

import (
	"log"
	"mindenairport/models"
)

func (db Database) GetTicketByID(id string) models.Ticket {
	var ticket models.Ticket

	query := `SELECT TICKET.ID,
		TICKET.SEAT_NUMBER,
		FLIGHT."FROM",FLIGHT."TO",TICKET.BOOKING_DATE, 
		CASE WHEN FLIGHT.SCHEDULED_DEPARTURE = FLIGHT.ACTUAL_DEPARTURE THEN TO_CHAR(FLIGHT.SCHEDULED_DEPARTURE, 'dd.mm.yyyy HH24:MI') ELSE TO_CHAR(FLIGHT.ACTUAL_DEPARTURE, 'dd.mm.yyyy HH24:MI') END AS DEPARTURE_TIME,
		TRAVEL_CLASS.NAME AS TRAVEL_CLASS,TICKET.PRICE,
		FLIGHT.GATE,
		FLIGHT.BAGGAGE_CLAIM,
		TICKET.STATUS FROM FLIGHT RIGHT JOIN TICKET ON TICKET.FLIGHT = FLIGHT.ID RIGHT JOIN TRAVEL_CLASS ON TICKET.TRAVEL_CLASS = TRAVEL_CLASS.ID 
		WHERE TICKET.ID = :1;`
	err := db.QueryRow(query, id).Scan(&ticket.ID, &ticket.SeatNumber, &ticket.From, &ticket.To, &ticket.BookingDate, &ticket.DepartureTime, &ticket.TravelClass, &ticket.Price, &ticket.Gate, &ticket.BaggageClaim, &ticket.Status)
	if err != nil {
		log.Fatal("Error querying the database:", err)
	}

	return ticket
}

// GetTicketsByUserID retrieves all tickets for a specific user
func (db Database) GetTicketsByUserID(userID string) ([]models.Ticket, error) {
	var tickets []models.Ticket

	query := `SELECT 
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
	WHERE TICKET.AIRPORTUSER = :1
	ORDER BY TICKET.BOOKING_DATE DESC`

	rows, err := db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var ticket models.Ticket
		err := rows.Scan(
			&ticket.ID,
			&ticket.SeatNumber,
			&ticket.From,
			&ticket.To,
			&ticket.BookingDate,
			&ticket.DepartureTime,
			&ticket.TravelClass,
			&ticket.Price,
			&ticket.Gate,
			&ticket.BaggageClaim,
			&ticket.Status,
			&ticket.AirportUserID,
			&ticket.Flight,
		)
		if err != nil {
			return nil, err
		}
		tickets = append(tickets, ticket)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return tickets, nil
}

// GetAllTickets retrieves all tickets with pagination for admin
func (db Database) GetAllTickets(page, limit int) ([]models.Ticket, int, error) {
	var tickets []models.Ticket
	var total int

	// First get the total count
	countQuery := `SELECT COUNT(*) FROM TICKET`
	err := db.QueryRow(countQuery).Scan(&total)
	if err != nil {
		return nil, 0, err
	}

	// Calculate offset
	offset := (page - 1) * limit

	// Get tickets with pagination and flight information
	query := `SELECT 
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
	OFFSET :1 ROWS FETCH NEXT :2 ROWS ONLY`

	rows, err := db.Query(query, offset, limit)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	for rows.Next() {
		var ticket models.Ticket
		err := rows.Scan(
			&ticket.ID,
			&ticket.SeatNumber,
			&ticket.From,
			&ticket.To,
			&ticket.BookingDate,
			&ticket.DepartureTime,
			&ticket.TravelClass,
			&ticket.Price,
			&ticket.Gate,
			&ticket.BaggageClaim,
			&ticket.Status,
			&ticket.AirportUserID,
			&ticket.Flight,
		)
		if err != nil {
			return nil, 0, err
		}
		tickets = append(tickets, ticket)
	}

	if err = rows.Err(); err != nil {
		return nil, 0, err
	}

	return tickets, total, nil
}
