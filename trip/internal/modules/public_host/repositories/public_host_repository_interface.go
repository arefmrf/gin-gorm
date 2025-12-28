package repositories

import (
	HostModels "trip/internal/persistence/host/model"
	"trip/pkg/pagination"
)

type PublicHostRepositoryInterface interface {
	List(pagination pagination.Pagination) ([]HostModels.Host, int64)
}
