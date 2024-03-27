package postgresql

import (
	"app/internal/domain"
	"context"
)

func (db *Database) getPages(ctx context.Context, bookID int) ([]domain.Page, error) {
	rawData := make([]Page, 0)
	err := db.db.SelectContext(ctx, &rawData, `SELECT * FROM pages WHERE book_id = $1 ORDER BY page_number;`, bookID)
	if err != nil {
		return nil, err
	}

	res := make([]domain.Page, len(rawData))
	domain.ConvertSlice(res, rawData, pageToDomain)

	return res, nil
}

func (db *Database) Books(ctx context.Context) ([]domain.Book, error) {
	rawData := make([]Book, 0)
	err := db.db.SelectContext(ctx, &rawData, `SELECT * FROM books;`)
	if err != nil {
		return nil, err
	}

	res := make([]domain.Book, len(rawData))
	for i, raw := range rawData {
		pages, err := db.getPages(ctx, raw.ID)
		if err != nil {
			return nil, err
		}

		b := bookToDomain(raw)
		b.Pages = pages
		res[i] = b
	}

	return res, nil
}
