package util

import "errors"

func IndexOf(value string, slice []string) (int, error) {
	for k, str := range slice {
		if str == value {
			return k, nil
		}
	}
	return -1, errors.New("Value does not exist in slice")
}
