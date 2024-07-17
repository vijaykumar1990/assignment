package controllers

import (
	"courses/services"
	"courses/model"
	"net/http"
	"github.com/labstack/echo"
	"encoding/json"
	"fmt"
	"strconv"
	
)

type CourseController interface {
	GetCourses(ctx echo.Context) error
    SignUpCourse(ctx echo.Context) error
	CancelCourse(ctx echo.Context) error
	GetClassMateCourses(ctx echo.Context) error
}

type courseController struct {
}

var service services.CourseService

func NewCourseController(courseService services.CourseService) CourseController {
	service = courseService
	return &courseController{}
}

func (*courseController) GetCourses(ctx echo.Context) error {
	studentId := ctx.Param("email")

	if studentId == ""  {
		return ctx.String(http.StatusBadRequest, "Missing Param Student Email")
	}
	data, err := service.GetCourses(studentId)

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, string(err.Error()))
	}

	return ctx.JSON(http.StatusOK, data)
}

func (*courseController) SignUpCourse(ctx echo.Context) error {
	input := model.SignUpInput{}
	err := json.NewDecoder(ctx.Request().Body).Decode(&input)

	if err != nil {
		errMessage := fmt.Sprintf("Error: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, errMessage)
	}

	data, err := service.SignUpCourse(input)

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, string(err.Error()))
	}
 
	return ctx.JSON(http.StatusOK, data)
}
func (*courseController) CancelCourse(ctx echo.Context) error {

	courseId,err := strconv.Atoi(ctx.QueryParam("id"))
    studentId := ctx.QueryParam("email")
	data, err := service.CancelCourse(courseId, studentId)

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, string(err.Error()))
	}

	return ctx.JSON(http.StatusOK, data)
}

func (*courseController) GetClassMateCourses(ctx echo.Context) error {
	studentId := ctx.QueryParam("email")

	if studentId == ""  {
		return ctx.String(http.StatusBadRequest, "Missing Param Student Email")
	}
	data, err := service.GetClassMateCourses(studentId)

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, string(err.Error()))
	}

	return ctx.JSON(http.StatusOK, data)
}