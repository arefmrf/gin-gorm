package repositories

import (
	"gorm.io/gorm"
	HostModels "trip/internal/persistence/host/model"
	"trip/pkg/database"
	"trip/pkg/pagination"
)

type PublicHostRepository struct {
	DB *gorm.DB
}

func New() *PublicHostRepository {
	return &PublicHostRepository{
		DB: database.Connection(),
	}
}

func (r *PublicHostRepository) List(pagination pagination.Pagination) ([]HostModels.Host, int64) {
	var hosts []HostModels.Host
	var total int64
	r.DB.Model(&HostModels.Host{}).Count(&total)
	//articleRepository.DB.Limit(limit).Joins("User").Find(&articles)
	r.DB.
		Limit(pagination.Limit).
		Offset(pagination.Offset()).
		Find(&hosts)

	return hosts, total
}
