package utils

func UnpackSliceString(s []string, vars ...*string) {
	for i, str := range s {
		*vars[i] = str
	}
}
