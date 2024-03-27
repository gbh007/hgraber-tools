package postgresql

import (
	"app/internal/domain"
	"context"
)

func (db *Database) getPages(ctx context.Context, bookID int) ([]domain.Page, error) {
	rawData := make([]Page, 0)
	err := db.db.SelectContext(ctx, &rawData, `SELECT * FROM pages WHERE book_id = $1 ORDER BY page_number;`, bookID)
	if err != nil {
		return nil, err
	}

	res := make([]domain.Page, len(rawData))
	domain.ConvertSlice(res, rawData, pageToDomain)

	return res, nil
}

func (db *Database) Books(ctx context.Context) ([]domain.Book, error) {
	rawData := make([]Book, 0)
	err := db.db.SelectContext(ctx, &rawData, `SELECT * FROM books;`)
	if err != nil {
		return nil, err
	}

	res := make([]domain.Book, len(rawData))
	for i, raw := range rawData {
		pages, err := db.getPages(ctx, raw.ID)
		if err != nil {
			return nil, err
		}

		attributes, err := db.getBookAttributes(ctx, raw.ID)
		if err != nil {
			return nil, err
		}

		b := bookToDomain(raw)

		b.Pages = pages
		b.Attributes = attributes

		res[i] = b
	}

	return res, nil
}

func (db *Database) getBookAttributes(ctx context.Context, bookID int) (domain.Attributes, error) {
	raw := make([]BookAttribute, 0)

	err := db.db.SelectContext(ctx, &raw, `SELECT * FROM book_attributes WHERE book_id = $1;`, bookID)
	if err != nil {
		return domain.Attributes{}, err
	}

	attributes := domain.Attributes{}

	for _, rawAttr := range raw {
		switch rawAttr.Attr {
		case "author":
			attributes.Authors = append(attributes.Authors, rawAttr.Value)
		case "category":
			attributes.Categories = append(attributes.Categories, rawAttr.Value)
		case "character":
			attributes.Characters = append(attributes.Characters, rawAttr.Value)
		case "group":
			attributes.Groups = append(attributes.Groups, rawAttr.Value)
		case "language":
			attributes.Languages = append(attributes.Languages, rawAttr.Value)
		case "parody":
			attributes.Parodies = append(attributes.Parodies, rawAttr.Value)
		case "tag":
			attributes.Tags = append(attributes.Tags, rawAttr.Value)
		}
	}

	return attributes, nil
}
