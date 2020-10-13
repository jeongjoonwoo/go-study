package main

import "fmt"

type student struct {
	name      string
	sex       string
	scoreData map[string]int
}

func newScore() *student {
	studentData := student{}
	studentData.scoreData = map[string]int{}
	return &studentData
}

func main() {
	var stuNum, subNum, score int
	var name, sex, subject string

	fmt.Scanln(&stuNum, &subNum)

	s := make([]student, stuNum)

	for i := 0; i < stuNum; i++ {
		fmt.Scanln(&name, &sex)
		student := newScore()
		student.name = name
		student.sex = sex
		for j := 0; j < subNum; j++ {
			fmt.Scanln(&subject, &score)

			student.scoreData[subject] = score
		}
		s[i] = *student
	}

	for i := 0; i < stuNum; i++ {
		fmt.Println("----------")
		fmt.Println(s[i].name, s[i].sex)

		for index, val := range s[i].scoreData {
			fmt.Println(index, val)
		}

	}
	fmt.Println("----------")
}
