package merge

import (
	"app/internal/domain"
	"time"
)

type Merger struct {
	logDiff bool
}

func New(logDiff bool) Merger {
	return Merger{
		logDiff: logDiff,
	}
}

func (m Merger) Merge(a, b []domain.Book) []domain.Book {
	return m.fromMap(m.mergeBooks(m.toMap(a), m.toMap(b)))
}

func mergeDate(a, b time.Time) time.Time {
	if isZeroTime(a) {
		return b
	}

	if isZeroTime(b) {
		return a
	}

	if a.Before(b) {
		return a
	}

	return b
}
