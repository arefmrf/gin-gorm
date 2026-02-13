package query

import (
	"gorm.io/gorm"
)

type Builder struct {
	db         *gorm.DB
	params     map[string]any
	order      []string
	joinsAdded map[string]bool
}

func New(db *gorm.DB, params map[string]any) *Builder {
	return &Builder{
		db:         db,
		params:     params,
		order:      []string{},
		joinsAdded: make(map[string]bool),
	}
}

func (b *Builder) addJoin(name, clause string, args ...interface{}) {
	if !b.joinsAdded[name] {
		b.db = b.db.Joins(clause, args...)
		b.joinsAdded[name] = true
	}
}
