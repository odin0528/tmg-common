package geoip

import (
	"errors"
	"mgmt/common/configs"
	"net"

	"github.com/oschwald/geoip2-golang"
)

func InitGeoIPRepository() error {
	fileName := configs.Get(configs.SECTION_GEOIP, configs.GEOIP_FILE_PATH, "/opt/geoip/GeoLite2-Country.mmdb")

	r, err := geoip2.Open(fileName)
	if err != nil {
		return errors.New("error opening database: " + err.Error())
	}

	reader = r

	return nil
}

func GetCountry(ip net.IP) (string, error) {
	if reader == nil {
		return "", errors.New("reader is not initialized")
	}

	record, err := reader.Country(ip)
	if err != nil {
		return "", err
	}

	return record.Country.IsoCode, nil
}
