package query

import (
	"math/rand"
	"testing"
	"time"

	"github.com/oklog/ulid/v2"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"

	"trip/internal/persistence/host/model"
	"trip/pkg/testutils"
)

func NewUID() string {
	entropy := ulid.Monotonic(rand.New(rand.NewSource(time.Now().UnixNano())), 0)
	return ulid.MustNew(ulid.Timestamp(time.Now()), entropy).String()
}

func seedHosts(db *gorm.DB) {
	db.Create(&model.Host{
		UID:        NewUID(),
		Identifier: "luxury-hotel",
		Title:      "Luxury Hotel",
		Place: &model.Place{
			UID: NewUID(),
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
			UID: NewUID(),
			Info: map[string]any{
				"city": map[string]string{
					"fa": "تهران",
					"en": "tehran",
				},
			},
		},
	})
}

func TestFilterSearch_ByCityJSONB(t *testing.T) {
	db := testutils.SetupTestDB(
		t,
		&model.Place{},
		&model.Host{},
	)

	seedHosts(db)
	builder := New(db, map[string]any{"search": "کرج"})
	//filter := builder.FilterSearch()

	var hosts []model.Host

	query := builder.
		FilterSearch().
		Preload("Place").
		Debug()
	err := query.Find(&hosts).Error

	assert.NoError(t, err)
	assert.Len(t, hosts, 1)
	assert.Equal(t, "Luxury Hotel", hosts[0].Title)
}
