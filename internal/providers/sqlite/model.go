package sqlite

import (
	"app/internal/domain"
	"database/sql"
	"time"
)

type Book struct {
	ID        int            `db:"id"`
	Name      sql.NullString `db:"name"`
	Url       sql.NullString `db:"url"`
	PageCount sql.NullInt32  `db:"page_count"`
	CreateAt  string         `db:"create_at"`
	Rate      int            `db:"rate"`
}

type Page struct {
	BookID     int            `db:"book_id"`
	PageNumber int            `db:"page_number"`
	Ext        string         `db:"ext"`
	Url        string         `db:"url"`
	Success    bool           `db:"success"`
	CreateAt   string         `db:"create_at"`
	LoadAt     sql.NullString `db:"load_at"`
	Rate       int            `db:"rate"`
}

func pageToDomain(p Page) domain.Page {
	return domain.Page{
		Number: p.PageNumber,
		Ext:    p.Ext,
		Url:    p.Url,
		LoadAt: stringToTime(p.LoadAt.String),
	}
}

func bookToDomain(b Book) domain.Book {
	return domain.Book{
		ID:        b.ID,
		Name:      b.Name.String,
		Url:       b.Url.String,
		PageCount: int(b.PageCount.Int32),
		CreateAt:  stringToTime(b.CreateAt),
	}
}

func stringToTime(s string) time.Time {
	if s == "" {
		return time.Time{}
	}

	t, _ := time.ParseInLocation(time.RFC3339Nano, s, time.UTC)

	return t
}
