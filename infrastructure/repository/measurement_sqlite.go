package repository

import (
	"database/sql"
	"log"
	"time"

	"visitors.it-zt.at/domain/entity"
)

type MeasurementSQLite struct {
	db *sql.DB
}

func NewMeasurementSQLite(db *sql.DB) *MeasurementSQLite {
	return &MeasurementSQLite{
		db: db,
	}
}

func (r *MeasurementSQLite) Create(e *entity.Measurement) (entity.ID, error) {
	tx, err := r.db.Begin()
	if err != nil {
		log.Fatalf("Could not create transaction: %v", err)
	}

	stmt, err := r.db.Prepare(`
		INSERT INTO visits (id, quantity, created_at) 
		values(?,?,?)`)

	if err != nil {
		return e.ID, err
	}
	_, err = stmt.Exec(
		e.ID,
		e.Quantity,
		time.Now().Unix(),
	)
	tx.Commit()

	if err != nil {
		return e.ID, err
	}
	err = stmt.Close()
	if err != nil {
		return e.ID, err
	}
	return e.ID, nil
}
