package utils

import "strconv"

func StrToUint64(v string) (uint64, error) {
	itemIDUint64, err := strconv.ParseUint(v, 10, 64)
	if err != nil {
		return 0, err
	}
	return itemIDUint64, nil
}
