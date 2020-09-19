package core

import (
	"time"
)

type Timer struct {
	Ticker	*time.Ticker
	Time	float64
	Code	*TokenList
}