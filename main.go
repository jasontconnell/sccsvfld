package main

import (
	"flag"
	"log"

	"github.com/google/uuid"
	"github.com/jasontconnell/lpgatags/conf"
	"github.com/jasontconnell/lpgatags/process"
	"github.com/jasontconnell/sitecore/data"
)

func main() {
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

	updFields, err := process.AssignTags(items, srcId, destId, data.English)
	if err != nil {
		log.Fatal(err)
	}

	err = process.Update(cfg.ConnectionString, updFields, data.FieldSource(cfg.DestTable))
	if err != nil {
		log.Fatal(err)
	}
}
