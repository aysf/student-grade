package grades

import (
	"fmt"
	"sync"
)

type Students []Student

var (
	students      Students
	studentsMutex *sync.Mutex
)

func (s Students) GetByID(id int) (*Student, error) {
	for i := range s {
		if s[i].ID == id {
			return &s[i], nil
		}
	}
	return nil, fmt.Errorf("student with id %v not found", id)
}

type Student struct {
	ID        int
	FirstName string
	LastName  string
	Grades    []Grade
}

func (s Student) Average() float32 {
	var sum float32
	for _, v := range s.Grades {
		sum += v.Score
	}
	return sum / float32(len(s.Grades))
}

type Grade struct {
	Title string
	Type  GradeType
	Score float32
}

type GradeType string

const (
	HomeWork = GradeType("HomeWork")
	Test     = GradeType("Test")
	Quiz     = GradeType("Quiz")
)
