package repository

import (
	"fmt"
	"courses/connections"
	"courses/model"
	 "reflect"
)

type CourseRepository interface {
	GetCourses(studentId string) ([]model.Course, error)
	SignUpCourse(model.SignUpInput) (bool, error)
	CancelCourse(courseId int, studentId string) (bool, error)
	GetClassMateCourses(studentId string) ([]model.Course, error)
}

type course struct {
	Db connections.DbConnection
}

var mysql = connections.New()

func NewCourseRepository() CourseRepository {
	return &course{Db: mysql}
}

func (c *course) GetCourses(studentId string) ([]model.Course, error) {

	var output []model.Course
	var rows, err = c.Db.Query(`SELECT e.email,c.name FROM courses.student_enrollment e inner join courses.mas_courses c on e.id = c.id where email = ?`,studentId)

	for rows.Next() {
		var x model.Course

		err_Rows := rows.Scan(&x.Email, &x.Name)

		if err_Rows != nil {
			fmt.Println(err_Rows)
		}

		output = append(output, x)
	}

	if err == nil {
		return output, nil
	} else {
		return nil, err
	}
}


func (m *course) SignUpCourse(i model.SignUpInput) (bool, error) {

	query := `INSERT INTO courses.student_enrollment VALUES (?,?,?,?)`

	_, err := m.Db.Query(query, i.Email, i.Id, i.SignupDate, "active")

	if err == nil {
		return true, err
	} else {
		return false, err
	}
}
func (m *course) CancelCourse(courseId int, studentId string) (bool, error) {

	_, err := m.Db.Query(`DELETE FROM courses.student_enrollment WHERE id=? and email = ?`, courseId, studentId)

	if err == nil {
		return true, err
	} else {
		return false, err
	}
}
func (c *course) GetClassMateCourses(studentId string) ([]model.Course, error) {
    fmt.Println(reflect.TypeOf(studentId))
	//fmt.println(type(studentId))
	var output []model.Course
	var rows, err = c.Db.Query(`SELECT e.email,c.name FROM courses.student_enrollment e inner join courses.mas_courses c on e.id = c.id and e.email <> ?`,studentId)

	for rows.Next() {
		var x model.Course

		err_Rows := rows.Scan(&x.Email, &x.Name)

		if err_Rows != nil {
			fmt.Println(err_Rows)
		}

		output = append(output, x)
	}

	if err == nil {
		return output, nil
	} else {
		return nil, err
	}
}