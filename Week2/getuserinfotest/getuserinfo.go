package main

import "fmt"

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
func GetUserBmi(weightKG, heightM float64) (bmi float64, err error) {
	if weightKG < 0 {
		err = fmt.Errorf("录入的体重小于0kg,请核对信息后重新录入")
		return
	}
	if heightM <= 0 {
		err = fmt.Errorf("输入的身高小于等于0m,请核对信息后重新录入")
		return
	}
	bmi = weightKG / (heightM * heightM)
	return
}

func GetUserFateRate(bmi float64, age int, sex string) (fateRate float64, suggest string, err error) {
	fateRate = 0.0
	suggest = ""
	if bmi <= 0 {
		err = fmt.Errorf("bmi录入不能为负数")
		return
	}

	if age <= 0 || age > 150 {
		err = fmt.Errorf("年龄不能为负数或者大于150")
		return
	}
	mapSex := map[string]int{
		"男": 1,
		"女": 0,
	}
	sexVal, ok := mapSex[sex]
	if ok == false {
		err = fmt.Errorf("性别输入异常，只能男/女")
	}
	fateRate = (1.2*bmi + 0.23*float64(age) - 5.4 - 10.8*(float64(sexVal))) / 100
	return
}
