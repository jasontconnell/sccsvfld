package process

import (
	"log"

	"github.com/jasontconnell/sitecore/api"
	"github.com/jasontconnell/sitecore/data"
)

func Update(connstr string, flds []data.UpdateField, tbl data.FieldSource) error {
	c, err := api.Update(connstr, nil, flds)

	log.Println("updated", c, "items")
	return err
}
