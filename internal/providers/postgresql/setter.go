package postgresql

import (
	"app/internal/domain"
	"context"
	"database/sql"
	"log"
	"time"
)

func (db *Database) UpdateDates(ctx context.Context, data []domain.Book) error {
	for _, book := range data {
		if !book.CreateAt.IsZero() {
			err := db.updateBookDate(ctx, book.ID, book.CreateAt)
			if err != nil {
				return err
			}
		}

		for _, page := range book.Pages {
			if !page.LoadAt.IsZero() {
				err := db.updatePageDate(ctx, book.ID, page.Number, page.LoadAt)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (db *Database) updatePageDate(ctx context.Context, id int, page int, t time.Time) error {
	res, err := db.db.ExecContext(
		ctx,
		`UPDATE pages SET load_at = $3 WHERE book_id = $1 AND page_number = $2;`,
		id, page, t.UTC(),
	)
	if err != nil {
		return err
	}

	ok, _ := isApplyWithErr(res)
	if !ok {
		log.Printf("not updated %d %d\n", id, page)
	}

	return nil
}

func isApplyWithErr(r sql.Result) (bool, error) {
	c, err := r.RowsAffected()
	if err != nil {
		return false, nil
	}

	return c != 0, nil
}

func (db *Database) updateBookDate(ctx context.Context, id int, t time.Time) error {
	res, err := db.db.ExecContext(
		ctx,
		`UPDATE books SET create_at = $1 WHERE id = $2;`,
		t.UTC(), id,
	)
	if err != nil {
		return err
	}

	ok, _ := isApplyWithErr(res)
	if !ok {
		log.Printf("not updated %d\n", id)
	}

	return nil
}
