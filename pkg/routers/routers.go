package routers

import (
	"app/diplom/pkg/handlers"
	"app/diplom/pkg/middleware"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

// SetupRouter - all endpoints
func SetupRouter() *gin.Engine {
	gin.ForceConsoleColor()
	router := gin.New()
	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d \n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
		)
	}))
	router.Use(gin.Recovery())

	v1 := router.Group("/api/v1")

	v1.GET("/", handlers.HealthCheck)
	v1.POST("/registration", handlers.Registration)
	v1.POST("/login", handlers.Login)

	v1.Use(middleware.Auth()).GET("/auth", handlers.AuthCheck)
	v1.Use(middleware.Auth()).GET("/languages", handlers.GetAllLanguage)
	v1.Use(middleware.Auth()).GET("/course/:language_id", handlers.GetCourseByLanguageID)
	v1.Use(middleware.Auth()).GET("/sections/:course_id", handlers.GetSectionsByCourseID)
	v1.Use(middleware.Auth()).GET("/lessons/:section_id", handlers.GetLessonsBySectionID)
	v1.Use(middleware.Auth()).GET("/test/:lesson_id", handlers.GetTestByLessonID)
	v1.Use(middleware.Auth()).GET("/questionAndAnswer/:test_id", handlers.GetQuestionAndAnswerByTestID)
	v1.Use(middleware.Auth()).GET("/userInfo/:user", handlers.GetUserInfoByUser)
	v1.Use(middleware.Auth()).POST("/changePassword/:user_id", handlers.ChangePassword)
	v1.Use(middleware.Auth()).GET("/writeOnCourse/:user_id/:course_id", handlers.WriteOnCourse)
	v1.Use(middleware.Auth()).GET("/changeCourse/:user_id/:course_id", handlers.ChangeCourse)
	v1.Use(middleware.Auth()).GET("/courseID/:user_id", handlers.GetCourseIDByUserID)
	v1.Use(middleware.Auth()).GET("/tasks/:lesson_id", handlers.GetTasksByLessonID)
	v1.Use(middleware.Auth()).POST("/lexCheck", handlers.LexCheckByTaskID)
	v1.Use(middleware.Auth()).POST("/syntaxCheck", handlers.SyntaxCheck)
	v1.Use(middleware.Auth()).GET("/lesson/:lesson_id", handlers.GetLessonByID)
	v1.Use(middleware.Auth()).GET("/tasks", handlers.GetAllTask)
	v1.Use(middleware.Auth()).POST("/runTask", handlers.RunTask)

	return router
}

