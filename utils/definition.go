package utils

import (
	"regexp"
	"time"
)

type TimeUnit string

const (
	UNIT_MONTH  TimeUnit = "month"
	UNIT_DAY    TimeUnit = "day"
	UNIT_HOUR   TimeUnit = "hour"
	UNIT_MINUTE TimeUnit = "minute"
	UNIT_SECOND TimeUnit = "second"

	DAY_PER_MONTH     = 31
	HOUR_PER_DAY      = 24
	MINUTE_PER_HOUR   = 60
	SECOND_PER_MINUTE = 60
)

const (
	// Microsecond
	DATE_FORMAT                                = "2006-01-02"
	TIME_FORMAT                                = "2006-01-02 15:04:05"
	TIME_FORMAT_WITH_MICRO_SEC          string = "2006-01-02 15:04:05.999999"
	TIME_FORMAT_WITH_MICRO_SEC_TIMEZONE string = "2006-01-02 15:04:05.999999 -0700"
	QUERY_TIME_FLOAT_LIMIT                     = 999999000
)

var letterRunes = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
var alphanumericRegex = regexp.MustCompile(`^[a-zA-Z0-9]+$`)
var TaiwanTimezone = time.FixedZone("", 8*60*60)
