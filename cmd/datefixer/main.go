package main

import (
	"app/internal/providers/export"
	"app/internal/providers/postgresql"
	"context"
	"flag"
	"log"
)

func main() {
	fromSource := flag.String("f", "", "Источник данных")
	toSource := flag.String("to", "", "Источник для обновления")
	flag.Parse()

	ctx := context.Background()

	exporter := export.New("")

	data, err := exporter.Import(ctx, *fromSource)
	if err != nil {
		log.Fatalln(err)
	}

	db, err := postgresql.Connect(ctx, *toSource)
	if err != nil {
		log.Fatalln(err)
	}

	err = db.UpdateDates(ctx, data)
	if err != nil {
		log.Fatalln(err)
	}
}
