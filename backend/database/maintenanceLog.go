package database

import (
	"database/sql"
	"database/sql/driver"
	"log"
	"mindenairport/models"
	"time"
)

func (db Database) GetMaintenanceLogById(id string) models.MaintenanceLog {
	stmt, err := db.Prepare(`
	BEGIN MindenAirport.GetMaintenanceLogByID(:1, :2); END;`)
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

	var maintenanceLog models.MaintenanceLog

	err = cursor.Next(r)
	if err == nil {
		maintenanceLog.ID = r[0].(string)
		maintenanceLog.PlaneID = r[1].(string)
		if r[2] != nil {
			maintenanceLog.MaintenanceDate = r[2].(time.Time)
		}
		maintenanceLog.Description = r[3].(string)
		maintenanceLog.Technician = r[4].(string)
		if r[5] != nil {
			t := r[5].(time.Time)
			maintenanceLog.NextMaintenance = &t
		}
	}

	return maintenanceLog
}

func (db Database) GetMaintenanceLogs() []models.MaintenanceLog {
	stmt, err := db.Prepare(`
	BEGIN MindenAirport.GetMaintenanceLogs(:1); END;`)
	if err != nil {
		log.Fatal("Error preparing statement:", err)
	}

	var cursor driver.Rows
	_, err = stmt.Exec(sql.Out{Dest: &cursor})
	if err != nil {
		log.Fatal("Failed to execute prepared statement:", err)
	}

	r := make([]driver.Value, len(cursor.Columns()))
	defer cursor.Close()

	var maintenanceLogs []models.MaintenanceLog

	for {
		err := cursor.Next(r)
		if err != nil {
			break
		}
		var maintenanceLog models.MaintenanceLog
		maintenanceLog.ID = r[0].(string)
		maintenanceLog.PlaneID = r[1].(string)
		if r[2] != nil {
			maintenanceLog.MaintenanceDate = r[2].(time.Time)
		}
		maintenanceLog.Description = r[3].(string)
		maintenanceLog.Technician = r[4].(string)
		if r[5] != nil {
			t := r[5].(time.Time)
			maintenanceLog.NextMaintenance = &t
		}
		maintenanceLogs = append(maintenanceLogs, maintenanceLog)
	}

	return maintenanceLogs
}
