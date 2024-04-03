package main_test

import (
	"testing"

	main "remedial-backend-dasar"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("StudentPortal", func() {
	var (
		manager *main.InMemoryStudentManager
	)

	BeforeEach(func() {
		manager = &main.InMemoryStudentManager{
			Students: []main.Student{
				{
					ID:        1,
					FirstName: "John",
					LastName:  "Doe",
					Courses:   []main.Course{{ID: 1, Name: "Mathematics"}, {ID: 2, Name: "Physics"}},
					Grades:    map[int]int{1: 85, 2: 90}, // Use course ID as the key
				},
			},
		}
	})

	Describe("Login", func() {
		Context("with valid student id", func() {
			It("should return a student", func() {
				student := manager.Login(1)
				Expect(student).To(Not(BeNil()))
				Expect(student.FirstName).To(Equal("John"))
			})
		})

		Context("with invalid student id", func() {
			It("should return nil", func() {
				student := manager.Login(99)
				Expect(student).To(BeNil())
			})
		})
	})

	Describe("ViewStudentDetails", func() {
		It("should return student details", func() {
			student := manager.Login(1)
			details := manager.ViewStudentDetails(student)
			expected := "Student ID: 1\nFirst Name: John\nLast Name: Doe"
			Expect(details).To(Equal(expected))
		})
	})

	Describe("ViewCourseDetails", func() {
		It("should return student course details", func() {
			student := manager.Login(1)
			details := manager.ViewCourseDetails(student)
			expected := "Courses:\nCourse ID: 1\nCourse Name: Mathematics\nCourse ID: 2\nCourse Name: Physics"
			Expect(details).To(Equal(expected))
		})
	})

	Describe("ViewGrades", func() {
		It("should return student grades", func() {
			student := manager.Login(1)
			details := manager.ViewGrades(student)
			expected := "Grades:\nCourse: Mathematics\nGrade: 85\nCourse: Physics\nGrade: 90"
			Expect(details).To(Equal(expected))
		})
	})
})

func TestStudentPortal(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "StudentPortal Suite")
}
