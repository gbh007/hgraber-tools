package main

import (
	"app/internal/domain"
	"app/internal/providers/jdb"
	"app/internal/providers/jsondata"
	"app/internal/providers/postgresql"
	"app/internal/providers/sqlite"
	"context"
	"flag"
	"log"
)

func main() {
	fromSource := flag.String("f", "", "Источник данных")
	fromType := flag.String("ft", "", "Тип источника данных")
	toFile := flag.String("to", "out.json", "Файл для сохранения")
	flag.Parse()

	ctx := context.Background()

	var data []domain.Book

	switch *fromType {
	case "jdb":
		db, err := jdb.New(*fromSource)
		if err != nil {
			log.Fatalln(err)
		}

		data, err = db.Books(ctx)
		if err != nil {
			log.Fatalln(err)
		}
	case "sqlite":
		db, err := sqlite.Connect(ctx, *fromSource)
		if err != nil {
			log.Fatalln(err)
		}

		data, err = db.Books(ctx)
		if err != nil {
			log.Fatalln(err)
		}
	case "postgresql":
		db, err := postgresql.Connect(ctx, *fromSource)
		if err != nil {
			log.Fatalln(err)
		}

		data, err = db.Books(ctx)
		if err != nil {
			log.Fatalln(err)
		}
	}

	exporter := jsondata.New("")

	err := exporter.Export(ctx, *toFile, data)
	if err != nil {
		log.Fatalln(err)
	}
}
