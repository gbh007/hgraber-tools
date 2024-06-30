package main

import (
	"app/internal/providers/hg4client"
	"app/internal/providers/hg5client"
	"context"
	"flag"
	"log"
)

func main() {
	hg4Addr := flag.String("hg4-addr", "", "")
	hg4Token := flag.String("hg4-token", "", "")
	hg5Addr := flag.String("hg5-addr", "", "")
	hg5Token := flag.String("hg5-token", "", "")
	from := flag.Int("from", 0, "")
	to := flag.Int("to", 0, "")
	flag.Parse()

	ctx := context.Background()

	hg4 := hg4client.New(*hg4Addr, *hg4Token)
	hg5 := hg5client.New(*hg5Addr, *hg5Token)

	for i := *from; i <= *to; i++ {
		body, err := hg4.BookArchive(ctx, i)
		if err != nil {
			log.Fatalln(i, err)
		}

		err = hg5.ImportArchive(ctx, body)
		if err != nil {
			log.Fatalln(i, err)
		}

		if i%100 == 0 {
			log.Println(i, "finished")
		}
	}
}
