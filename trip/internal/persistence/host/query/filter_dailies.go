package query

import (
	"fmt"
	"strconv"
	"time"

	"gorm.io/gorm"
)

func noon(ts int64) time.Time {
	t := time.Unix(ts, 0).UTC()
	return time.Date(
		t.Year(), t.Month(), t.Day(),
		12, 0, 0, 0,
		time.UTC,
	)
}

func (b *Builder) FilterDailies() *gorm.DB {
	startRaw, ok1 := b.params["start_date"]
	endRaw, ok2 := b.params["end_date"]

	if !ok1 || !ok2 {
		return b.db
	}

	startTS, err1 := strconv.ParseInt(fmt.Sprint(startRaw), 10, 64)
	endTS, err2 := strconv.ParseInt(fmt.Sprint(endRaw), 10, 64)
	fmt.Println("======", startTS, "-----", endTS)
	if err1 != nil || err2 != nil {
		return b.db
	}

	checkIn := noon(startTS)
	checkOut := noon(endTS)
	totalNights := int(checkOut.Sub(checkIn).Hours() / 24)
	if totalNights <= 0 {
		return b.db
	}

	maleCap, _ := strconv.Atoi(fmt.Sprint(b.params["male_capacity"]))
	femaleCap, _ := strconv.Atoi(fmt.Sprint(b.params["female_capacity"]))
	capacity, _ := strconv.Atoi(fmt.Sprint(b.params["capacity"]))

	// required joins
	b.addJoin("host_daily", `
			JOIN host_daily hd
			  ON hd.host_id = hosts.id
			 -- AND hd.date >= ?
			 -- AND hd.date < ?
		`)
	//`, checkIn.Unix(), checkOut.Unix())

	b.addJoin("host_types", `
			JOIN host_types ht
			  ON ht.id = hosts.host_type_id
		`)

	// CASE 1: male/female capacity (ALL days must match)
	if maleCap > 0 || femaleCap > 0 {
		return b.db.
			Group("hosts.id").
			Having(`
					COUNT(
						CASE
							WHEN hd.date >= ? AND hd.date < ?
							AND (
								(ht.category = '1'
									AND (hd.capacity->>'male')::int - hd.male_booked >= ?
									AND (hd.capacity->>'female')::int - hd.female_booked >= ?
								)
								OR
								(ht.category = '2'
									AND hd.male_booked = 0
									AND hd.total_capacity >= ?
								)
							)
						THEN 1 END
					) = ?
	`,
				checkIn.Unix(), checkOut.Unix(),
				maleCap, femaleCap,
				maleCap+femaleCap,
				totalNights,
			)
	} else {
		// CASE 2: capacity only (AT LEAST ONE day)
		if capacity > 0 {
			return b.db.
				Group("hosts.id").
				Having(`
					COUNT(
						CASE
							WHEN hd.date >= ? AND hd.date < ?
							AND (
								(ht.category = '1'
									AND (hd.total_capacity - hd.male_booked - hd.female_booked) >= ?
								)
								OR
								(ht.category = '2'
									AND hd.male_booked = 0
									AND hd.total_capacity >= ?
								)
							)
						THEN 1 END
					) >= 1
		`, checkIn.Unix(), checkOut.Unix(), capacity, capacity)
		}
	}

	// CASE 3: only date availability (ALL days must exist)
	return b.db.
		Group("hosts.id").
		Having(`
			COUNT(
				CASE
					WHEN hd.date >= ? AND hd.date < ?
					THEN 1 END
				) = ?
		`, checkIn.Unix(), checkOut.Unix(), totalNights)
}
