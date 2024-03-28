package domain

import "slices"

type MatchRecord struct {
	ID        int
	Name      string
	Url       string
	MatchRate float64 // [0;1]
}

type MatchFileResult struct {
	Path   string
	Result []MatchRecord
}

type MatchFilesResult struct {
	NotMatched []string
	Matched    []MatchFileResult
}

type MatchBookResult struct {
	Book   Book
	Result []MatchRecord
}

type MatchBooksResult struct {
	NotMatched []Book
	Matched    []MatchBookResult
}

func (r *MatchBooksResult) ApplyFilter() {
	ConvertSlice(r.NotMatched, r.NotMatched, func(in Book) Book {
		in.Pages = nil
		in.Attributes = Attributes{}
		return in
	})
	ConvertSlice(r.Matched, r.Matched, func(in MatchBookResult) MatchBookResult {
		in.Book.Pages = nil
		in.Book.Attributes = Attributes{}
		return in
	})

	r.Matched = slices.DeleteFunc(r.Matched, func(e MatchBookResult) bool {
		return len(e.Result) < 2
	})
}
