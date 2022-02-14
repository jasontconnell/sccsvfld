package main

import (
	"flag"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/jasontconnell/sccsvfld/conf"
	"github.com/jasontconnell/sccsvfld/process"
	"github.com/jasontconnell/sitecore/data"
)

func main() {
	start := time.Now()
	fn := flag.String("c", "config.json", "config filename")
	flag.Parse()

	cfg := conf.LoadConfig(*fn)

	destId, err := uuid.Parse(cfg.DestFieldId)
	if err != nil {
		log.Fatal(err)
	}

	srcId, err := uuid.Parse(cfg.SrcFieldId)
	if err != nil {
		log.Fatal(err)
	}

	items, err := process.Read(cfg.ConnectionString, []uuid.UUID{destId, srcId})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("read", len(items), "items.", time.Since(start))

	updFields, err := process.AssignTags(items, srcId, destId, data.English, data.FieldSource(cfg.DestTable))
	if err != nil {
		log.Fatal(err)
	}
	log.Println("updating", len(updFields), "fields.", time.Since(start))

	err = process.Update(cfg.ConnectionString, updFields)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("finished. ", time.Since(start))
}
