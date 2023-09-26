package utils

import (
	"time"
)

func ParseTimeInterval(from, to string) (fromTime, toTime time.Time, success bool) {
	var err error

	if fromTime, err = time.ParseInLocation(TIME_FORMAT, from, TaiwanTimezone); err != nil {
		return fromTime, toTime, false
	}

	if toTime, err = time.ParseInLocation(TIME_FORMAT, to, TaiwanTimezone); err != nil {
		return fromTime, toTime, false
	}

	return fromTime, toTime, true
}

// ParseTimeInterval2 return the 'toTime' which will add 0.999999000 seconds
// i.e. input toTime = 2023-07-09 00:00:00, then output toTime = 2023-07-08 23:59:59.999999000
func ParseTimeIntervalExcludeToTime(from, to string) (fromTime, toTime time.Time, success bool) {
	var err error

	if fromTime, err = time.Parse(TIME_FORMAT, from); err != nil {
		return fromTime, toTime, false
	}

	if toTime, err = time.Parse(TIME_FORMAT, to); err != nil {
		return fromTime, toTime, false
	}

	newToTime := toTime.Unix() - 1
	excludeToTime := time.Unix(newToTime, 0)
	excludeToTime = time.Unix(excludeToTime.Unix(), QUERY_TIME_FLOAT_LIMIT).UTC()

	return fromTime, excludeToTime, true
}

func IsTimeDiffLessThan(from, to time.Time, interval int, unit TimeUnit) bool {
	timeDiff := to.Sub(from).Seconds()
	if 0 >= timeDiff {
		return false
	}

	maxBoundOfSeconds := getMaxBoundOfSeconds(interval, unit)

	if timeDiff >= maxBoundOfSeconds {
		return false
	}

	return true
}

func getMaxBoundOfSeconds(interval int, unit TimeUnit) float64 {
	var maxBoundOfSeconds float64
	switch unit {
	case UNIT_MONTH:
		maxBoundOfSeconds = float64(interval * DAY_PER_MONTH * HOUR_PER_DAY * MINUTE_PER_HOUR * SECOND_PER_MINUTE)
	case UNIT_DAY:
		maxBoundOfSeconds = float64(interval * HOUR_PER_DAY * MINUTE_PER_HOUR * SECOND_PER_MINUTE)
	case UNIT_HOUR:
		maxBoundOfSeconds = float64(interval * MINUTE_PER_HOUR * SECOND_PER_MINUTE)
	case UNIT_MINUTE:
		maxBoundOfSeconds = float64(interval * SECOND_PER_MINUTE)
	default:
		maxBoundOfSeconds = float64(interval)
	}

	return maxBoundOfSeconds
}

func FormatTimeToMicrosecondString(when time.Time) string {
	return when.Format(TIME_FORMAT_WITH_MICRO_SEC)
}
func GetUTC8Time() time.Time {
	cstZone := time.FixedZone("CST", 8*3600)
	now := time.Now()
	utc8 := now.In(cstZone)
	return utc8
}
