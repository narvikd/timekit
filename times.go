package timekit

// IsBetweenTimesBool returns true or false depending on the result of "IsNowBetweenTimes"
//
// "start" and "end" must be in the hh:mm:ss format.
func IsBetweenTimesBool(start string, end string) bool {
	return returnFalseIfErr(IsNowBetweenTimes(start, end))
}

// IsMidnight returns true if "now" is between "00:00:00" and "01:00:59"
func IsMidnight() bool {
	start := "00:00:00"
	end := "01:00:59"
	return returnFalseIfErr(IsNowBetweenTimes(start, end))
}

// IsDawn returns true if "now" is between "01:01:00" and "05:59:59"
func IsDawn() bool {
	start := "01:01:00"
	end := "05:59:59"
	return returnFalseIfErr(IsNowBetweenTimes(start, end))
}

// IsMorning returns true if "now" is between "06:00:00" and "11:59:59"
func IsMorning() bool {
	start := "06:00:00"
	end := "11:59:59"
	return returnFalseIfErr(IsNowBetweenTimes(start, end))
}

// IsAfternoon returns true if "now" is between "12:00:00" and "17:59:59"
func IsAfternoon() bool {
	start := "12:00:00"
	end := "17:59:59"
	return returnFalseIfErr(IsNowBetweenTimes(start, end))
}

// IsEvening returns true if "now" is between "18:00:00" and "23:59:59"
func IsEvening() bool {
	start := "18:00:00"
	end := "23:59:59"
	return returnFalseIfErr(IsNowBetweenTimes(start, end))
}

// returnFalseIfErr is a function I had to make in case I'm programming at 2 AM and I can't distinguish between left and right
func returnFalseIfErr(err error) bool {
	if err != nil {
		return false
	}
	return true
}
