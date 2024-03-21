package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/nozzlium/heymat_backend/response"
)

func TestMake(t *testing.T) {
	wow := make([]response.MonthlyBalance, 12, 12)
	fmt.Println(wow)
	fmt.Println(time.Now().String())
}
