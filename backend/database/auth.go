package database

import (
	"database/sql"
	"fmt"
	"log"
	"mindenairport/models"
	"mindenairport/utils"
	"time"

	"github.com/google/uuid"
)

// GetUserByEmail retrieves a user by email
func (db Database) GetUserByEmail(email string) (*models.AirportUser, error) {
	query := `BEGIN GetUserByEmail(:1, :2); END;`

	var cursor *sql.Rows
	_, err := db.Exec(query, email, sql.Out{Dest: &cursor})
	if err != nil {
		return nil, err
	}
	defer cursor.Close()

	var user models.AirportUser
	var active int

	if cursor.Next() {
		err := cursor.Scan(
			&user.ID,
			&user.FirstName,
			&user.LastName,
			&user.Birthdate,
			&user.Password,
			&active,
			&user.Email,
			&user.Phone,
			&user.Role,
		)
		if err != nil {
			return nil, err
		}

		// Convert active int to bool
		user.Active = active == 1
		return &user, nil
	}

	return nil, nil // User not found
}

// GetUserByID retrieves a user by ID
func (db Database) GetUserByID(id string) (*models.AirportUser, error) {
	query := `BEGIN GetUserByID(:1, :2); END;`

	var cursor *sql.Rows
	_, err := db.Exec(query, id, sql.Out{Dest: &cursor})
	if err != nil {
		return nil, err
	}
	defer cursor.Close()

	var user models.AirportUser
	var active int

	if cursor.Next() {
		err := cursor.Scan(
			&user.ID,
			&user.FirstName,
			&user.LastName,
			&user.Birthdate,
			&user.Password,
			&active,
			&user.Email,
			&user.Phone,
			&user.Role,
		)
		if err != nil {
			return nil, err
		}

		// Convert active int to bool
		user.Active = active == 1
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
	query := `BEGIN UserExistsByEmail(:1, :2); END;`
	_, err := db.Exec(query, email, sql.Out{Dest: &exists})
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
	countQuery := `BEGIN GetUserCount(:1); END;`
	_, err := db.Exec(countQuery, sql.Out{Dest: &total})
	if err != nil {
		return nil, 0, err
	}

	// Calculate offset
	offset := (page - 1) * limit

	// Get users with pagination using stored procedure
	query := `BEGIN GetAllUsers(:1, :2, :3); END;`
	var cursor *sql.Rows
	_, err = db.Exec(query, offset, limit, sql.Out{Dest: &cursor})
	if err != nil {
		return nil, 0, err
	}
	defer cursor.Close()

	for cursor.Next() {
		var user models.AirportUser
		var active int
		err := cursor.Scan(
			&user.ID,
			&user.FirstName,
			&user.LastName,
			&user.Birthdate,
			&user.Password,
			&active,
			&user.Email,
			&user.Phone,
			&user.Role,
		)
		if err != nil {
			return nil, 0, err
		}

		// Convert active int to bool
		user.Active = active == 1
		users = append(users, user)
	}

	if err = cursor.Err(); err != nil {
		return nil, 0, err
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
	query := `BEGIN GetUserCount(:1); END;`
	_, err := db.Exec(query, sql.Out{Dest: &count})
	if err != nil {
		log.Printf("Error getting user count: %v", err)
		return 0
	}
	return count
}
