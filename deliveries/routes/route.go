package routes

import (
	"instructors-api/deliveries/controllers/auth"
	"instructors-api/deliveries/controllers/instructor"
	"instructors-api/deliveries/controllers/course"
	"instructors-api/deliveries/middlewares"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RegisterPaths(e *echo.Echo, ac *auth.AuthController, ic *instructor.InstructorController, cc *course.CourseController) {
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	// Auth Route
	a := e.Group("/login")
	a.POST("", ac.Login())

	// Instructor Route
	i := e.Group("/instructors")
	i.POST("/register", ic.Create())
	i.PUT("/update-profile", ic.Update(), middlewares.JWTMiddleware())

	// Course Route
	c := e.Group("/courses")
	c.Use(middlewares.JWTMiddleware())
	c.POST("", cc.Create())
	c.GET("", cc.Get())
	c.GET("", cc.GetDetail())
	c.GET("/:id", cc.Update())
	c.DELETE("/:id", cc.Delete())
}
