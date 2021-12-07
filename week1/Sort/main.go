package main

import (
	"fmt"
	"sort"
)

type student struct {
	idx      int
	score    int
	strength int
}

type ByScore []student

func (a ByScore) Len() int { return len(a) }
func (a ByScore) Less(i, j int) bool {
	if a[i].score == a[j].score {
		return a[i].idx < a[j].idx
	}
	return a[i].score > a[j].score
}
func (a ByScore) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func cmp(a student, b student) bool {
	if a.score == b.score {
		return a.idx < b.idx
	}
	return a.score > b.score
}

func main() {
	var n, r, q int
	const N int = 1e5*2 + 20
	fmt.Scanln(&n, &r, &q)
	var students []student
	var p1 [N]int
	var p2 [N]int
	for i := 0; i < 2*n; i++ {
		fmt.Scan(&p1[i])
	}
	for i := 0; i < 2*n; i++ {
		fmt.Scan(&p2[i])
	}

	for i := 0; i < 2*n; i++ {
		students = append(students, student{i + 1, p1[i], p2[i]})
	}
	sort.Sort(ByScore(students))
	var q1 []student
	var q2 []student
	var l1, l2 int
	for i := 1; i <= r; i++ {
		q1 = nil
		q2 = nil
		for j := 0; j < 2*n; j += 2 {
			if students[j].strength > students[j+1].strength {
				students[j].score += 1
				q1 = append(q1, students[j])
				q2 = append(q2, students[j+1])
			} else {
				students[j+1].score += 1
				q1 = append(q1, students[j+1])
				q2 = append(q2, students[j])
			}
		}
		l1 = 0
		l2 = 0
		for {
			if (l1 >= len(q1)) || (l2 >= len(q2)) {
				break
			}
			if cmp(q1[l1], q2[l2]) {
				students[l1+l2] = q1[l1]
				l1++
			} else {
				students[l1+l2] = q2[l2]
				l2++
			}
		}
		for {
			if l1 >= len(q1) {
				break
			}
			students[l1+l2] = q1[l1]
			l1++
		}
		for {
			if l2 >= len(q2) {
				break
			}
			students[l1+l2] = q2[l2]
			l2++
		}
	}
	fmt.Println(students[q-1].idx)
}
