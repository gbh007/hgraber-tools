package domain

import "time"

type Book struct {
	ID        int
	Name      string
	Url       string
	PageCount int
	CreateAt  time.Time

	Pages []Page
}

type Page struct {
	PageNumber int
	Ext        string
	Url        string
	LoadAt     time.Time
}

func ConvertSlice[From any, To any](to []To, from []From, conv func(From) To) {
	for i, v := range from {
		if i >= len(to) {
			return
		}

		to[i] = conv(v)
	}
}
