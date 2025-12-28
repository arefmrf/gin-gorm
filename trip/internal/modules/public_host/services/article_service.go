package services

import (
	PublicHostRepository "trip/internal/modules/public_host/repositories"
	PublicHostResponse "trip/internal/modules/public_host/responses"
	defaultPagination "trip/pkg/pagination"
)

type PublicHostService struct {
	publicHostRepository PublicHostRepository.PublicHostRepositoryInterface
}

func New() *PublicHostService {
	return &PublicHostService{
		publicHostRepository: PublicHostRepository.New(),
	}
}

func (p *PublicHostService) GetPublicHosts(pagination defaultPagination.Pagination) defaultPagination.Result[PublicHostResponse.PublicHosts] {
	hosts, total := p.publicHostRepository.List(pagination)
	//return PublicHostResponse.ToPublicHosts(hosts)
	return defaultPagination.NewResult(
		PublicHostResponse.ToPublicHosts(hosts),
		pagination,
		total,
	)
}
