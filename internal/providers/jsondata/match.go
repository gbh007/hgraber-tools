package jsondata

import (
	"app/internal/domain"
	"context"
	"encoding/json"
	"os"
	"path"
)

func (ep ExportProvider) ExportPageMatches(ctx context.Context, filename string, data []domain.MatchRecord) error {
	finalPath := path.Join(ep.basePath, filename)

	dataToExport := convertMatchRecord(data)

	f, err := os.Create(finalPath)
	if err != nil {
		return err
	}
	defer f.Close()

	encoder := json.NewEncoder(f)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(dataToExport)
	if err != nil {
		return err
	}

	return nil
}

func (ep ExportProvider) ExportMatchFilesResult(ctx context.Context, filename string, data domain.MatchFilesResult) error {
	finalPath := path.Join(ep.basePath, filename)

	dataToExport := MatchFilesResult{
		NotMatched: data.NotMatched,
		Matched:    convertMatchFileResult(data.Matched),
	}

	f, err := os.Create(finalPath)
	if err != nil {
		return err
	}
	defer f.Close()

	encoder := json.NewEncoder(f)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(dataToExport)
	if err != nil {
		return err
	}

	return nil
}

func (ep ExportProvider) ExportMatchBooksResult(ctx context.Context, filename string, data domain.MatchBooksResult) error {
	finalPath := path.Join(ep.basePath, filename)

	dataToExport := MatchBooksResult{
		NotMatched: convertBooks(data.NotMatched),
		Matched:    convertMatchBookResult(data.Matched),
	}

	f, err := os.Create(finalPath)
	if err != nil {
		return err
	}
	defer f.Close()

	encoder := json.NewEncoder(f)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(dataToExport)
	if err != nil {
		return err
	}

	return nil
}
