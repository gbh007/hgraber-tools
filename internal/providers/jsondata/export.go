package jsondata

import (
	"app/internal/domain"
	"context"
	"encoding/json"
	"os"
	"path"
)

type ExportProvider struct {
	basePath string
}

func New(path string) ExportProvider {
	return ExportProvider{
		basePath: path,
	}
}

func (ep ExportProvider) Export(ctx context.Context, filename string, data []domain.Book) error {
	finalPath := path.Join(ep.basePath, filename)

	dataToExport := convertBooks(data)

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
