package services

import (
	"courses/model"
	"courses/repository"
)

type CourseService interface {
	GetCourses(studentId string) ([]model.Course, error)
	SignUpCourse(data model.SignUpInput) (bool, error)
	CancelCourse(courseId int, studentId string) (bool, error)
	GetClassMateCourses(studentId string) ([]model.Course, error)
}
type courseService struct {
}

var repo repository.CourseRepository

func NewCourseService(courseRepo repository.CourseRepository) CourseService {
	repo = courseRepo
	return &courseService{}
}

func (*courseService) GetCourses(studentId string) ([]model.Course, error) {
	data, err := repo.GetCourses(studentId)
	return data, err
}
func (*courseService) SignUpCourse(data model.SignUpInput) (bool, error) {
	output, err := repo.SignUpCourse(data)
	return output, err
}
func (*courseService) CancelCourse(courseId int, studentId string) (bool, error) {
	output, err := repo.CancelCourse(courseId, studentId)
	return output, err
}
func (*courseService) GetClassMateCourses(studentId string) ([]model.Course, error) {
	data, err := repo.GetClassMateCourses(studentId)
	return data, err
}