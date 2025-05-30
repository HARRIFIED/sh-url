package postgresql

import (
	"database/sql"
	"errors"
	"time"

	"url-shortner-go/internal/domain"
)

type urlRepo struct {
	db *sql.DB
}

func NewURLRepo(db *sql.DB) domain.URLRepository {
	return &urlRepo{db: db}
}

func (r *urlRepo) Create(u *domain.URL) error {
	now := time.Now().Unix()
	err := r.db.QueryRow(
		`INSERT INTO urls (code, original_url, created_at) VALUES ($1, $2, $3) RETURNING id`,
		u.Code, u.Original, now,
	).Scan(&u.ID)
	if err != nil {
		return err
	}
	u.Created = now
	return nil
}

func (r *urlRepo) FindByCode(code string) (*domain.URL, error) {
	var u domain.URL
	row := r.db.QueryRow(
		`SELECT id, code, original_url, created_at FROM urls WHERE code=$1`, code,
	)
	if err := row.Scan(&u.ID, &u.Code, &u.Original, &u.Created); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &u, nil
}
