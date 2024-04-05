package main

import (
	"fmt"
	"strconv"
)

type Student struct {
	ID        int
	FirstName string
	LastName  string
	Courses   []Course
	Grades    map[int]int
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
	for i := 0; i < len(manager.Students); i++ {
		if id == manager.Students[i].ID {
			return &manager.Students[i]
		}
	}

	return nil

}

func (manager InMemoryStudentManager) ViewStudentDetails(student *Student) string {
	return "Student ID: " + strconv.Itoa(student.ID) + "\nFirst Name: " + student.FirstName + "\nLast Name: " + student.LastName
}

func (manager InMemoryStudentManager) ViewCourseDetails(student *Student) string {
	var details string = "Courses: "
	for i := 0; i < len(student.Courses); i++ {
		details += "\nCourse ID: " + strconv.Itoa(student.Courses[i].ID) + "\nCourse Name: " + student.Courses[i].Name
	}
	return details
}

func (manager InMemoryStudentManager) ViewGrades(student *Student) string {
	var details string = "Grades:"
	for key, val := range student.Grades {
		var courseName string
		for i := 0; i < len(student.Courses); i++ {
			if student.Courses[i].ID == key {
				courseName = student.Courses[i].Name
			}
		}
		details += "\nCourse: " + courseName + "\nGrade: " + strconv.Itoa(val)
	}
	return details
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
			{
				ID:        2,
				FirstName: "John",
				LastName:  "Doe",
				Courses:   []Course{{ID: 1, Name: "Math"}, {ID: 2, Name: "Physics"}},
				Grades:    map[int]int{1: 80, 2: 85}, // Use course ID as the key
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
