package jdb

import (
	"encoding/json"
	"os"
)

type Database struct {
	data DatabaseData
}

func New(path string) (*Database, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	dbData := DatabaseData{}

	err = json.NewDecoder(f).Decode(&dbData)
	if err != nil {
		return nil, err
	}

	return &Database{
		data: dbData,
	}, nil
}
