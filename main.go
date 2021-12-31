package main

import (
	"fmt"
	gotdsbmi "github.com/armstrongli/go-bmi"
)

func main() {
	intPeopleNum := giveUserTip()
	floatTotalBodyFat := 0.0
	for i := 0; i < intPeopleNum; i++ {
		name, age, sexVal, weight, tall := GetUserInfo()
		bmi, err := gotdsbmi.BMI(weight, tall)
		if err != nil {
			fmt.Println(err)
		}
		fateRate := CalcFateRate(bmi, age, sexVal)
		suggestions := GetSuggestions(fateRate, sexVal, age)
		fmt.Printf("姓名:%s的bmi是:%f,体脂率是:%f,给你的建议是:%s\n", name, bmi, fateRate, suggestions)
		floatTotalBodyFat += fateRate
	}
	fmt.Printf("总共输入%d人的信息，%d人的体脂率平均值为:%f", intPeopleNum, intPeopleNum, floatTotalBodyFat/float64(intPeopleNum))
}

func giveUserTip() int {
	intPeopleNum := 0
	fmt.Println("请输入要录入的人数")
	fmt.Scanln(&intPeopleNum)
	return intPeopleNum
}

func GetUserInfo() (name string, age int, sexVal int, weight float64, tall float64) {
	fmt.Print("请输入姓名：")
	fmt.Scanln(&name)

	fmt.Print("请输入年龄（18-150岁）：")
	fmt.Scanln(&age)
	for age < 18 || age > 150 {
		fmt.Print("年龄超出指定的范围(18-150)，请重新输入")
		fmt.Print("请输入年龄（18-150岁）：")
		fmt.Scanln(&age)
	}
	mapSex := map[string]int{
		"男": 1,
		"女": 0,
	}
	var sex string
	fmt.Print("请输入性别（男/女）：")
	fmt.Scanln(&sex)
	for {
		_, ok := mapSex[sex]
		if ok == false {
			fmt.Print("性别输入异常（男/女），请重新输入")
			fmt.Print("请输入性别（男/女）：")
			fmt.Scanln(&sex)
		}
		break
	}
	sexVal = mapSex[sex]
	fmt.Print("请输入体重（20kg-1000kg）：")
	fmt.Scanln(&weight)
	for weight < 20 || weight > 1000 {
		fmt.Print("体重超出指定的范围(20-1000)，请重新输入")
		fmt.Print("请输入体重（20kg-1000kg）：")
		fmt.Scanln(&weight)
	}

	fmt.Print("请输入身高（0.5米-3米）：")
	fmt.Scanln(&tall)
	for tall < 0.5 || tall > 3 {
		fmt.Print("身高超出指定的范围(0.5米-3米)")
		fmt.Scanln(&tall)
	}
	return
}

func CalcFateRate(bmi float64, age int, sexVal int) (bfr float64) {
	bfr = (1.2*bmi + 0.23*float64(age) - 5.4 - 10.8*(float64(sexVal))) / 100
	return
}

func GetSuggestions(rate float64, val int, age int) (suggestions string) {
	if val == 1 {
		suggestions = getMaleSuggestion(rate, age)
	} else {
		suggestions = getFeMaleSuggestion(rate, age)
	}
	return
}

