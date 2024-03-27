package postgresql

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
	CreateAt  time.Time      `db:"create_at"`
	Rate      int            `db:"rate"`
}

type Page struct {
	BookID     int          `db:"book_id"`
	PageNumber int          `db:"page_number"`
	Ext        string       `db:"ext"`
	Url        string       `db:"url"`
	Success    bool         `db:"success"`
	CreateAt   time.Time    `db:"create_at"`
	LoadAt     sql.NullTime `db:"load_at"`
	Rate       int          `db:"rate"`

	Hash sql.NullString `db:"hash"`
	Size sql.NullInt64  `db:"size"`
}

type BookAttribute struct {
	BookID int    `db:"book_id"`
	Attr   string `db:"attr"`
	Value  string `db:"value"`
}

func pageToDomain(p Page) domain.Page {
	return domain.Page{
		Number: p.PageNumber,
		Ext:    p.Ext,
		Url:    p.Url,
		LoadAt: p.LoadAt.Time,
		Hash:   p.Hash.String,
		Size:   p.Size.Int64,
	}
}

func bookToDomain(b Book) domain.Book {
	return domain.Book{
		ID:        b.ID,
		Name:      b.Name.String,
		Url:       b.Url.String,
		PageCount: int(b.PageCount.Int32),
		CreateAt:  b.CreateAt,
	}
}
