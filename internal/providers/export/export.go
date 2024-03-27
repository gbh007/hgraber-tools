package export

import (
	"app/internal/domain"
	"context"
	"encoding/json"
	"os"
	"path"
	"slices"
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

	dataToExport := make([]Book, len(data))
	for i, raw := range data {
		pages := make([]Page, len(raw.Pages))
		domain.ConvertSlice(pages, raw.Pages, func(p domain.Page) Page {
			return Page{
				Number: p.Number,
				Ext:    p.Ext,
				Url:    p.Url,
				LoadAt: p.LoadAt,
				Hash:   p.Hash,
				Size:   p.Size,
			}
		})

		slices.SortFunc(pages, func(a, b Page) int {
			return a.Number - b.Number
		})

		dataToExport[i] = Book{
			ID:         raw.ID,
			Name:       raw.Name,
			Url:        raw.Url,
			PageCount:  raw.PageCount,
			CreateAt:   raw.CreateAt,
			Pages:      pages,
			Attributes: Attributes(raw.Attributes),
		}
	}

	slices.SortFunc(dataToExport, func(a, b Book) int {
		return a.ID - b.ID
	})

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