func getMaleSuggestion(floatBodyFat float64, intAge int) (suggestions string) {
	suggestions = ""
	switch {
	case intAge >= 18 && intAge <= 39:
		switch {
		case floatBodyFat > 0.0 && floatBodyFat <= 0.10:
			suggestions = "你的体型属于偏瘦"
		case floatBodyFat > 0.1 && floatBodyFat <= 0.16:
			suggestions = "你的体型属于正常"
		case floatBodyFat > 0.16 && floatBodyFat <= 0.21:
			suggestions = "你的体型属于偏重"
		case floatBodyFat > 0.21 && floatBodyFat <= 0.26:
			suggestions = "你的体型属于肥胖"
		case floatBodyFat > 0.26:
			suggestions = "你的体型属于严重肥胖"
		default:
			suggestions = "你的体脂率小于0，请核对输入的信息"
		}
	case intAge >= 40 && intAge <= 59:
		switch {
		case floatBodyFat > 0 && floatBodyFat <= 0.11:
			suggestions = "你的体型属于偏瘦"
		case floatBodyFat > 0.11 && floatBodyFat <= 0.17:
			suggestions = "你的体型属于正常"
		case floatBodyFat > 0.17 && floatBodyFat <= 0.22:
			suggestions = "你的体型属于偏重"
		case floatBodyFat > 0.22 && floatBodyFat <= 0.27:
			suggestions = "你的体型属于肥胖"
		case floatBodyFat > 0.27:
			suggestions = "你的体型属于严重肥胖"
		default:
			suggestions = "你的体脂率小于0，请核对输入的信息"
		}
	case intAge >= 60:
		switch {
		case floatBodyFat > 0 && floatBodyFat <= 0.13:
			suggestions = "你的体型属于偏瘦"
		case floatBodyFat > 0.13 && floatBodyFat <= 0.19:
			suggestions = "你的体型属于正常"
		case floatBodyFat > 0.19 && floatBodyFat <= 0.24:
			suggestions = "你的体型属于偏重"
		case floatBodyFat > 0.24 && floatBodyFat <= 0.29:
			suggestions = "你的体型属于肥胖"
		case floatBodyFat > 0.29:
			suggestions = "你的体型属于严重肥胖"
		default:
			suggestions = "你的体脂率小于0，请核对输入的信息"
		}
	default:
		suggestions = "输入的年龄异常，请核对输入的信息"
	}
	return
}

//女生输出的建议
func getFeMaleSuggestion(floatBodyFat float64, intAge int) (suggestions string) {
	suggestions = ""
	switch {
	case intAge >= 18 && intAge <= 39:
		switch {
		case floatBodyFat > 0 && floatBodyFat <= 0.2:
			suggestions = "你的体型属于偏瘦"
		case floatBodyFat > 0.2 && floatBodyFat <= 0.27:
			suggestions = "你的体型属于正常"
		case floatBodyFat > 0.27 && floatBodyFat <= 0.34:
			suggestions = "你的体型属于偏重"
		case floatBodyFat > 0.34 && floatBodyFat <= 0.39:
			suggestions = "你的体型属于肥胖"
		case floatBodyFat > 0.39:
			suggestions = "你的体型属于严重肥胖"
		default:
			suggestions = "你的体脂率小于0，请核对输入的信息"
		}
	case intAge >= 40 && intAge <= 59:
		switch {
		case floatBodyFat > 0 && floatBodyFat <= 0.21:
			suggestions = "你的体型属于偏瘦"
		case floatBodyFat > 0.21 && floatBodyFat <= 0.28:
			suggestions = "你的体型属于正常"
		case floatBodyFat > 0.28 && floatBodyFat <= 0.35:
			suggestions = "你的体型属于偏重"
		case floatBodyFat > 0.35 && floatBodyFat <= 0.40:
			suggestions = "你的体型属于肥胖"
		case floatBodyFat > 0.40:
			suggestions = "你的体型属于严重肥胖"
		default:
			suggestions = "你的体脂率小于0，请核对输入的信息"
		}
	case intAge >= 60:
		switch {
		case floatBodyFat > 0 && floatBodyFat <= 0.22:
			suggestions = "你的体型属于偏瘦"
		case floatBodyFat > 0.22 && floatBodyFat <= 0.2:
			suggestions = "你的体型属于正常"
		case floatBodyFat > 0.29 && floatBodyFat <= 0.36:
			suggestions = "你的体型属于偏重"
		case floatBodyFat > 0.36 && floatBodyFat <= 0.41:
			suggestions = "你的体型属于肥胖"
		case floatBodyFat > 0.41:
			suggestions = "你的体型属于严重肥胖"
		default:
			suggestions = "你的体脂率小于0，请核对输入的信息"
		}
	default:
		suggestions = "输入的年龄异常，请核对输入的信息"
	}
	return
}
