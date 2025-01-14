package database

import (
	"log"
	"mindenairport/models"
)

func (db Database) GetMaintenanceLogById(id string) models.MaintenanceLog {
	var maintenanceLog models.MaintenanceLog

	err := db.QueryRow("SELECT * FROM maintenanceLog WHERE id = :1", id).Scan(&maintenanceLog.ID, &maintenanceLog.PlaneID, &maintenanceLog.MaintenanceDate, &maintenanceLog.Description, &maintenanceLog.Technician, &maintenanceLog.NextMaintenance)
	if err != nil {
		log.Fatal("Error querying the database:", err)
	}

	return maintenanceLog
}

func (db Database) GetMaintenanceLogs() []models.MaintenanceLog {
	var maintenanceLogs []models.MaintenanceLog

	rows, err := db.Query("SELECT * FROM maintenanceLog")
	if err != nil {
		log.Fatal("Error querying the database:", err)
	}
	defer rows.Close()

	for rows.Next() {
		var maintenanceLog models.MaintenanceLog
		err := rows.Scan(&maintenanceLog.ID, &maintenanceLog.PlaneID, &maintenanceLog.MaintenanceDate, &maintenanceLog.Description, &maintenanceLog.Technician, &maintenanceLog.NextMaintenance)
		if err != nil {
			log.Fatal("Error scanning the database:", err)
		}
		maintenanceLogs = append(maintenanceLogs, maintenanceLog)
	}

	return maintenanceLogs
}
