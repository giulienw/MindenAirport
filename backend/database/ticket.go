package database

import (
	"database/sql"
	"database/sql/driver"
	"log"
	"mindenairport/models"
	"strconv"
	"time"

	"github.com/godror/godror"
)

func (db Database) GetTicketByID(id string) (models.Ticket, error) {
	stmt, err := db.Prepare(`
	BEGIN 
	GetTicketByID(:1, :2); END;
	`)
	if err != nil {
		log.Fatal("Error preparing statement:", err)
	}

	var cursor driver.Rows
	_, err = stmt.Exec(id, sql.Out{Dest: &cursor})
	if err != nil {
		log.Fatal("Failed to execute prepared statement:", err)
	}

	r := make([]driver.Value, len(cursor.Columns()))
	defer cursor.Close()

	var ticket models.Ticket

	err = cursor.Next(r)
	if err == nil {
		ticket.ID = r[0].(string)
		if r[1] != nil {
			ticket.SeatNumber = r[1].(string)
		}
		ticket.From = r[2].(string)
		ticket.To = r[3].(string)
		if r[4] != nil {
			ticket.BookingDate = r[4].(time.Time)
		}
		if r[5] != nil {
			ticket.DepartureTime = r[5].(string)
		}
		if r[6] != nil {
			ticket.TravelClass = r[6].(string)
		}
		if r[7] != nil {
			ticket.Price, _ = strconv.ParseFloat(r[7].(godror.Number).String(), 64)
		}
		if r[8] != nil {
			ticket.Gate = r[8].(string)
		}
		if r[9] != nil {
			ticket.BaggageClaim = r[9].(string)
		}
		if r[10] != nil {
			ticket.Status = r[10].(string)
		}
	}

	return ticket, nil
}

// GetTicketsByUserID retrieves all tickets for a specific user
func (db Database) GetTicketsByUserID(userID string) ([]models.Ticket, error) {
	stmt, err := db.Prepare(`
	BEGIN 
	GetTicketsByUserID(:1, :2); END;
	`)
	if err != nil {
		return nil, err
	}

	var cursor driver.Rows
	_, err = stmt.Exec(userID, sql.Out{Dest: &cursor})
	if err != nil {
		return nil, err
	}

	r := make([]driver.Value, len(cursor.Columns()))
	defer cursor.Close()

	var tickets []models.Ticket

	for {
		err := cursor.Next(r)
		if err != nil {
			break
		}
		var ticket models.Ticket
		ticket.ID = r[0].(string)
		if r[1] != nil {
			ticket.SeatNumber = r[1].(string)
		}
		ticket.From = r[2].(string)
		ticket.To = r[3].(string)
		if r[4] != nil {
			ticket.BookingDate = r[4].(time.Time)
		}
		if r[5] != nil {
			ticket.DepartureTime = r[5].(string)
		}
		if r[6] != nil {
			ticket.TravelClass = r[6].(string)
		}
		if r[7] != nil {
			ticket.Price, _ = strconv.ParseFloat(r[7].(godror.Number).String(), 64)
		}
		if r[8] != nil {
			ticket.Gate = r[8].(string)
		}
		if r[9] != nil {
			ticket.BaggageClaim = r[9].(string)
		}
		if r[10] != nil {
			ticket.Status = r[10].(string)
		}
		if r[11] != nil {
			ticket.AirportUserID = r[11].(string)
		}
		if r[12] != nil {
			ticket.Flight = r[12].(string)
		}
		tickets = append(tickets, ticket)
	}

	return tickets, nil
}

// GetAllTickets retrieves all tickets with pagination for admin
func (db Database) GetAllTickets(page, limit int) ([]models.Ticket, int, error) {
	var tickets []models.Ticket
	var total int

	// First get the total count using stored procedure
	countStmt, err := db.Prepare(`BEGIN MindenAirport.GetTicketCount(:1); END;`)
	if err != nil {
		return nil, 0, err
	}
	_, err = countStmt.Exec(sql.Out{Dest: &total})
	if err != nil {
		return nil, 0, err
	}

	// Calculate offset
	offset := (page - 1) * limit

	// Get tickets with pagination using stored procedure
	stmt, err := db.Prepare(`BEGIN MindenAirport.GetAllTickets(:1, :2, :3); END;`)
	if err != nil {
		return nil, 0, err
	}

	var cursor driver.Rows
	_, err = stmt.Exec(offset, limit, sql.Out{Dest: &cursor})
	if err != nil {
		return nil, 0, err
	}

	r := make([]driver.Value, len(cursor.Columns()))
	defer cursor.Close()

	for {
		err := cursor.Next(r)
		if err != nil {
			break
		}
		var ticket models.Ticket
		ticket.ID = r[0].(string)
		if r[1] != nil {
			ticket.SeatNumber = r[1].(string)
		}
		ticket.From = r[2].(string)
		ticket.To = r[3].(string)
		if r[4] != nil {
			ticket.BookingDate = r[4].(time.Time)
		}
		if r[5] != nil {
			ticket.DepartureTime = r[5].(string)
		}
		if r[6] != nil {
			ticket.TravelClass = r[6].(string)
		}
		if r[7] != nil {
			ticket.Price, _ = strconv.ParseFloat(r[7].(godror.Number).String(), 64)
		}
		if r[8] != nil {
			ticket.Gate = r[8].(string)
		}
		if r[9] != nil {
			ticket.BaggageClaim = r[9].(string)
		}
		if r[10] != nil {
			ticket.Status = r[10].(string)
		}
		if r[11] != nil {
			ticket.AirportUserID = r[11].(string)
		}
		if r[12] != nil {
			ticket.Flight = r[12].(string)
		}
		tickets = append(tickets, ticket)
	}

	return tickets, total, nil
}

func (db Database) CalculateRevenue() (int, error) {
	var total int

	// Call stored procedure
	stmt, err := db.Prepare(`BEGIN MindenAirport.CalculateRevenue(:1); END;`)
	if err != nil {
		return 0, err
	}
	_, err = stmt.Exec(sql.Out{Dest: &total})
	if err != nil {
		return 0, err
	}

	return total, nil
}
