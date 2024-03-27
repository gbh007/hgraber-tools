package jdb

import (
	"app/internal/domain"
	"context"
)

func (db *Database) Books(_ context.Context) ([]domain.Book, error) {
	res := make([]domain.Book, 0, len(db.data.Data))

	for _, rawBook := range db.data.Data {
		b := domain.Book{
			ID:        rawBook.ID,
			Name:      rawBook.Data.Name,
			Url:       rawBook.URL,
			PageCount: len(rawBook.Pages),
			CreateAt:  rawBook.Created,
		}

		pages := make([]domain.Page, len(rawBook.Pages))
		for i, rawPage := range rawBook.Pages {
			pages[i] = domain.Page{
				PageNumber: i + 1,
				Ext:        rawPage.Ext,
				Url:        rawPage.URL,
				LoadAt:     rawPage.LoadedAt,
			}
		}

		b.Pages = pages

		res = append(res, b)
	}

	return res, nil
}
