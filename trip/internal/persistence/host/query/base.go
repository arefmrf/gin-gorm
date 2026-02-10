package query

import (
	HostModels "trip/internal/persistence/host/model"

	"gorm.io/gorm"
)

func (b *Builder) Base() *gorm.DB {
	return b.db.
		Model(&HostModels.Host{}).
		Where("hidden = ?", false).
		Where("status = ?", HostModels.StatusApproved).
		Where("out_of_service = ?", false)
}
