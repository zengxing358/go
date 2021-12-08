package main

import (
	"fmt"
	"sort"
)

type student struct {
	score int
	name  string
}
type ByScore []student

func (a ByScore) Len() int { return len(a) }
func (a ByScore) Less(i, j int) bool {
	return a[i].score > a[j].score
}
func (a ByScore) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func main() {
	names := []string{"王强", "李静", "苗文", "周毅", "旭东", "张松", "小强", "郑荒", "杨田", "李玄", "王田", "欧阳地", "张荒", "陈木", "王山", "沈石", "郑地", "沈冬", "王火", "郑夏"}
	scores := []int{60, 83, 99, 95, 67, 57, 83, 96, 60, 32, 94, 96, 100, 94, 94, 95, 32, 62, 67, 57}
	maps := map[string]int{}
	for i := 0; i < 20; i++ {
		maps[names[i]] = scores[i]
	}
	fmt.Println(maps)
	var all float32
	for _, score := range maps {
		all += float32(score)
	}
	fmt.Printf("平均分是%.2f\n", all/20)
	var students []student
	for name, score := range maps {
		students = append(students, student{score, name})
	}
	sort.Sort(ByScore(students))
	var sortarray [][]string
	pre := -1
	var tmp []string
	for i := 0; i < 20; i++ {
		if pre == -1 || pre == students[i].score {
			tmp = append(tmp, students[i].name)
		} else {
			sortarray = append(sortarray, tmp)
			tmp = nil
			tmp = append(tmp, students[i].name)
		}
		pre = students[i].score
	}
	if len(tmp) > 0 {
		sortarray = append(sortarray, tmp)
	}
	fmt.Println(sortarray)
}
