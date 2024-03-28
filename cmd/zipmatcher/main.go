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
	fromSource := flag.String("f", "", "Папка для поиска")
	source := flag.String("src", "", "Источник данных БД (дамп exporter)")
	toFile := flag.String("to", "out.json", "Файл для сохранения результата")
	matchRate := flag.Float64("rate", 0.3, "Минимальный рейтинг совпадения")
	recursive := flag.Bool("r", false, "Рекурсивно обходить папки")
	flag.Parse()

	ctx := context.Background()
	zipScanner := zipscaner.New()

	exporter := jsondata.New("")

	data, err := exporter.Import(ctx, *source)
	if err != nil {
		log.Fatalln(err)
	}

	matchResult, err := matcher.New(zipScanner, data).MatchFiles(*fromSource, *recursive, *matchRate)
	if err != nil {
		log.Fatalln(err)
	}

	err = exporter.ExportMatchFilesResult(ctx, *toFile, matchResult)
	if err != nil {
		log.Fatalln(err)
	}
}
