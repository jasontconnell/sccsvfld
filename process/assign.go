package process

import (
	"log"
	"strings"

	"github.com/google/uuid"
	"github.com/jasontconnell/sitecore/data"
)

func AssignTags(m data.ItemMap, srcFieldId, destFieldId uuid.UUID, lang data.Language) ([]data.UpdateField, error) {
	updflds := []data.UpdateField{}
	for _, item := range m {
		fv := item.GetFieldValue(srcFieldId, lang)
		if fv == nil || len(strings.Trim(fv.GetValue(), " ")) == 0 {
			continue
		}

		ids := strings.Split(fv.GetValue(), "|")

		csvstr := ""
		for _, idstr := range ids {
			uid, err := uuid.Parse(idstr)
			if err != nil {
				log.Println("couldn't parse id from ", idstr, err)
			}

			refitem, ok := m[uid]
			if ok {
				csvstr += refitem.GetName() + ","
			}
		}

		csvstr = strings.TrimRight(csvstr, ",")

		uptype := data.Insert
		exists := item.GetFieldValue(destFieldId, lang) != nil
		if exists {
			uptype = data.Update
		}

		upd := data.UpdateField{ItemID: item.GetId(), FieldID: destFieldId, Value: csvstr, Source: data.VersionedFields, Version: fv.GetVersion(), Language: lang, UpdateType: uptype}
		updflds = append(updflds, upd)
	}

	return updflds, nil
}
