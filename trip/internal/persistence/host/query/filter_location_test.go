package query

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"

	"trip/internal/persistence/host/model"
	"trip/pkg/testutils"
)

func ptr(s string) *string {
	return &s
}

func seedHostsLocation(db *gorm.DB) {
	db.Create(&model.Host{
		UID:        NewUID(),
		Identifier: "luxury-hotel",
		Title:      "Luxury Hotel",
		Place: &model.Place{
			UID:     NewUID(),
			City:    ptr("karaj"),
			Country: ptr("iran"),
			Info: map[string]any{
				"city": map[string]string{
					"fa": "کرج",
					"en": "karaj",
				},
			},
		},
	})

	db.Create(&model.Host{
		UID:        NewUID(),
		Title:      "Cheap Hostel",
		Identifier: "Cheap Hostel",
		Place: &model.Place{
			UID:     NewUID(),
			City:    ptr("tehran"),
			Country: ptr("iran"),
			Info: map[string]any{
				"city": map[string]string{
					"fa": "تهران",
					"en": "tehran",
				},
			},
		},
	})
}

func TestFilterLocation(t *testing.T) {
	db := testutils.SetupTestDB(
		t,
		&model.Place{},
		&model.Host{},
	)

	seedHostsLocation(db)
	builder := New(db, map[string]any{"location": "karaj"})
	//filter := builder.FilterSearch()

	var hosts []model.Host

	query := builder.
		FilterLocation().
		Preload("Place")
	err := query.Find(&hosts).Error

	assert.NoError(t, err)
	assert.Len(t, hosts, 1)
	assert.Equal(t, "Luxury Hotel", hosts[0].Title)
}
