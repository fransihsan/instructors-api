package main

import (
	"fmt"
	"instructors-api/configs"
	_AuthController "instructors-api/deliveries/controllers/auth"
	_CourseController "instructors-api/deliveries/controllers/course"
	_InstructorController "instructors-api/deliveries/controllers/instructor"
	"instructors-api/deliveries/routes"
	_AuthRepo "instructors-api/repositories/auth"
	_CourseRepo "instructors-api/repositories/course"
	_InstructorRepo "instructors-api/repositories/instructor"
	"instructors-api/utils"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func main() {
	config := configs.GetConfig(false)
	db := utils.InitDB(config)

	authRepo := _AuthRepo.NewAuthRepository(db)
	instructorRepo := _InstructorRepo.NewInstructorRepository(db)
	courseRepo := _CourseRepo.NewCourseRepository(db)

	ac := _AuthController.NewAuthController(authRepo)
	ic := _InstructorController.NewInstructorController(instructorRepo)
	cc := _CourseController.NewCourseController(courseRepo)

	e := echo.New()

	routes.RegisterPaths(e, ac, ic, cc)

	log.Fatal(e.Start(fmt.Sprintf(":%d", config.PORT)))
}
