package main

func Primeiro[T any](s []T) (T, bool) {
	if len(s) == 0 {
		var zero T
		return zero, false
	}

	return s[0], true
}

func Unicos[T comparable](s []T) []T {
	vistos := make(map[T]struct{}, len(s))
	out := make([]T, 0, len(s))

	for _, v := range s {
		if _, ja := vistos[v]; ja {
			continue
		}

		vistos[v] = struct{}{}
		out = append(out, v)
	}

	return out
}
