package jsondata

import (
	"app/internal/domain"
	"slices"
)

func convertMatchRecord(in []domain.MatchRecord) []MatchRecord {
	out := make([]MatchRecord, len(in))

	domain.ConvertSlice(out, in, func(raw domain.MatchRecord) MatchRecord {
		return MatchRecord(raw)
	})

	slices.SortFunc(out, func(a, b MatchRecord) int {
		if a.MatchRate > b.MatchRate {
			return -1
		}

		if a.MatchRate < b.MatchRate {
			return 1
		}

		return 0
	})

	return out
}

func convertMatchFileResult(in []domain.MatchFileResult) []MatchFileResult {
	out := make([]MatchFileResult, len(in))

	domain.ConvertSlice(out, in, func(raw domain.MatchFileResult) MatchFileResult {
		return MatchFileResult{
			Path:   raw.Path,
			Result: convertMatchRecord(raw.Result),
		}
	})

	return out
}

func convertBook(in domain.Book) Book {
	pages := make([]Page, len(in.Pages))
	domain.ConvertSlice(pages, in.Pages, func(p domain.Page) Page {
		return Page{
			Number: p.Number,
			Ext:    p.Ext,
			Url:    p.Url,
			LoadAt: p.LoadAt,
			Hash:   p.Hash,
			Size:   p.Size,
		}
	})

	slices.SortFunc(pages, func(a, b Page) int {
		return a.Number - b.Number
	})

	return Book{
		ID:         in.ID,
		Name:       in.Name,
		Url:        in.Url,
		PageCount:  in.PageCount,
		CreateAt:   in.CreateAt,
		Pages:      pages,
		Attributes: Attributes(in.Attributes),
	}
}

func convertBooks(in []domain.Book) []Book {
	out := make([]Book, len(in))

	domain.ConvertSlice(out, in, convertBook)

	slices.SortFunc(out, func(a, b Book) int {
		return a.ID - b.ID
	})

	return out
}

func convertMatchBookResult(in []domain.MatchBookResult) []MatchBookResult {
	out := make([]MatchBookResult, len(in))

	domain.ConvertSlice(out, in, func(raw domain.MatchBookResult) MatchBookResult {
		return MatchBookResult{
			Book:   convertBook(raw.Book),
			Result: convertMatchRecord(raw.Result),
		}
	})

	return out
}
