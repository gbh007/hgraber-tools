package matcher

import (
	"app/internal/domain"
	"slices"
)

func (m *Matcher) MatchPages(pages []domain.Page, minMatchRate float64) []domain.MatchRecord {
	matchedBooks := make(map[int]int, 10) // Ожидается что не будет сильно много совпадений
	pageCountFloat := float64(len(pages))

	for _, page := range pages {
		foundedBooks := make(map[int]struct{}, 10)

		for _, info := range m.hashes[page.Hash] {
			if _, ok := foundedBooks[info.BookID]; ok { // Возможен случай когда у книги несколько страниц с одинаковым хешом
				continue
			}

			foundedBooks[info.BookID] = struct{}{}
			matchedBooks[info.BookID]++
		}
	}

	res := make([]domain.MatchRecord, 0, len(matchedBooks))
	for id, count := range matchedBooks {
		book := m.books[id]

		rate := float64(count) / pageCountFloat
		if rate < minMatchRate {
			continue
		}

		res = append(res, domain.MatchRecord{
			ID:        id,
			Name:      book.Name,
			Url:       book.Url,
			MatchRate: rate,
		})
	}

	slices.SortFunc(res, func(a, b domain.MatchRecord) int {
		if a.MatchRate > b.MatchRate {
			return -1
		}

		if a.MatchRate < b.MatchRate {
			return 1
		}

		return 0
	})

	return res
}
