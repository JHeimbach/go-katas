package kata

const DalmatiansHardlyAny string = "Hardly any"
const DalmatiansMoreThanHandful string = "More than a handful!"
const Dalmatians101 string = "101 DALMATIONS!!!"
const DalmatiansLotsOfDogs string = "Woah that's a lot of dogs!"

func HowManyDalmatians(number int) string {
	message := ""
	switch {
	case number <= 10:
		message = DalmatiansHardlyAny
	case number <= 50:
		message = DalmatiansMoreThanHandful
	case number == 101:
		message = Dalmatians101
	default:
		message = DalmatiansLotsOfDogs
	}
	return message
}
