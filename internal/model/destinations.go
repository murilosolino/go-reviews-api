package model

import (
	"database/sql"
	"log/slog"

	"github.com/murilosolino/challenge-backend-7/internal/apperrors"
	"github.com/murilosolino/challenge-backend-7/internal/dto"
)

type DestinationModel struct {
	Bm BaseModel
}

func NewDestinationModel(bm BaseModel) *DestinationModel {
	bm.TableName = "destinations"
	return &DestinationModel{Bm: bm}
}

func (d *DestinationModel) ListAllDestinations() ([]dto.Destination, error) {
	rows, err := d.Bm.List()
	if err != nil {
		return nil, err
	}
	return hydration(rows)
}

func (d *DestinationModel) FindByName(name string) (dto.Destination, error) {
	var destination dto.Destination
	row := d.Bm.db.QueryRow("SELECT * FROM "+d.Bm.TableName+" WHERE name = ?", name)
	err := row.Scan(&destination.Id, &destination.Img, &destination.Name, &destination.Price, &destination.DescriptiveText)
	if err != nil {
		slog.Error("[DATBASE:ERROR][DestinationModel][hydration()]"+apperrors.APP_ERR_SCAN_SQL_RESULT, "error", err)
		return dto.Destination{}, err
	}
	return destination, nil
}

func hydration(rows *sql.Rows) ([]dto.Destination, error) {
	var d dto.Destination
	var destinations []dto.Destination

	for {
		r := rows.Next()
		if !r {
			break
		}
		err := rows.Scan(&d.Id, &d.Img, &d.Name, &d.Price, &d.DescriptiveText)
		if err != nil {
			slog.Error("[DATBASE:ERROR][DestinationModel][hydration()]"+apperrors.APP_ERR_SCAN_SQL_RESULT, "error", err)
			return nil, err
		}
		destinations = append(destinations, d)
	}
	return destinations, nil
}
