package helper

import (
	"fmt"
	"math"
	"strconv"
)

func IntToCurrency(currency int64) string {
	format := "Rp%s"
	if currency < 0 {
		format = "-Rp%s"
	}
	return fmt.Sprintf(format, GroupDecimals2(currency))
}

func GroupDecimals(currency int64) string {
	res := ""
	count := 0
	for currency > 0 {
		one := currency % 10
		res = strconv.Itoa(int(one)) + res
		currency = currency / 10
		if count++; count == 3 {
			res = "." + res
			count = 0
		}
	}
	return res
}

func GroupDecimals2(currency int64) string {
	currency = int64(math.Abs(float64(currency)))
	bytes := make([]byte, 0, 10)
	count := 0
	for currency > 0 {
		one := currency % 10
		bytes = append(bytes, byte(one+48))
		currency = currency / 10
		if count++; count == 3 {
			bytes = append(bytes, byte(rune('.')))
			count = 0
		}
	}
	if len(bytes) == 0 {
		return "0"
	}
	res := make([]byte, 0, len(bytes))
	for i := len(bytes) - 1; i >= 0; i-- {
		res = append(res, bytes[i])
	}
	return string(res)
}
