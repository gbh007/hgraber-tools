package export

import "time"

type Book struct {
	ID        int       `json:"id"`
	Name      string    `json:"name,omitempty"`
	Url       string    `json:"url,omitempty"`
	PageCount int       `json:"page_count,omitempty"`
	CreateAt  time.Time `json:"create_at,omitempty"`
	Pages     []Page    `json:"pages,omitempty"`
}

type Page struct {
	Number int       `json:"number"`
	Ext    string    `json:"ext"`
	Url    string    `json:"url"`
	LoadAt time.Time `json:"load_at,omitempty"`
}
