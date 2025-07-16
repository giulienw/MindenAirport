package database

import (
	"database/sql"
	"log"
	"mindenairport/models"
)

func (db Database) GetTicketByID(id string) (models.Ticket, error) {
	query := `BEGIN GetTicketByID(:1, :2); END;`

	var cursor *sql.Rows
	_, err := db.Exec(query, id, sql.Out{Dest: &cursor})
	if err != nil {
		log.Fatal("Error calling stored procedure:", err)
	}
	defer cursor.Close()

	var ticket models.Ticket

	if cursor.Next() {
		err := cursor.Scan(&ticket.ID, &ticket.SeatNumber, &ticket.From, &ticket.To, &ticket.BookingDate, &ticket.DepartureTime, &ticket.TravelClass, &ticket.Price, &ticket.Gate, &ticket.BaggageClaim, &ticket.Status)
		if err != nil {
			log.Println("Error scanning ticket data:", err)
			return models.Ticket{}, err
		}
	}

	return ticket, nil
}

// GetTicketsByUserID retrieves all tickets for a specific user
func (db Database) GetTicketsByUserID(userID string) ([]models.Ticket, error) {
	query := `BEGIN GetTicketsByUserID(:1, :2); END;`

	var cursor *sql.Rows
	_, err := db.Exec(query, userID, sql.Out{Dest: &cursor})
	if err != nil {
		return nil, err
	}
	defer cursor.Close()

	var tickets []models.Ticket

	for cursor.Next() {
		var ticket models.Ticket
		err := cursor.Scan(
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

	if err = cursor.Err(); err != nil {
		return nil, err
	}

	return tickets, nil
}

// GetAllTickets retrieves all tickets with pagination for admin
func (db Database) GetAllTickets(page, limit int) ([]models.Ticket, int, error) {
	var tickets []models.Ticket
	var total int

	// First get the total count using stored procedure
	countQuery := `BEGIN GetTicketCount(:1); END;`
	_, err := db.Exec(countQuery, sql.Out{Dest: &total})
	if err != nil {
		return nil, 0, err
	}

	// Calculate offset
	offset := (page - 1) * limit

	// Get tickets with pagination using stored procedure
	query := `BEGIN GetAllTickets(:1, :2, :3); END;`
	var cursor *sql.Rows
	_, err = db.Exec(query, offset, limit, sql.Out{Dest: &cursor})
	if err != nil {
		return nil, 0, err
	}
	defer cursor.Close()

	for cursor.Next() {
		var ticket models.Ticket
		err := cursor.Scan(
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

	if err = cursor.Err(); err != nil {
		return nil, 0, err
	}

	return tickets, total, nil
}

func (db Database) CalculateRevenue() (int, error) {
	var total int

	// Call stored procedure
	query := `BEGIN CalculateRevenue(:1); END;`
	_, err := db.Exec(query, sql.Out{Dest: &total})
	if err != nil {
		return 0, err
	}

	return total, nil
}
