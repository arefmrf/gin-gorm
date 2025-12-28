package responses

import (
	HostModels "trip/internal/persistence/host/model"
	"trip/pkg/mathutils"
)

type HostListItem struct {
	UID       string  `json:"uid"`
	Title     string  `json:"title"`
	Rate      float64 `json:"rate"`
	RateCount int     `json:"rate_count"`
	Priority  int     `json:"priority"`
}

type PublicHosts struct {
	Data []HostListItem
}

func toPublicHost(host HostModels.Host) HostListItem {
	return HostListItem{
		UID:       host.UID,
		Title:     host.Title,
		Rate:      mathutils.RoundFloat(host.Rate, 2),
		RateCount: int(host.RateCount),
		Priority:  int(host.Priority),
	}
}

func ToPublicHosts(articles []HostModels.Host) PublicHosts {
	var respons PublicHosts
	for _, article := range articles {
		respons.Data = append(respons.Data, toPublicHost(article))
	}
	return respons
}
