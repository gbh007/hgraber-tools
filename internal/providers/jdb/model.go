package jdb

import "time"

type DatabaseData struct {
	Data map[int]RawTitle `json:"titles,omitempty"`
}

type RawPage struct {
	URL      string    `json:"url"`
	Ext      string    `json:"ext"`
	Success  bool      `json:"success"`
	LoadedAt time.Time `json:"loaded_at"`
	Rate     int       `json:"rate,omitempty"`
}

type RawTitleInfoParsed struct {
	Name       bool `json:"name,omitempty"`
	Page       bool `json:"page,omitempty"`
	Tags       bool `json:"tags,omitempty"`
	Authors    bool `json:"authors,omitempty"`
	Characters bool `json:"characters,omitempty"`
	Languages  bool `json:"languages,omitempty"`
	Categories bool `json:"categories,omitempty"`
	Parodies   bool `json:"parodies,omitempty"`
	Groups     bool `json:"groups,omitempty"`
}

type RawTitle struct {
	ID      int       `json:"id"`
	Created time.Time `json:"created"`
	URL     string    `json:"url"`

	Pages []RawPage    `json:"pages"`
	Data  RawTitleInfo `json:"info"`
}

type RawTitleInfo struct {
	Parsed     RawTitleInfoParsed `json:"parsed,omitempty"`
	Name       string             `json:"name,omitempty"`
	Rate       int                `json:"rate,omitempty"`
	Tags       []string           `json:"tags,omitempty"`
	Authors    []string           `json:"authors,omitempty"`
	Characters []string           `json:"characters,omitempty"`
	Languages  []string           `json:"languages,omitempty"`
	Categories []string           `json:"categories,omitempty"`
	Parodies   []string           `json:"parodies,omitempty"`
	Groups     []string           `json:"groups,omitempty"`
}
