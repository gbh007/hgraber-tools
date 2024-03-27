package fsscan

import (
	"app/internal/domain"
	"context"
	"fmt"
	"log"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

type Provider struct {
	path string
}

func New(path string) Provider {
	return Provider{
		path: path,
	}
}

func (p Provider) Books(_ context.Context) ([]domain.Book, error) {
	entries, err := os.ReadDir(p.path)
	if err != nil {
		return nil, err
	}

	res := make([]domain.Book, 0, len(entries))
	for _, e := range entries {
		if !e.IsDir() {
			continue
		}

		name := e.Name()
		id, err := strconv.Atoi(name)
		if err != nil {
			log.Println(err)
			continue
		}

		pages, err := p.scanPages(path.Join(p.path, name))
		if err != nil {
			return nil, err
		}

		bookCreateAt := time.Time{}

		for _, p := range pages {
			if bookCreateAt.IsZero() || (!p.LoadAt.IsZero() && p.LoadAt.Before(bookCreateAt)) {
				bookCreateAt = p.LoadAt
			}
		}

		res = append(res, domain.Book{
			ID:       id,
			Pages:    pages,
			CreateAt: bookCreateAt,
		})
	}

	return res, nil
}

func (Provider) scanPages(dirPath string) ([]domain.Page, error) {
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		return nil, err
	}

	res := make([]domain.Page, 0, len(entries))
	for _, e := range entries {
		if e.IsDir() {
			continue
		}

		name := e.Name()

		fileParts := strings.Split(name, ".")
		if len(fileParts) != 2 {
			return nil, fmt.Errorf("scan %s/%s : invalid file parts", dirPath, name)
		}

		number, err := strconv.Atoi(fileParts[0]) // Костыль на поиск первой части пути (заменить аналогом из path если есть)
		if err != nil {
			return nil, fmt.Errorf("scan %s/%s : %w", dirPath, name, err)
		}

		info, err := e.Info()
		if err != nil {
			return nil, err
		}

		res = append(res, domain.Page{
			Number: number,
			Ext:    fileParts[1],
			LoadAt: info.ModTime(),
		})
	}

	return res, nil
}
