package helper

import (
	"fmt"
	"time"
)

var idMonthNames = map[time.Month]string{
	time.January:   "Januari",
	time.February:  "Februari",
	time.March:     "Maret",
	time.April:     "April",
	time.May:       "Mei",
	time.June:      "Juni",
	time.July:      "Juli",
	time.August:    "Agustus",
	time.September: "September",
	time.October:   "Oktober",
	time.November:  "November",
	time.December:  "Desember",
}

var idDayNames = map[time.Weekday]string{
	time.Monday:    "Senin",
	time.Tuesday:   "Selasa",
	time.Wednesday: "Rabu",
	time.Thursday:  "Kamis",
	time.Friday:    "Jum'at",
	time.Saturday:  "Sabtu",
	time.Sunday:    "Minggu",
}

func GetIdDateStringFull(tme time.Time) string {
	return fmt.Sprintf(
		"%s, %d %s %d",
		idDayNames[tme.Weekday()],
		tme.Day(),
		idMonthNames[tme.Month()],
		tme.Year(),
	)
}

func GetIdDateStringMonth(tme time.Time) string {
	return fmt.Sprintf("%s %d", idMonthNames[tme.Month()], tme.Year())
}
