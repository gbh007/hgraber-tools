package zipscaner

import (
	"os"
	"path"
)

func (p Provider) ScanFiles(dirPath string, recursive bool) ([]string, error) {
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		return nil, err
	}

	res := make([]string, 0)
	for _, e := range entries {
		name := e.Name()
		fullPath := path.Join(dirPath, name)

		if e.IsDir() {
			if recursive {
				subData, err := p.ScanFiles(fullPath, recursive)
				if err != nil {
					return nil, err
				}

				res = append(res, subData...)
			}

			continue
		}

		if path.Ext(name) != ".zip" {
			continue
		}

		res = append(res, fullPath)
	}

	return res, nil
}
