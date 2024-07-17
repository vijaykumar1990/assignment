package main

import (
	"courses/controllers"
	router "courses/http"
	"courses/repository"
	"courses/services"
)

var echoRouter = router.NewEchoRouter()
var courseRepo = repository.NewCourseRepository()
var courseservice = services.NewCourseService(courseRepo)
var coursecontroller = controllers.NewCourseController(courseservice)

func main() {

	echoRouter.GET("/courses/:email", coursecontroller.GetCourses)
	echoRouter.POST("/courses/signup", coursecontroller.SignUpCourse)
	echoRouter.DELETE("/courses/cancel", coursecontroller.CancelCourse)
	echoRouter.GET("/courses/others", coursecontroller.GetClassMateCourses)
	echoRouter.SERVE(":8085")
}
