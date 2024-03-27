package merge

import (
	"app/internal/domain"
	"log"
)

func (m Merger) mergeBooks(a, b map[int]mappedBook) map[int]mappedBook {
	handled := make(map[int]struct{})
	res := make(map[int]mappedBook, len(a))

	for id, bookA := range a {
		handled[id] = struct{}{}

		bookB, ok := b[id]
		if !ok {
			res[id] = bookA
			continue
		}

		res[id] = m.mergeBook(bookA, bookB)
	}

	for id, bookB := range b {
		if _, ok := handled[id]; ok {
			continue
		}

		res[id] = bookB
	}

	return res
}

func (m Merger) mergeBook(a, b mappedBook) mappedBook {
	c := mappedBook{
		Book: domain.Book{
			ID:        a.Book.ID,
			Name:      a.Book.Name,
			Url:       a.Book.Url,
			PageCount: a.Book.PageCount,
			CreateAt:  mergeDate(a.Book.CreateAt, b.Book.CreateAt),
		},
		Pages: m.mergePages(a.Book.ID, a.Pages, b.Pages),
	}

	if m.logDiff {
		if a.Book.Name != b.Book.Name {
			log.Printf("book %d > name: %s != %s\n", a.Book.ID, a.Book.Name, b.Book.Name)
		}

		if a.Book.Url != b.Book.Url {
			log.Printf("book %d > url: %s != %s\n", a.Book.ID, a.Book.Url, b.Book.Url)
		}

		if a.Book.PageCount != b.Book.PageCount {
			log.Printf("book %d > url: %d != %d\n", a.Book.ID, a.Book.PageCount, b.Book.PageCount)
		}
	}

	return c
}
