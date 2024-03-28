package matcher

import (
	"app/internal/domain"
)

type zipScanner interface {
	ScanFilePages(path string) ([]domain.Page, error)
	ScanFiles(path string, recursive bool) ([]string, error)
}

type hashInfo struct {
	BookID     int
	PageNumber int
}

type Matcher struct { // Это не корректный юзкейс, т.к. у него есть состояние, но поскольку это прототип то плохой код допустим
	books  map[int]domain.Book
	hashes map[string][]hashInfo

	zipScanner zipScanner
}

func New(zipScanner zipScanner, rawData []domain.Book) *Matcher {
	books := make(map[int]domain.Book, len(rawData))
	hashes := make(map[string][]hashInfo, len(rawData)*30) // В среднем у книги по 30 страниц

	for _, book := range rawData {
		books[book.ID] = book
		for _, page := range book.Pages {
			hashes[page.Hash] = append(hashes[page.Hash], hashInfo{
				BookID:     book.ID,
				PageNumber: page.Number,
			})
		}
	}

	return &Matcher{
		books:      books,
		hashes:     hashes,
		zipScanner: zipScanner,
	}
}
