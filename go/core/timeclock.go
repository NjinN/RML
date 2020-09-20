package core

import "fmt"
import "strings"
import "strconv"
import "regexp"

var _DAYS_OF_MONTHS = [12]int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

type TimeClock struct {
	Negative		bool
	Date			int
	Second			int
	FloatSecond		float64
}

func ParseTimeStr(timeStr string) *TimeClock {
	result := TimeClock{}
	str := timeStr

	
	matched, err := regexp.MatchString("^\\-?[0-9]{4}-[0-9]{2}-[0-9]{2}$", str)
	if !matched || err != nil {
		matched, err = regexp.MatchString("^\\-?[0-9]{2}:[0-9]{2}:[0-9]{2}(\\.[0-9]{1,8})?$", str)
		if !matched || err != nil {
			matched, err = regexp.MatchString("^\\-?[0-9]{4}-[0-9]{2}-[0-9]{2}\\+[0-9]{2}:[0-9]{2}:[0-9]{2}(\\.[0-9]{1,8})?$", str)
		}
	}
	if !matched || err != nil {
		return nil
	}

	if str[0] == '-' {
		result.Negative = true
		str = str[1:]
	}

	slices := strings.Split(str, "+")

	result.Date = DateStrToDays(slices[0])
	result.Second = TimeStrToSecs(slices[1])

	if strings.Index(slices[1], ".") >= 0 {
		nsStr := strings.Split(slices[1], ".")[1]
		nsStr = fmt.Sprintf("s%09", nsStr)
		ns, err := strconv.ParseFloat(nsStr, 64)
		if err != nil {
			fmt.Println(err.Error())
			return nil
		}
		result.FloatSecond = ns
	}

	return &result
}

func IsLeapYear(y int) bool{
	if y % 400 == 0 {
		return true
	}else if y % 4 == 0 {
		if y % 100 != 0 {
			return true
		}
	}
	return false
}

func DateStrToDays(dateStr string) int{
	dateSlice :=  strings.Split(dateStr, "-")
	if len(dateSlice) != 3 {
		return 0
	}

	y, err := strconv.Atoi(dateSlice[0])
	if err != nil {
		return 0
	}

	m, err := strconv.Atoi(dateSlice[1])
	if err != nil {
		return 0
	}

	d, err := strconv.Atoi(dateSlice[2])
	if err != nil {
		return 0
	}

	leaps := 0
	for i := 4; i < y; i++ {
		if IsLeapYear(i){
			leaps++
		}
	}
	
	days := (y - 1) * 365 + leaps

	for i := 1; i < m; i++ {
		days += _DAYS_OF_MONTHS[i-1]
	}

	if IsLeapYear(y) && m > 2 {
		days++
	}

	days += d
	
	return days
}

func DateToDays(y int, m int, d int) int {
	leaps := 0
	for i := 4; i < y; i++ {
		if IsLeapYear(i){
			leaps++
		}
	}
	
	days := (y - 1) * 365 + leaps

	for i := 1; i < m; i++ {
		days += _DAYS_OF_MONTHS[i-1]
	}

	if IsLeapYear(y) && m > 2 {
		days++
	}

	days += d
	
	return days
}


func DaysToDate(days int) string{
	y := 1
	m := 1

	for days > 365 {
		if IsLeapYear(y){
			if days == 366 {
				break
			}else{
				days -= 366
			}
		}else{
			days -= 365
		}
		y++
	}

	for days > _DAYS_OF_MONTHS[m-1] {
		if IsLeapYear(y) && m == 2 {
			if days == 29 {
				break
			}else{
				days -= 29
			}
		}else{
			days -= _DAYS_OF_MONTHS[m-1]
		}

		m++
	} 

	yearStr := fmt.Sprintf("%04d", y)
	monthStr := fmt.Sprintf("%02d", m)
	dayStr := fmt.Sprintf("%02d", days)

	return yearStr + "-" + monthStr + "-" + dayStr
}


func TimeStrToSecs(timeStr string) int{
	timeSlice :=  strings.Split(timeStr, ":")
	if len(timeSlice) != 3 {
		return -1
	}

	h, err := strconv.Atoi(timeSlice[0])
	if err != nil {
		return -1
	}

	m, err := strconv.Atoi(timeSlice[1])
	if err != nil {
		return -1
	}

	s, err := strconv.Atoi(timeSlice[2])
	if err != nil {
		return -1
	}

	return  h * 60 * 60 + m * 60 + s

}

func SecsToTimeStr(secs int) string{
	h := int(secs / 3600)
	secs -= h * 3600
	m := int(secs / 60)
	secs -= m * 60

	hStr := fmt.Sprintf("%02d", h)
	mStr := fmt.Sprintf("%02d", m)
	sStr := fmt.Sprintf("%02d", secs)

	return hStr + ":" + mStr + ":" + sStr
}


func (tc *TimeClock) Format(){
	if tc.FloatSecond > 0 {
		for tc.FloatSecond >= 1 {
			tc.Second++
			tc.FloatSecond--
		}
	}else{
		for tc.FloatSecond < 0 {
			tc.Second--
			tc.FloatSecond++
		}
	}

	if tc.Second > 0 {
		for tc.Second > 60*60*24 {
			tc.Second -= 60*60*24
			tc.Date++
		}
	}else{
		for tc.Second < 0 {
			tc.Second += 60*60*24
			tc.Date--
		}
	}

	if tc.Date < 0 {
		tc.Negative = !tc.Negative
		tc.Date = -1 * tc.Date
	}
}


func TimeClockAdd(tc1 *TimeClock, tc2 *TimeClock) *TimeClock{
	result := TimeClock{}

	if !tc1.Negative {
		if !tc2.Negative {
			result.Date = tc1.Date + tc2.Date
			result.Second = tc1.Second + tc2.Second
			result.FloatSecond = tc1.FloatSecond + tc2.FloatSecond
		}else{
			result.Date = tc1.Date - tc2.Date
			result.Second = tc1.Second - tc2.Second
			result.FloatSecond = tc1.FloatSecond - tc2.FloatSecond
		}

	}else{
		if !tc2.Negative {
			result.Date = tc1.Date - tc2.Date
			result.Second = tc1.Second - tc2.Second
			result.FloatSecond = tc1.FloatSecond - tc2.FloatSecond
		}else{
			result.Date = tc1.Date + tc2.Date
			result.Second = tc1.Second + tc2.Second
			result.FloatSecond = tc1.FloatSecond + tc2.FloatSecond
		}

	}

	result.Format()
	return &result
}

func TimeClockSub(tc1 *TimeClock, tc2 *TimeClock) *TimeClock{
	result := TimeClock{}

	if !tc1.Negative {
		if !tc2.Negative {
			result.Date = tc1.Date - tc2.Date
			result.Second = tc1.Second - tc2.Second
			result.FloatSecond = tc1.FloatSecond - tc2.FloatSecond
		}else{
			result.Date = tc1.Date + tc2.Date
			result.Second = tc1.Second + tc2.Second
			result.FloatSecond = tc1.FloatSecond + tc2.FloatSecond
		}

	}else{
		if !tc2.Negative {
			result.Date = tc1.Date + tc2.Date
			result.Second = tc1.Second + tc2.Second
			result.FloatSecond = tc1.FloatSecond + tc2.FloatSecond
		}else{
			result.Date = tc1.Date - tc2.Date
			result.Second = tc1.Second - tc2.Second
			result.FloatSecond = tc1.FloatSecond - tc2.FloatSecond
		}

	}

	result.Format()
	return &result
}




