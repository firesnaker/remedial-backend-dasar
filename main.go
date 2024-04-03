package main

import (
	"fmt"
)

type Student struct {
	ID        int
	FirstName string
	LastName  string
	Courses   []Course
	Grades    map[int]int // Use course ID as the key
}

type Course struct {
	ID   int
	Name string
}

type StudentManager interface {
	Login(id int) *Student
	ViewStudentDetails(student *Student) string
	ViewCourseDetails(student *Student) string
	ViewGrades(student *Student) string
}

type InMemoryStudentManager struct {
	Students []Student
}

func (manager InMemoryStudentManager) Login(id int) *Student {
	return &Student{}
}

func (manager InMemoryStudentManager) ViewStudentDetails(student *Student) string {
	return ""
}

func (manager InMemoryStudentManager) ViewCourseDetails(student *Student) string {
	return ""
}

func (manager InMemoryStudentManager) ViewGrades(student *Student) string {
	return ""
}

func main() {
	manager := &InMemoryStudentManager{
		Students: []Student{
			{
				ID:        1,
				FirstName: "Jane",
				LastName:  "Doe",
				Courses:   []Course{{ID: 1, Name: "Biology"}, {ID: 2, Name: "Chemistry"}},
				Grades:    map[int]int{1: 85, 2: 90}, // Use course ID as the key
			},
		},
	}

	// Assume a student logs in with ID 1
	student := manager.Login(1)
	if student != nil {
		fmt.Println(manager.ViewStudentDetails(student))
		fmt.Println(manager.ViewCourseDetails(student))
		fmt.Println(manager.ViewGrades(student))
	} else {
		fmt.Println("Student not found!")
	}
}
