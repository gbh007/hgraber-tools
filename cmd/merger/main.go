package main

import (
	"app/internal/providers/export"
	"app/internal/usecases/merge"
	"context"
	"flag"
	"log"
)

func main() {
	aSource := flag.String("a", "", "Источник данных A")
	bSource := flag.String("b", "", "Источник данных B")
	toFile := flag.String("to", "out.json", "Файл для сохранения")
	logDiff := flag.Bool("diff", false, "Показать различие A и B")
	flag.Parse()

	ctx := context.Background()

	exporter := export.New("")

	aData, err := exporter.Import(ctx, *aSource)
	if err != nil {
		log.Fatalln(err)
	}

	bData, err := exporter.Import(ctx, *bSource)
	if err != nil {
		log.Fatalln(err)
	}

	data := merge.New(*logDiff).Merge(aData, bData)

	err = exporter.Export(ctx, *toFile, data)
	if err != nil {
		log.Fatalln(err)
	}
}
