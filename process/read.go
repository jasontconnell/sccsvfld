package process

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/jasontconnell/sitecore/api"
	"github.com/jasontconnell/sitecore/data"
)

func Read(connstr string, fieldIds []uuid.UUID) (data.ItemMap, error) {
	items, err := api.LoadItems(connstr)
	if err != nil {
		return nil, fmt.Errorf("loading items %w", err)
	}

	_, m := api.LoadItemMap(items)

	fv, err := api.LoadFilteredFieldValues(connstr, fieldIds, 20)

	api.AssignFieldValues(m, fv)

	return m, err
}
