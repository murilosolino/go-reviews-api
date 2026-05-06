package model

import (
	"database/sql"
	"fmt"
	"log/slog"
	"strings"

	"github.com/murilosolino/challenge-backend-7/internal/apperrors"
)

type BaseModel struct {
	db        *sql.DB
	TableName string
}

func NewBaseModel(db *sql.DB) *BaseModel {
	return &BaseModel{
		db:        db,
		TableName: "",
	}
}

func (bm *BaseModel) Save(m map[string]interface{}) error {
	query, values := bm.buildInsertQuery(m)
	_, err := bm.db.Exec(query, values...)
	if err != nil {
		slog.Error("[DATABASE::ERROR][BaseModel][Save()]"+apperrors.APP_ERR_SAVE_REGISTERS, "error", err)
		return err
	}
	return nil
}

func (bm *BaseModel) List() (*sql.Rows, error) {
	query := "SELECT * FROM " + bm.TableName

	rows, err := bm.db.Query(query)
	if err != nil {
		slog.Error("[DATABASE::ERROR][BaseModel][List()]"+apperrors.APP_ERR_EXEC_QUERY, "error", err)
		return &sql.Rows{}, err
	}
	return rows, nil
}

func (bm *BaseModel) Exclude(id int) error {
	query := "DELETE FROM " + bm.TableName + " WHERE id = ?"
	_, err := bm.db.Exec(query, id)
	if err != nil {
		slog.Error("[DATABASE::ERROR][BaseModel][Exclude()]"+apperrors.APP_ERR_DELETE_REGISTER, "error", err)
		return err
	}
	return nil
}

func (bm *BaseModel) Update(id int, m map[string]interface{}) error {

	query, values := bm.buildUpdateQuery(m)
	values = append(values, id)
	_, err := bm.db.Exec(query, values...)
	if err != nil {
		slog.Error("[DATABASE::ERROR][BaseModel][Update()]"+apperrors.APP_ERR_EXEC_QUERY, "error", err)
		return err
	}
	return nil
}

func (bm *BaseModel) buildInsertQuery(m map[string]interface{}) (string, []any) {
	var values []any
	query := "INSERT INTO " + bm.TableName + "("

	for k, v := range m {
		query += fmt.Sprintf("%s,", k)
		values = append(values, v)
	}

	query = strings.TrimSuffix(query, ",") + ")"
	query += " VALUES("

	for i := 0; i < len(m); i++ {
		query += "?,"
	}
	query = strings.TrimSuffix(query, ",") + ")"
	return query, values
}

func (bm *BaseModel) buildUpdateQuery(m map[string]interface{}) (string, []any) {
	query := "UPDATE " + bm.TableName + " SET "
	var values []any

	for k, v := range m {
		query += fmt.Sprintf("%s = ?, ", k)
		values = append(values, v)
	}
	query = strings.TrimSuffix(query, ", ") + " WHERE id = ?"
	return query, values
}
