package kata

const (
	NorthDirection = "NORTH"
	SouthDirection = "SOUTH"
	WestDirection  = "WEST"
	EastDirection  = "EAST"
)

func DirReduc(arr []string) []string {
	for i := 0; i < len(arr)-1; i++ {
		if cancelOut(arr[i], arr[i+1]) {
			return DirReduc(append(arr[:i], arr[i+2:]...))
		}
	}
	return arr
}

func cancelOut(s1, s2 string) bool {
	if s1 == NorthDirection && s2 == SouthDirection || s1 == SouthDirection && s2 == NorthDirection {
		return true
	}

	if s1 == WestDirection && s2 == EastDirection || s1 == EastDirection && s2 == WestDirection {
		return true
	}

	return false
}
