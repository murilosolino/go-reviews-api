package model

import (
	"database/sql"
	"fmt"
	"log/slog"
	"strings"
)

type ReviewsRow struct {
	Id         *int64  `json:"id"`
	Review     *string `json:"review"`
	AuthorName *string `json:"author"`
	Url_photo  *string `json:"url_photo"`
}

type ReviewModel struct {
	db *sql.DB
}

func NewReviewModel(db *sql.DB) *ReviewModel {
	return &ReviewModel{db: db}
}

func (m *ReviewModel) List() ([]ReviewsRow, error) {

	rows, err := m.db.Query("SELECT * FROM reviews")
	if err != nil {
		slog.Error("[DATBASE:ERROR][ReviewModel][List()] Erro ao realizar consulta", "error", err)
		return nil, err
	}
	return m.hydration(rows)
}

func (m *ReviewModel) FindById(id int) (ReviewsRow, error) {
	var result ReviewsRow
	stmt, err := m.db.Prepare("SELECT * FROM reviews WHERE id = ?")
	if err != nil {
		slog.Error("[DATBASE:ERROR][ReviewModel][FindById()] Erro ao montar statment", "error", err)
		return ReviewsRow{}, err
	}

	row := stmt.QueryRow(id)
	row.Scan(&result.Id, &result.Review, &result.AuthorName, &result.Url_photo)
	return result, nil
}

func (m *ReviewModel) Update(id int, r map[string]interface{}) error {
	query := "UPDATE reviews SET "
	var values []any

	for k, v := range r {
		query += fmt.Sprintf("%s = ?, ", k)
		values = append(values, v)
	}

	query = strings.TrimSuffix(query, ", ") + " WHERE id = ?"
	values = append(values, id)
	fmt.Println(query)
	_, err := m.db.Exec(query, values...)
	if err != nil {
		slog.Error("[DATBASE:ERROR][ReviewModel][Update()] Erro ao atualizar registro", "error", err)
		return err
	}
	return nil
}

func (m *ReviewModel) Save(r ReviewsRow) error {
	stmt, err := m.db.Prepare("INSERT INTO reviews (review, author_name, url_photo) VALUES (?,?,?)")
	if err != nil {
		slog.Error("[DATBASE:ERROR][ReviewModel][Save()] Erro ao montar statment", "error", err)
		return err
	}
	_, err = stmt.Exec(r.Review, r.AuthorName, r.Url_photo)
	if err != nil {
		slog.Error("[DATBASE:ERROR][ReviewModel][Save()] Erro ao salvar registro", "error", err)
		return err
	}
	return nil
}

func (m *ReviewModel) Delete(id int) error {
	stmt, err := m.db.Prepare("DELETE FROM reviews WHERE id = ?")
	if err != nil {
		slog.Error("[DATBASE:ERROR][ReviewModel][Delete()] Erro ao montar statment", "error", err)
		return err
	}

	_, err = stmt.Exec(id)
	if err != nil {
		slog.Error("[DATBASE:ERROR][ReviewModel][Save()] Erro ao deletar registro", "error", err)
		return err
	}
	return nil
}

func (m *ReviewModel) hydration(rows *sql.Rows) ([]ReviewsRow, error) {
	var row ReviewsRow
	var result []ReviewsRow
	for {
		next := rows.Next()
		if !next {
			break
		}
		err := rows.Scan(&row.Id, &row.Review, &row.AuthorName, &row.Url_photo)
		if err != nil {
			slog.Error("[DATBASE:ERROR][ReviewModel][hydration()] Erro ao scanear resultados", "error", err)
			return nil, err
		}
		result = append(result, row)
	}
	return result, nil
}
