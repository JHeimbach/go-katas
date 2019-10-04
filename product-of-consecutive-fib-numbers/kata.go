package kata

func ProductFib(searchFor uint64) [3]uint64 {
	var product uint64
	var fibOne, fibTwo uint64 = 0, 1
	var found uint64

	for {
		fibOne, fibTwo = fibTwo, fibOne+fibTwo
		product = fibOne * fibTwo

		if product == searchFor {
			found = 1
		}

		if product >= searchFor {
			return [3]uint64{fibOne, fibTwo, found}
		}
	}
}
