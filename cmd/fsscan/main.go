package main

import (
	"app/internal/providers/fsscan"
	"app/internal/providers/jsondata"
	"context"
	"flag"
	"log"
)

func main() {
	fromSource := flag.String("f", "", "Источник данных")
	toFile := flag.String("to", "out.json", "Файл для сохранения")
	flag.Parse()

	ctx := context.Background()

	db := fsscan.New(*fromSource)

	data, err := db.Books(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	exporter := jsondata.New("")

	err = exporter.Export(ctx, *toFile, data)
	if err != nil {
		log.Fatalln(err)
	}
}
