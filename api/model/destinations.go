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
	Bm BaseModel
}

func NewDestinationModel(bm BaseModel) *DestinationModel {
	bm.TableName = tableName
	return &DestinationModel{Bm: bm}
}

func (d *DestinationModel) ListAllDestinations() ([]DestinationRow, error) {
	rows, err := d.Bm.List()
	if err != nil {
		return nil, err
	}
	return hydration(rows)
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
