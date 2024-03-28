package matcher

import "app/internal/domain"

func (m *Matcher) MatchFiles(path string, recursive bool, minMatchRate float64) (domain.MatchFilesResult, error) {
	res := domain.MatchFilesResult{}

	paths, err := m.zipScanner.ScanFiles(path, recursive)
	if err != nil {
		return domain.MatchFilesResult{}, err
	}

	for _, path := range paths {
		pages, err := m.zipScanner.ScanFilePages(path)
		if err != nil {
			return domain.MatchFilesResult{}, err
		}

		if len(pages) == 0 {
			res.NotMatched = append(res.NotMatched, path)
			continue
		}

		records := m.MatchPages(pages, minMatchRate)
		if len(records) == 0 {
			res.NotMatched = append(res.NotMatched, path)
			continue
		}

		res.Matched = append(res.Matched, domain.MatchFileResult{
			Path:   path,
			Result: records,
		})
	}

	return res, nil
}
