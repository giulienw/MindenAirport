package database

import (
	"database/sql"
	"log"
	"mindenairport/models"
)

func (db Database) GetMaintenanceLogById(id string) models.MaintenanceLog {
	query := `BEGIN GetMaintenanceLogByID(:1, :2); END;`

	var cursor *sql.Rows
	_, err := db.Exec(query, id, sql.Out{Dest: &cursor})
	if err != nil {
		log.Fatal("Error calling stored procedure:", err)
	}
	defer cursor.Close()

	var maintenanceLog models.MaintenanceLog

	if cursor.Next() {
		err := cursor.Scan(&maintenanceLog.ID, &maintenanceLog.PlaneID, &maintenanceLog.MaintenanceDate, &maintenanceLog.Description, &maintenanceLog.Technician, &maintenanceLog.NextMaintenance)
		if err != nil {
			log.Fatal("Error scanning maintenance log data:", err)
		}
	}

	return maintenanceLog
}

func (db Database) GetMaintenanceLogs() []models.MaintenanceLog {
	query := `BEGIN GetMaintenanceLogs(:1); END;`

	var cursor *sql.Rows
	_, err := db.Exec(query, sql.Out{Dest: &cursor})
	if err != nil {
		log.Fatal("Error calling stored procedure:", err)
	}
	defer cursor.Close()

	var maintenanceLogs []models.MaintenanceLog

	for cursor.Next() {
		var maintenanceLog models.MaintenanceLog
		err := cursor.Scan(&maintenanceLog.ID, &maintenanceLog.PlaneID, &maintenanceLog.MaintenanceDate, &maintenanceLog.Description, &maintenanceLog.Technician, &maintenanceLog.NextMaintenance)
		if err != nil {
			log.Fatal("Error scanning maintenance log data:", err)
		}
		maintenanceLogs = append(maintenanceLogs, maintenanceLog)
	}

	return maintenanceLogs
}
