package pg_storage

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"klp/models"
)

type LineStorage struct {
	db *sql.DB
}

func (ls *LineStorage) AddLine(line *models.Line) error {
	stmt, err := ls.db.Prepare("INSERT INTO sports_cfs(sport, coeff) VALUES ($1, $2);")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(line.Sport, line.Cf)
	if err != nil {
		return err
	}

	return nil
}

func (ls *LineStorage) GetLine(sport string) (float64, error) {
	panic("todo")
}

func InitDB(env *models.Env) (*sql.DB, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	dbConnectionStr := fmt.Sprintf("user=%s password=%s dbname=%s port=%d sslmode=disable", env.DBUser, env.DBPassword, env.DBName, env.DBPort)

	db, err := sql.Open("postgres", dbConnectionStr)
	if err != nil {
		return nil, err
	}

	if err := db.PingContext(ctx); err != nil {
		return nil, err
	}

	createQuery := `CREATE TABLE IF NOT EXISTS sport_cfs  (
		sport varchar(100),
		cfs FLOAT
	);`

	_, err = db.ExecContext(ctx, createQuery)
	if err != nil {
		return nil, err
	}

	return db, nil
}
