package matcher

import "app/internal/domain"

func (m *Matcher) MatchBooks(books []domain.Book, minMatchRate float64) (domain.MatchBooksResult, error) {
	res := domain.MatchBooksResult{}

	for _, book := range books {
		if len(book.Pages) == 0 {
			res.NotMatched = append(res.NotMatched, book)
			continue
		}

		records := m.MatchPages(book.Pages, minMatchRate)
		if len(records) == 0 {
			res.NotMatched = append(res.NotMatched, book)
			continue
		}

		res.Matched = append(res.Matched, domain.MatchBookResult{
			Book:   book,
			Result: records,
		})
	}

	return res, nil
}
