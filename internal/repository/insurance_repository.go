package repository

import (
	"database/sql"
	"fmt"

	"github.com/cabralfbenja/segurointeligente/internal/entities"
)

type InsuranceRepository interface {
	Insert(insurance *entities.Insurance) error
	GetAllByUserID(userID int64) ([]*entities.Insurance, error)
	Update(insurance *entities.Insurance) error
	GetById(id int64) (*entities.Insurance, error)
}

type mysqlInsuranceRepository struct {
	db *sql.DB
}

func NewMySQLInsuranceRepository(db *sql.DB) InsuranceRepository {
	return &mysqlInsuranceRepository{db: db}
}

func (r *mysqlInsuranceRepository) Insert(insurance *entities.Insurance) error {
	query := `
        INSERT INTO insurances (user_id, insurance_type, time_from, time_to, value, created_at)
        VALUES (?, ?, ?, ?, ?, NOW())
    `
	_, err := r.db.Exec(query,
		insurance.UserID,
		insurance.InsuranceType,
		insurance.TimeFrom,
		insurance.TimeTo,
		insurance.Value,
	)

	if err != nil {
		return fmt.Errorf("error inserting insurance: %w", err)
	}

	return nil
}

func (r *mysqlInsuranceRepository) GetAllByUserID(userID int64) ([]*entities.Insurance, error) {
	query := `
        SELECT id, user_id, insurance_type, time_from, time_to, value
        FROM insurances
        WHERE user_id = ?
    `
	rows, err := r.db.Query(query, userID)
	if err != nil {
		return nil, fmt.Errorf("error querying insurances: %w", err)
	}
	defer rows.Close()

	var insurances []*entities.Insurance
	for rows.Next() {
		var insurance entities.Insurance
		if err := rows.Scan(&insurance.ID, &insurance.UserID, &insurance.InsuranceType, &insurance.TimeFrom, &insurance.TimeTo, &insurance.Value); err != nil {
			return nil, fmt.Errorf("error scanning insurance: %w", err)
		}
		insurances = append(insurances, &insurance)
	}

	return insurances, nil
}

func (r *mysqlInsuranceRepository) Update(insurance *entities.Insurance) error {
	query := `
        UPDATE insurances
        SET insurance_type = ?, time_from = ?, time_to = ?, value = ?
        WHERE id = ?
    `
	_, err := r.db.Exec(query,
		insurance.InsuranceType,
		insurance.TimeFrom,
		insurance.TimeTo,
		insurance.Value,
		insurance.ID,
	)

	if err != nil {
		return fmt.Errorf("error updating insurance: %w", err)
	}

	return nil
}

func (r *mysqlInsuranceRepository) GetById(id int64) (*entities.Insurance, error) {
	query := `
        SELECT id, user_id, insurance_type, time_from, time_to, value
        FROM insurances
        WHERE id = ?
    `
	row := r.db.QueryRow(query, id)

	var insurance entities.Insurance
	if err := row.Scan(&insurance.ID, &insurance.UserID, &insurance.InsuranceType, &insurance.TimeFrom, &insurance.TimeTo, &insurance.Value); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("insurance with id %d not found", id)
		}
		return nil, fmt.Errorf("error scanning insurance: %w", err)
	}

	return &insurance, nil
}
