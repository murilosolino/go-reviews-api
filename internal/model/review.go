package model

import (
	"database/sql"
	"log/slog"

	"github.com/murilosolino/challenge-backend-7/internal/apperrors"
	"github.com/murilosolino/challenge-backend-7/internal/dto"
)

type ReviewModel struct {
	Bm BaseModel
}

func NewReviewModel(bm BaseModel) *ReviewModel {
	bm.TableName = "reviews"
	return &ReviewModel{Bm: bm}
}

func (m *ReviewModel) List() ([]dto.Review, error) {
	rows, err := m.Bm.List()
	if err != nil {
		slog.Error("[DATBASE:ERROR][ReviewModel][List()]"+apperrors.APP_ERR_BUILD_STMT, "error", err)
		return nil, err
	}
	return m.hydration(rows)
}

func (m *ReviewModel) FindById(id int) (dto.Review, error) {
	var result dto.Review
	stmt, err := m.Bm.db.Prepare("SELECT * FROM reviews WHERE id = ?")
	if err != nil {
		slog.Error("[DATBASE:ERROR][ReviewModel][FindById()]"+apperrors.APP_ERR_BUILD_STMT, "error", err)
		return dto.Review{}, err
	}

	row := stmt.QueryRow(id)
	_ = row.Scan(&result.Id, &result.Review, &result.AuthorName, &result.Url_photo)
	return result, nil
}

func (m *ReviewModel) FindRandomRegisters(limit int) ([]dto.Review, error) {
	rows, err := m.Bm.db.Query("SELECT * FROM reviews ORDER BY RAND() LIMIT ?", limit)
	if err != nil {
		slog.Error("[DATBASE:ERROR][ReviewModel][FindRandomRegisters()]"+apperrors.APP_ERR_EXEC_QUERY, "error", err)
		return nil, err
	}
	return m.hydration(rows)
}

func (m *ReviewModel) hydration(rows *sql.Rows) ([]dto.Review, error) {
	var row dto.Review
	var result []dto.Review
	for {
		next := rows.Next()
		if !next {
			break
		}
		err := rows.Scan(&row.Id, &row.Review, &row.AuthorName, &row.Url_photo)
		if err != nil {
			slog.Error("[DATBASE:ERROR][ReviewModel][hydration()]"+apperrors.APP_ERR_SCAN_SQL_RESULT, "error", err)
			return nil, err
		}
		result = append(result, row)
	}
	return result, nil
}
