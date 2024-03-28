package main

import (
	"app/internal/providers/jsondata"
	"app/internal/providers/zipscaner"
	"app/internal/usecases/matcher"
	"context"
	"flag"
	"log"
)

func main() {
	fromSource := flag.String("f", "", "Данные для поиска")
	source := flag.String("src", "", "Источник данных БД (дамп exporter)")
	toFile := flag.String("to", "out.json", "Файл для сохранения результата")
	matchRate := flag.Float64("rate", 0.3, "Минимальный рейтинг совпадения")
	applyFilter := flag.Bool("filter", false, "Применить фильтр для сокращения данных")
	flag.Parse()

	ctx := context.Background()
	zipScanner := zipscaner.New()

	exporter := jsondata.New("")

	data, err := exporter.Import(ctx, *source)
	if err != nil {
		log.Fatalln(err)
	}

	books, err := exporter.Import(ctx, *fromSource)
	if err != nil {
		log.Fatalln(err)
	}

	matchResult, err := matcher.New(zipScanner, data).MatchBooks(books, *matchRate)
	if err != nil {
		log.Fatalln(err)
	}

	if *applyFilter {
		matchResult.ApplyFilter()
	}

	err = exporter.ExportMatchBooksResult(ctx, *toFile, matchResult)
	if err != nil {
		log.Fatalln(err)
	}
}
