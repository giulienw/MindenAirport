// Package database provides authentication-related database operations
// for user management in the MindenAirport system.
package database

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"log"
	"mindenairport/models"
	"mindenairport/utils"
	"time"

	"github.com/google/uuid"
)

// GetUserByEmail retrieves a user from the database by their email address.
// Used primarily for login authentication to validate user credentials.
//
// Parameters:
//   - email: The user's email address to search for
//
// Returns:
//   - *models.AirportUser: The user record if found, nil if not found
//   - error: Any database error that occurred during the operation
func (db Database) GetUserByEmail(email string) (*models.AirportUser, error) {
	stmt, err := db.Prepare(`
	BEGIN 
	GetUserByEmail(:1, :2); END;
	`)
	if err != nil {
		return nil, err
	}

	var cursor driver.Rows
	_, err = stmt.Exec(email, sql.Out{Dest: &cursor})
	if err != nil {
		return nil, err
	}

	r := make([]driver.Value, len(cursor.Columns()))
	defer cursor.Close()

	var user models.AirportUser

	err = cursor.Next(r)
	if err == nil {
		user.ID = r[0].(string)
		user.FirstName = r[1].(string)
		user.LastName = r[2].(string)
		if r[3] != nil {
			user.Birthdate = r[3].(time.Time)
		}
		user.Password = r[4].(string)
		user.Active = r[5].(int64) == 1
		user.Email = r[6].(string)
		user.Phone = r[7].(string)
		user.Role = r[8].(string)
		return &user, nil
	}

	return nil, nil // User not found
}

// GetUserByID retrieves a user from the database by their unique ID.
// Used for token validation and user profile operations.
//
// Parameters:
//   - id: The unique user identifier
//
// Returns:
//   - *models.AirportUser: The user record if found, nil if not found
//   - error: Any database error that occurred during the operation
func (db Database) GetUserByID(id string) (*models.AirportUser, error) {
	stmt, err := db.Prepare(`
	BEGIN 
	GetUserByID(:1, :2); END;
	`)
	if err != nil {
		return nil, err
	}

	var cursor driver.Rows
	_, err = stmt.Exec(id, sql.Out{Dest: &cursor})
	if err != nil {
		return nil, err
	}

	r := make([]driver.Value, len(cursor.Columns()))
	defer cursor.Close()

	var user models.AirportUser

	err = cursor.Next(r)
	if err == nil {
		user.ID = r[0].(string)
		user.FirstName = r[1].(string)
		user.LastName = r[2].(string)
		if r[3] != nil {
			user.Birthdate = r[3].(time.Time)
		}
		user.Password = r[4].(string)
		user.Active = r[5].(int64) == 1
		user.Email = r[6].(string)
		user.Phone = r[7].(string)
		user.Role = r[8].(string)
		return &user, nil
	}

	return nil, nil // User not found
}

// CreateUser creates a new user
func (db Database) CreateUser(req models.RegisterRequest) (*models.AirportUser, error) {
	// Hash the password
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	// Generate new UUID for user
	userID := uuid.New().String()

	// Call stored procedure
	_, err = db.Exec("BEGIN CreateUserWithRole(:1, :2, :3, :4, :5, :6, :7, :8, :9); END;",
		userID,
		req.FirstName,
		req.LastName,
		req.Birthdate,
		hashedPassword,
		1, // Active by default
		req.Email,
		req.Phone,
		"USER", // Default role
	)

	if err != nil {
		log.Printf("Error creating user: %v", err)
		return nil, err
	}

	// Return the created user
	user := &models.AirportUser{
		ID:        userID,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Birthdate: req.Birthdate,
		Password:  hashedPassword,
		Active:    true,
		Email:     req.Email,
		Phone:     req.Phone,
		Role:      "USER",
	}

	return user, nil
}

// UpdateUserLastLogin updates the user's last login time
func (db Database) UpdateUserLastLogin(userID string) error {
	query := `UPDATE AIRPORTUSER SET LAST_LOGIN = :1 WHERE ID = :2`
	_, err := db.Exec(query, time.Now(), userID)
	return err
}

// DeactivateUser deactivates a user account
func (db Database) DeactivateUser(userID string, active int) error {
	fmt.Println("Deactivating user:", userID, "Active:", active)
	query := `BEGIN SetUserActiveStatus(:1, :2); END;`
	_, err := db.Exec(query, userID, active)
	return err
}

// CheckEmailExists checks if an email already exists in the database
func (db Database) CheckEmailExists(email string) (bool, error) {
	var exists int
	stmt, err := db.Prepare(`BEGIN UserExistsByEmail(:1, :2); END;`)
	if err != nil {
		return false, err
	}
	_, err = stmt.Exec(email, sql.Out{Dest: &exists})
	if err != nil {
		return false, err
	}
	return exists == 1, nil
}

// GetAllUsers retrieves all users with pagination for admin
func (db Database) GetAllUsers(page, limit int) ([]models.AirportUser, int, error) {
	var users []models.AirportUser
	var total int

	// First get the total count using stored procedure
	countStmt, err := db.Prepare(`BEGIN GetUserCount(:1); END;`)
	if err != nil {
		return nil, 0, err
	}
	_, err = countStmt.Exec(sql.Out{Dest: &total})
	if err != nil {
		return nil, 0, err
	}

	// Calculate offset
	offset := (page - 1) * limit

	// Get users with pagination using stored procedure
	stmt, err := db.Prepare(`
	BEGIN 
	GetAllUsers(:1, :2, :3); END;
	`)
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
		var user models.AirportUser
		user.ID = r[0].(string)
		user.FirstName = r[1].(string)
		user.LastName = r[2].(string)
		if r[3] != nil {
			user.Birthdate = r[3].(time.Time)
		}
		user.Password = r[4].(string)
		user.Active = r[5].(int64) == 1
		user.Email = r[6].(string)
		user.Phone = r[7].(string)
		user.Role = r[8].(string)
		users = append(users, user)
	}

	return users, total, nil
}

// UpdateUserByAdmin updates user information by admin
func (db Database) UpdateUserByAdmin(userID, firstName, lastName, email, phone string, active *bool, role string) error {
	activeValue := 1
	if active != nil && !*active {
		activeValue = 0
	}

	// Call stored procedure
	query := `BEGIN UpdateUserByAdmin(:1, :2, :3, :4, :5, :6, :7); END;`
	_, err := db.Exec(query, userID, firstName, lastName, email, phone, activeValue, role)
	return err
}

func (db Database) GetUserCount() int {
	var count int
	stmt, err := db.Prepare(`BEGIN GetUserCount(:1); END;`)
	if err != nil {
		log.Printf("Error preparing statement: %v", err)
		return 0
	}
	_, err = stmt.Exec(sql.Out{Dest: &count})
	if err != nil {
		log.Printf("Error getting user count: %v", err)
		return 0
	}
	return count
}
