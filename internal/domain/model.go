package domain

import "time"

type Book struct {
	ID        int
	Name      string
	Url       string
	PageCount int
	CreateAt  time.Time

	Pages []Page

	Attributes Attributes
}

type Page struct {
	Number int
	Ext    string
	Url    string
	LoadAt time.Time

	Hash string
	Size int64
}

type Attributes struct {
	Tags       []string
	Authors    []string
	Characters []string
	Languages  []string
	Categories []string
	Parodies   []string
	Groups     []string
}

func ConvertSlice[From any, To any](to []To, from []From, conv func(From) To) {
	for i, v := range from {
		if i >= len(to) {
			return
		}

		to[i] = conv(v)
	}
}
