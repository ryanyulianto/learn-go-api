package util

func OpParam[T any](param []T, default_p T) T {
	if len(param) > 0 {
		return param[0]
	}
	return default_p
}
