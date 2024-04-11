package lib

import (
	"os"
	"strconv"
)

func GetStringEnv(
	varKey string,
	target *string,
) {
	varVal, ok := os.LookupEnv(varKey)
	if !ok {
		return
	}
	*target = varVal
}

func GetIntEnv(
	varKey string,
	target *int,
) {
	var varValStr string
	GetStringEnv(varKey, &varValStr)
	varValInt, err := strconv.Atoi(
		varValStr,
	)
	if err != nil {
		return
	}
	*target = varValInt
}
