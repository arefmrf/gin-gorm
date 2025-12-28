package query

import (
	"context"

	"trip/internal/persistence/host/model"
	"trip/internal/services/public_host/responses"
	"trip/pkg/database"
)

func FetchPublicHosts(ctx context.Context) ([]responses.HostListItem, error) {
	var hosts []responses.HostListItem

	err := database.Connection().
		WithContext(ctx).
		Model(&model.Host{}).
		Select("uid, title, rate, rate_count, priority").
		//Where("status = ?", model.StatusApproved).
		Where("hidden = false").
		//Order("priority DESC").
		Find(&hosts).Error

	return hosts, err
}
