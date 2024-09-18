package geoip

import (
	"github.com/oschwald/geoip2-golang"
)

const (
	UNKNOWN_COUNTRY = "UNKNOWN"
)

var (
	reader *geoip2.Reader
)
