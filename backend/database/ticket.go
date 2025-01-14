package database

import (
	"log"
	"mindenairport/models"
)

func (db Database) GetTicketByID(id string) models.Ticket {
	var ticket models.Ticket

	err := db.QueryRow("SELECT TICKET.ID,TICKET.SEAT_NUMBER,FLIGHT.FROM,FLIGHT.TO,TO_CHAR(TICKET.BOOKING_DATE,'dd.mm.yyyy HH24:MI') AS BOOKING_DATE, CASE WHEN FLIGHT.SCHEDULED_DEPARTURE = FLIGHT.ACTUAL_DEPARTURE THEN TO_CHAR(FLIGHT.SCHEDULED_DEPARTURE, 'dd.mm.yyyy HH24:MI') ELSE TO_CHAR(FLIGHT.ACTUAL_DEPARTURE, 'dd.mm.yyyy HH24:MI') END AS DEPARTURE_TIME,TRAVEL_CLASS.NAME AS TRAVEL_CLASS,TICKET.PRICE,FLIGHT.GATE,FLIGHT.BAGGAGE_CLAIM,TICKET.STATUS FROM FLIGHT RIGHT JOIN TICKET ON TICKET.FLIGHT = FLIGHT.ID RIGHT JOIN TRAVEL_CLASS ON TICKET.TRAVEL_CLASS = TRAVEL_CLASS.ID WHERE TICKET.ID = :1;", id).Scan(&ticket.ID, &ticket.SeatNumber, &ticket.From, &ticket.To, &ticket.BookingDate, &ticket.DepartureTime, &ticket.TravelClass, &ticket.Price, &ticket.Gate, &ticket.BaggageClaim, &ticket.Status)
	if err != nil {
		log.Fatal("Error querying the database:", err)
	}

	return ticket
}
