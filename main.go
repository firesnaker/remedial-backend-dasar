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
	Grades    map[int]int // Use course ID as the key
}

//{[{1 Jane Doe [{1 Biology} {2 Chemistry}] map[1:85 2:90]}

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
	for z := 0; z < len(manager.Students); z++ {
		if id == manager.Students[z].ID {
			return &manager.Students[z]
		}
	}
	return nil
}

func (manager InMemoryStudentManager) ViewStudentDetails(student *Student) string {

	fmt.Println(student.FirstName)
	return "Student ID: " + strconv.Itoa(student.ID) + "\nFirst Name: " + student.FirstName + "\nLast Name: " + student.LastName
	// student ID : 1\nFirst Name: John\nLast Name : Doe
}

func (manager InMemoryStudentManager) ViewCourseDetails(student *Student) string {


	var detailCourse string = "Courses:"
	for z := 0; z < len(student.Courses); z++ {
		detailCourse += "\nCourse ID: " + strconv.Itoa(student.Courses[z].ID) + "\nCourse Name: " + student.Courses[z].Name
	}
	return detailCourse
}

func (manager InMemoryStudentManager) ViewGrades(student *Student) string {
	var grades string = "Grades:"
	for key, val := range student.Grades {
		fmt.Println(key, " \t:", val)
		var courseName string
		for z := 0; z < len(student.Courses); z++ {
			if student.Courses[z].ID == key {
				courseName = student.Courses[z].Name

			}
		}
		grades += "\nCourse: " + courseName + "\nGrade: " + strconv.Itoa(val)
	}
	return grades

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
			//coba tambah
			{
				ID:        2,
				FirstName: "Johnny",
				LastName:  "Suh",
				Courses:   []Course{{ID: 1, Name: "Mathematics"}, {ID: 2, Name: "History"}},
				Grades:    map[int]int{1: 95, 2: 80}, // Use course ID as the key
			},
		},
	}

	// Assume a student logs in with ID 1
	student := manager.Login(1)
	fmt.Println(student)
	if student != nil {
		fmt.Println(manager.ViewStudentDetails(student))
		fmt.Println(manager.ViewCourseDetails(student))
		fmt.Println(manager.ViewGrades(student))
	} else {
		fmt.Println("Student not found!")
	}
}
