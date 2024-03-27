package merge

import (
	"app/internal/domain"
	"log"
)

func (m Merger) mergePages(bookID int, a, b map[int]domain.Page) map[int]domain.Page {
	handled := make(map[int]struct{})
	res := make(map[int]domain.Page, len(a))

	for number, pageA := range a {
		handled[number] = struct{}{}

		pageB, ok := b[number]
		if !ok {
			res[number] = pageA
			continue
		}

		res[number] = m.mergePage(bookID, pageA, pageB)
	}

	for number, pageB := range b {
		if _, ok := handled[number]; ok {
			continue
		}

		res[number] = pageB
	}

	return res
}

func (m Merger) mergePage(bookID int, a, b domain.Page) domain.Page {
	c := domain.Page{
		PageNumber: a.PageNumber,
		Ext:        a.Ext,
		Url:        a.Url,
		LoadAt:     mergeDate(a.LoadAt, b.LoadAt),
	}

	if m.logDiff {
		if a.PageNumber != b.PageNumber {
			log.Printf("book %d page > number: %d != %d\n", bookID, a.PageNumber, b.PageNumber)
		}

		if a.Ext != b.Ext {
			log.Printf("book %d page > ext: %s != %s\n", bookID, a.Ext, b.Ext)
		}

		if a.Url != b.Url {
			log.Printf("book %d page > url: %s != %s\n", bookID, a.Url, b.Url)
		}
	}

	return c
}
