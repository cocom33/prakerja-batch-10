package main

import "fmt"

func main() {
	// soal 1
	var car Car
	car.tipe = "kijang"
	car.fuelIn = 30
	jarak := car.HitungJarak()
	fmt.Println(jarak, "mil")

	// soal 2
	students := []Student{
		{name: "agus", score: 60},
		{name: "joni", score: 50},
		{name: "budi", score: 75},
		{name: "joko", score: 80},
		{name: "tone", score: 90},
	}

	var studentCheck StudentCheck
	studentCheck.Average(students)
	studentCheck.MaxScore(students)
	studentCheck.MinScore(students)

	// soal 3
	var a1, a2, a3, a4, a5, a6 int
	fmt.Scan(&a1)
	fmt.Scan(&a2)
	fmt.Scan(&a3)
	fmt.Scan(&a4)
	fmt.Scan(&a5)
	fmt.Scan(&a6)
	min, max := getMinMax(&a1, &a2, &a3, &a4, &a5, &a6)
	fmt.Println()
	fmt.Println(min)
	fmt.Println(max)
}

// soal 3
func getMinMax(numbers ...*int) (int, int) {
	var min, max int = 1, 1
	for _, i := range numbers {
		if *i < min {
			min = *i
		}
		if *i > max {
			max = *i
		}
	}
	return min, max
}
// soal 2
type Student struct {
	name  string
	score int
}

// soal 2
type StudentCheck interface {
	Average(students []Student)
	MaxScore(students []Student)
	MinScore(students []Student)
}

// soal 2
func (v Student) Average(students []Student) float64 {
	totalScore := 0
	for _, student := range students {
		totalScore += student.score
	}
	return float64(totalScore) / float64(len(students))
}
func (v Student) MaxScore(students []Student) Student {
	maxScore := students[0]
	for _, student := range students {
		if student.score > maxScore.score {
			maxScore.score = student.score
			maxScore.name = student.name
		}
	}
	return maxScore
}
func (v Student) MinScore(students []Student) Student {
	minScore := students[0]
	for _, student := range students {
		if student.score < minScore.score {
			minScore.score = student.score
			minScore.name = student.name
		}
	}
	return minScore
}

// soal 1
type Car struct {
	tipe   string
	fuelIn float32
}

// soal 1
func (car Car) HitungJarak() float32 {
	return car.fuelIn / 1.5
}