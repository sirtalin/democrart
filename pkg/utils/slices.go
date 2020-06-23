package utils

func UintInSlice(elem uint, list []uint) bool {
	for _, b := range list {
		if b == elem {
			return true
		}
	}
	return false
}
