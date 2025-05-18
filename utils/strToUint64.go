package utils

import "strconv"

func StrToint(v string) (int, error) {
	itemIDint, err := strconv.ParseInt(v, 10, 64)
	if err != nil {
		return 0, err
	}
	return int(itemIDint), nil

}
