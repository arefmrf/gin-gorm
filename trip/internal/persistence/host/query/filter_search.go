package query

import (
	"gorm.io/gorm"
)

func (b *Builder) FilterSearch() *gorm.DB {
	title, ok := b.params["search"].(string)
	if !ok || title == "" {
		return b.db
	}

	like := "%" + title + "%"

	b.addJoin("places", "LEFT JOIN places ON places.id = hosts.place_id")

	db := b.db.Where(`
		(
			hosts.title ILIKE ?
			OR hosts.identifier = ?
			OR EXISTS (
				SELECT 1
				FROM jsonb_each_text(places.info->'country') kv
				WHERE kv.value ILIKE ?
			)
			OR EXISTS (
				SELECT 1
				FROM jsonb_each_text(places.info->'province') kv
				WHERE kv.value ILIKE ?
			)
			OR EXISTS (
				SELECT 1
				FROM jsonb_each_text(places.info->'city') kv
				WHERE kv.value ILIKE ?
			)
		)
	`,
		like,
		title,
		like,
		like,
		like,
	)

	arabic := ArabicNormalize(title)
	if arabic != title {
		db = db.Or("hosts.title ILIKE ?", "%"+arabic+"%")
	}

	return db
}

func ArabicNormalize(input string) string {
	// TODO: implement real normalization
	return input
}
