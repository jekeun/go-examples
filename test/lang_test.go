package test

import (
	"fmt"
	"strings"
	"testing"
)

func TestArrayToString(t *testing.T) {

	// arr := [3]string{"A","B","C"}

	arr := []string{"KRW-BTC", "KRW-BCH", "KRW-XRP"}

	coinStr := strings.Join(arr, ",")

	fmt.Println(coinStr)
}

