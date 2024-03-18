package helper

import "fmt"

func IntToCurrency(currency uint64) string {
	format := "Rp%d"
	if currency < 0 {
		format = "-Rp%d"
	}
	return fmt.Sprintf(format, currency)
}
