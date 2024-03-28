package zipscaner

import (
	"app/internal/domain"
	"archive/zip"
	"crypto/md5"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func (Provider) ScanFilePages(path string) ([]domain.Page, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	stats, err := f.Stat()
	if err != nil {
		return nil, err
	}

	zipReader, err := zip.NewReader(f, stats.Size())
	if err != nil {
		return nil, err
	}

	res := make([]domain.Page, 0, len(zipReader.File))

	for _, file := range zipReader.File {
		number, _ := strconv.Atoi(strings.Split(file.Name, ".")[0])
		if number < 1 {
			continue
		}

		r, err := file.Open()
		if err != nil {
			return nil, err
		}

		data, err := io.ReadAll(r)
		r.Close()
		if err != nil {
			return nil, err
		}

		hash := fmt.Sprintf("%x", md5.Sum(data))

		res = append(res, domain.Page{
			Number: number,
			Size:   int64(file.UncompressedSize64),
			Hash:   hash,
		})
	}

	return res, nil
}
