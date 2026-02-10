package query

import (
	"trip/pkg/utils"

	"gorm.io/gorm"
)

func (b *Builder) FilterLocation() *gorm.DB {
	locVal, ok := b.params["location"]
	if !ok {
		return b.db
	}

	locations := utils.NormalizeList(locVal)
	if len(locations) == 0 {
		return b.db
	}
	b.addJoin("places", "LEFT JOIN places ON places.id = hosts.place_id")

	return b.db.
		//Where(
		//	b.db.Where("places.city IN ?", locations).
		//		Or("places.country IN ?", locations).
		//		Or("places.province IN ?", locations),
		//)
		Where("places.city IN ? OR places.country IN ? OR places.province IN ?", locations, locations, locations)
}
