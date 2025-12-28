package services

import (
	PublicHostResponse "trip/internal/modules/public_host/responses"
	"trip/pkg/pagination"
)

type ArticleServiceInterface interface {
	GetPublicHosts(pagination pagination.Pagination) pagination.Result[PublicHostResponse.PublicHosts]
}
