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
