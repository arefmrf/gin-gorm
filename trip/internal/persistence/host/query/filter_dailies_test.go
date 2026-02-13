package query

import (
	"fmt"
	"testing"
	"time"
	"trip/internal/persistence"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"

	"trip/internal/persistence/host/model"
	"trip/pkg/testutils"
)

func GenerateDailies(hostID uint, start time.Time, nights int) []model.HostDaily {
	var dailies []model.HostDaily

	start = time.Date(
		start.Year(), start.Month(), start.Day(),
		12, 0, 0, 0,
		time.UTC,
	)

	for i := 0; i < nights; i++ {
		day := start.AddDate(0, 0, i)

		dailies = append(dailies, model.HostDaily{
			UID:           fmt.Sprintf("HD-%d-%d", hostID, i),
			HostID:        hostID,
			Date:          day.Unix(),
			Price:         100,
			Discount:      0,
			TotalCapacity: 20,
			MaleBooked:    10,
			FemaleBooked:  10,
			Capacity: persistence.JSONB{
				"male":         10,
				"female":       10,
				"base_person":  0,
				"extra_person": 0,
			},
		})
	}

	return dailies
}

func getDate(timeDelta *int) time.Time {
	now := time.Now()

	date := time.Date(
		now.Year(),
		now.Month(),
		now.Day(),
		12, 0, 0, 0,
		time.UTC,
	)
	if timeDelta != nil {
		date = date.AddDate(0, 0, *timeDelta)
	}
	return date
}

func seedHostsDailies(db *gorm.DB) {
	sharedHostType := model.HostType{
		UID:      NewUID(),
		Category: 1,
		Title:    "Test",
	}
	if err := db.Create(&sharedHostType).Error; err != nil {
		panic(err)
	}
	host := model.Host{
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
		HostTypeID: &sharedHostType.ID,
	}
	if err := db.Create(&host).Error; err != nil {
		panic(err)
	}

	start := time.Now().UTC()
	dailies := GenerateDailies(host.ID, start, 30)
	if err := db.Create(&dailies).Error; err != nil {
		panic(err)
	}

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

func TestFilterDailies(t *testing.T) {
	db := testutils.SetupTestDB(
		t,
		&model.Place{},
		&model.Host{},
		&model.HostDaily{},
		&model.HostType{},
	)

	startDate := getDate(nil)
	timeDelta := 1
	endDate := getDate(&timeDelta)

	seedHostsDailies(db)
	builder := New(db, map[string]any{
		"start_date": startDate.Unix(),
		"end_date":   endDate.Unix(),
	})

	var hosts []model.Host

	query := builder.
		FilterDailies().
		Preload("Place").Debug()
	err := query.Find(&hosts).Error

	assert.NoError(t, err)
	assert.Len(t, hosts, 1)
	assert.Equal(t, "Luxury Hotel", hosts[0].Title)
}
