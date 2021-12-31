package main

import (
	"fmt"
	"testing"
)

func TestGetUserBmi(t *testing.T) {
	bmiTest1, error1 := GetUserBmi(77, 1.78)
	fmt.Println("bmiTest1:", bmiTest1, " error1:", error1)
	if bmiTest1 < 0 || error1 != nil {
		t.Errorf("%s", error1)
	}
	bmiTest2, error2 := GetUserBmi(0, 1.78)
	fmt.Println("bmiTest2:", bmiTest2, " error2: ", error2)
	if bmiTest2 != 0 || error2 == nil {
		t.Errorf("%s", error2)
	}
	bmiTest3, error3 := GetUserBmi(77, 0)
	fmt.Println("bmiTest3:", bmiTest3, " error3:", error3)
	if bmiTest3 != 0 || error3 == nil {
		t.Errorf("%s", error3)
	}
	bmiTest4, error4 := GetUserBmi(77, -1)
	fmt.Println("bmiTest4:", bmiTest4, " error3:", error4)
	if bmiTest4 != 0 || error4 == nil {
		t.Errorf("%s", error4)
	}
}

func TestGetUserFateRate(t *testing.T) {
	fatRate1, suggest1, error1 := GetUserFateRate(21.3, 29, "男")
	fmt.Println("fatRate1:", fatRate1, " suggest1:", suggest1, " error1:", error1)
	if fatRate1 < 0 || error1 != nil {
		t.Errorf("%s", error1)
	}

	_, _, error2 := GetUserFateRate(-1, 29, "女")
	if error2 != nil {
		t.Errorf("%s", error2)
	}

	_, _, error3 := GetUserFateRate(23.5, 0, "男")
	if error3 != nil {
		t.Errorf("%s", error3)
	}

	_, _, error4 := GetUserFateRate(29, 151, "男")
	if error4 != nil {
		t.Errorf("%s", error4)
	}

	_, _, error5 := GetUserFateRate(29, 33, "男女")
	if error5 != nil {
		t.Errorf("%s", error5)
	}
}
