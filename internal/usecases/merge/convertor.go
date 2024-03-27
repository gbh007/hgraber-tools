package merge

import (
	"app/internal/domain"
	"time"
)

type mappedBook struct {
	Book  domain.Book
	Pages map[int]domain.Page
}

func (Merger) toMap(raw []domain.Book) map[int]mappedBook {
	res := make(map[int]mappedBook, len(raw))

	for _, rawBook := range raw {
		mb := mappedBook{
			Book:  rawBook,
			Pages: make(map[int]domain.Page, len(rawBook.Pages)),
		}

		mb.Book.Pages = nil

		for _, rawPage := range rawBook.Pages {
			mb.Pages[rawPage.PageNumber] = rawPage
		}

		res[rawBook.ID] = mb
	}

	return res
}

func (Merger) fromMap(raw map[int]mappedBook) []domain.Book {
	res := make([]domain.Book, 0, len(raw))

	for _, rawBook := range raw {
		pages := make([]domain.Page, 0, len(rawBook.Pages))
		for _, rawPage := range rawBook.Pages {
			pages = append(pages, rawPage)
		}

		rawBook.Book.Pages = pages

		res = append(res, rawBook.Book)
	}

	return res
}

// Специальный конвертор, с защитой от некорректных данных,т.к. все валидные данные гарантировано после 2000
func isZeroTime(t time.Time) bool {
	return t.Before(time.Date(2000, 0, 0, 0, 0, 0, 0, time.UTC))
}
