package main

import "fmt"

type Student struct {
	name   string
	grades []int
	age    int
}

/*func (s Student) setAge(age int) {
	s.age = age
}*/

func (s Student) getAverageGrade() float32 {
	sum := 0
	for _, v := range s.grades {
		sum += v
	}
	return float32(sum) / float32(len(s.grades))
}

func (s *Student) getMaxGrade() int {
	curMax := 0
	for _, v := range s.grades {
		if curMax < v {
			curMax = v
		}
	}
	return curMax
}

func main() {
	s1 := Student{"Tim", []int{70, 90, 80, 85}, 19}
	s2 := Student{"Joe", []int{70, 90, 80, 85, 90, 90, 95}, 21}
	averge := s1.getAverageGrade()
	averge1 := s2.getAverageGrade()
	max := s1.getMaxGrade()
	max1 := s2.getMaxGrade()
	fmt.Println(averge, averge1)
	fmt.Println(max, max1)

}
