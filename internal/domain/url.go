package domain

type URL struct {
	ID       int64  `db:"id"`
	Code     string `db:"code"`
	Original string `db:"original_url"`
	Created  int64  `db:"created_at"`
}

type URLRepository interface {
	Create(u *URL) error
	FindByCode(code string) (*URL, error)
}
