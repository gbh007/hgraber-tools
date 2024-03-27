package export

import (
	"app/internal/domain"
	"context"
	"encoding/json"
	"os"
	"path"
)

func (ep ExportProvider) Import(ctx context.Context, filename string) ([]domain.Book, error) {
	finalPath := path.Join(ep.basePath, filename)

	f, err := os.Open(finalPath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	dataToImport := make([]Book, 0)
	err = json.NewDecoder(f).Decode(&dataToImport)
	if err != nil {
		return nil, err
	}

	data := make([]domain.Book, len(dataToImport))

	for i, raw := range dataToImport {
		pages := make([]domain.Page, len(raw.Pages))
		domain.ConvertSlice(pages, raw.Pages, func(p Page) domain.Page {
			return domain.Page{
				PageNumber: p.Number,
				Ext:        p.Ext,
				Url:        p.Url,
				LoadAt:     p.LoadAt,
			}
		})

		data[i] = domain.Book{
			ID:        raw.ID,
			Name:      raw.Name,
			Url:       raw.Url,
			PageCount: raw.PageCount,
			CreateAt:  raw.CreateAt,
			Pages:     pages,
		}
	}

	return data, nil
}
