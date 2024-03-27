package export

import "time"

type Book struct {
	ID         int        `json:"id"`
	Name       string     `json:"name,omitempty"`
	Url        string     `json:"url,omitempty"`
	CreateAt   time.Time  `json:"create_at,omitempty"`
	Attributes Attributes `json:"attributes,omitempty"`
	PageCount  int        `json:"page_count,omitempty"`
	Pages      []Page     `json:"pages,omitempty"`
}

type Page struct {
	Number int       `json:"number"`
	Ext    string    `json:"ext"`
	Url    string    `json:"url"`
	LoadAt time.Time `json:"load_at,omitempty"`

	Hash string `json:"hash,omitempty"`
	Size int64  `json:"size,omitempty"`
}

type Attributes struct {
	Tags       []string `json:"tags,omitempty"`
	Authors    []string `json:"authors,omitempty"`
	Characters []string `json:"characters,omitempty"`
	Languages  []string `json:"languages,omitempty"`
	Categories []string `json:"categories,omitempty"`
	Parodies   []string `json:"parodies,omitempty"`
	Groups     []string `json:"groups,omitempty"`
}
