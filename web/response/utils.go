package response

import (
	"errors"
	"fmt"
	"strings"
)

func GetCustomFormatError(str string, args ...interface{}) error {
	count := strings.Count(str, "%s")
	if count > len(args) {
		return errors.New("unexpected string:" + str)
	}
	return fmt.Errorf(str, args[:count]...)
}
