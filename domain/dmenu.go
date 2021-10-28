package domain

import (
	"strconv"
	"strings"
)

func ParseDmenuEntryNumber(dmenuOutput string) (int, error) {
	num, err := strconv.ParseInt(strings.Trim(dmenuOutput, "\n"), 10, 32)
	return int(num), err
}
