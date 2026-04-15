package model

import (
	"database/sql"
	"log/slog"

	"github.com/murilosolino/challenge-backend-7/api/apperrors"
)

const tableName = "destinations"

type DestinationRow struct {
	Id    int     `json:"id"`
	Img   *string `json:"image"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type DestinationModel struct {
	db *sql.DB
}

func NewDestinationModel(db *sql.DB) *DestinationModel {
	return &DestinationModel{db: db}
}

func (d DestinationModel) Save(destination DestinationRow) error {
	query := "INSERT INTO " + tableName +
		" (img, name, price) VALUES (?,?,?)"

	_, err := d.db.Exec(query, destination.Img, destination.Name, destination.Price)
	if err != nil {
		slog.Error("[DATABASE::ERROR][DestinationModel][Save()]"+apperrors.APP_ERR_SAVE_REGISTERS, "error", err)
		return err
	}
	return nil
}

func (d DestinationModel) List() ([]DestinationRow, error) {
	query := "SELECT * FROM " + tableName

	rows, err := d.db.Query(query)
	if err != nil {
		slog.Error("[DATABASE::ERROR][DestinationModel][List()]"+apperrors.APP_ERR_EXEC_QUERY, "error", err)
		return nil, err
	}
	return hydration(rows)
}

func (d DestinationModel) Exclude(id int) error {
	query := "DELETE FROM " + tableName + " WHERE id = ?"
	_, err := d.db.Exec(query, id)
	if err != nil {
		slog.Error("[DATABASE::ERROR][DestinationModel][Exclude()]"+apperrors.APP_ERR_DELETE_REGISTER, "error", err)
		return err
	}
	return nil
}

func hydration(rows *sql.Rows) ([]DestinationRow, error) {
	var d DestinationRow
	var destinations []DestinationRow

	for {
		r := rows.Next()
		if !r {
			break
		}
		err := rows.Scan(&d.Id, &d.Img, &d.Name, &d.Price)
		if err != nil {
			slog.Error("[DATBASE:ERROR][DestinationModel][hydration()]"+apperrors.APP_ERR_SCAN_SQL_RESULT, "error", err)
			return nil, err
		}
		destinations = append(destinations, d)
	}
	return destinations, nil
}
