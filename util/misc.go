package util

func IndexOf(value string, slice []string) int {
	for k, str := range slice {
		if str == value {
			return k
		}
	}
	return -1
}
