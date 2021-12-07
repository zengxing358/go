package main

import "fmt"

func getarr(n int) []int {
	var res []int
	for i := 1; i <= n; i++ {
		res = append(res, i)
	}
	return res
}

func main() {
	//反转
	var ages [6]int = [6]int{35, 32, 33, 37, 59, 78}
	for i := 0; i < len(ages)>>1; i++ {
		ages[i], ages[len(ages)-1-i] = ages[len(ages)-1-i], ages[i]
	}
	fmt.Println(ages)

	//日历
	var calendar [12][]int
	var year int
	fmt.Print("请输入年份：")
	fmt.Scanln(&year)
	is_leap := false
	if (year%400 == 0) || ((year%4 == 0) && (year%100 != 0)) {
		is_leap = true
	}
	var month []int
	var cur int
	for i := 0; i < 12; i++ {
		cur = i + 1
		if cur == 1 || cur == 3 || cur == 5 || cur == 7 || cur == 8 || cur == 10 || cur == 12 {
			month = getarr(31)
		} else if cur == 2 {
			if is_leap {
				month = getarr(29)
			} else {
				month = getarr(28)
			}
		} else {
			month = getarr(30)
		}
		calendar[i] = month
	}
	fmt.Println(calendar)

	//日历_Week
	var calendar2 [12][][]int
	var month2 [][]int
	var cur2 int
	var tmp int
	for i := 0; i < 12; i++ {
		cur2 = i + 1
		tmp = ZellerFunction2Week(year, cur2, 1)
		if cur2 == 1 || cur2 == 3 || cur2 == 5 || cur2 == 7 || cur2 == 8 || cur2 == 10 || cur2 == 12 {
			month2 = getweekarr(31, tmp)
		} else if cur2 == 2 {
			if is_leap {
				month2 = getweekarr(29, tmp)
			} else {
				month2 = getweekarr(28, tmp)
			}
		} else {
			month2 = getweekarr(30, tmp)
		}
		calendar2[i] = month2
	}
	fmt.Println(calendar2)
}
func getweekarr(days, start int) [][]int {
	var res [][]int
	var tmp []int
	for i := 1; i < start; i++ {
		tmp = append(tmp, -1)
	}
	for i := 1; i <= days; i++ {
		tmp = append(tmp, i)
		if len(tmp) == 7 {
			if start > 0 {
				tmp = tmp[start-1:]
				start = 0
			}
			res = append(res, tmp)
			tmp = nil
		}
	}
	if len(tmp) > 0 {
		res = append(res, tmp)
	}
	return res
}
func ZellerFunction2Week(year, month, day int) int {
	var y, m, c int
	var weekday = [7]int{7, 1, 2, 3, 4, 5, 6}
	if month >= 3 {
		m = month
		y = year % 100
		c = year / 100
	} else {
		m = month + 12
		y = (year - 1) % 100
		c = (year - 1) / 100
	}
	week := y + (y / 4) + (c / 4) - 2*c + ((26 * (m + 1)) / 10) + day - 1
	if week < 0 {
		week = 7 - (-week)%7
	} else {
		week = week % 7
	}
	which_week := int(week)
	return weekday[which_week]
}
