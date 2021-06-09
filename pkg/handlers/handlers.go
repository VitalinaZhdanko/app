package handlers

import (
	"app/diplom/pkg/auth"
	"app/diplom/pkg/controllers"
	"app/diplom/pkg/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

// HealthCheck view
func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, "ok")
}

// Registration handler for Gin router
func Registration(c *gin.Context) {
	user := new(models.User)
	err := user.PopulateFromRequest(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err)
	}
	if !user.IsValid() {
		c.JSON(http.StatusUnprocessableEntity, "Invalid user data")
		return
	}

	err = auth.RegisterUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "User did not register")
		return
	}

	response := user.PrepareResponse()

	c.JSON(http.StatusOK, response)
}

// Login handler for Gin router
func Login(c *gin.Context) {
	user := new(models.User)
	err := user.PopulateFromRequest(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err)
	}
	if !user.IsValid() {
		c.JSON(http.StatusUnprocessableEntity, "Invalid user data")
		return
	}
	response, err := auth.LoginUser(user)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "Unauthorized")
		return
	}

	c.JSON(http.StatusOK, response)
}

// AuthCheck needed to test authentication
func AuthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, "secure endpoint")
}

// GetAllLanguage returns all tasks list
func GetAllLanguage(c *gin.Context) {
	allLanguage, err := controllers.GetAllLanguage()
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Error at Marshall data")
		log.Println(err)
		return
	}
	c.JSON(http.StatusOK, allLanguage)
}

func GetAllTask(c *gin.Context) {
	tasks, err := controllers.GetAllTask()
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Error at Marshall data")
		log.Println(err)
		return
	}
	c.JSON(http.StatusOK, tasks)
}


func GetCourseByLanguageID(c *gin.Context) {
	languageIDVal := c.Param("language_id")
	languageID, err := strconv.Atoi(languageIDVal)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Wrong ID by language")
		return
	}
	course, err := controllers.GetCourseByLanguageID(languageID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Error at Marshall data")
		log.Println(err)
		return
	}
	c.JSON(http.StatusOK, course)
}

func GetSectionsByCourseID(c *gin.Context) {
	courseIDVal := c.Param("course_id")
	courseID, err := strconv.Atoi(courseIDVal)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Wrong ID by course")
		return
	}

	courses, err := controllers.GetSectionsByCourseID(courseID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Error at Marshall data")
		log.Println(err)
		return
	}
	fmt.Println(courses)
	c.JSON(http.StatusOK, courses)
}

func GetLessonsBySectionID(c *gin.Context){
	sectionIDVal := c.Param("section_id")
	sectionID, err := strconv.Atoi(sectionIDVal)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Wrong ID by section")
		return
	}
	lessons, err := controllers.GetLessonsBySectionID(sectionID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Error at Marshall data")
		log.Println(err)
		return
	}
	c.JSON(http.StatusOK, lessons)
}

func GetTestByLessonID(c *gin.Context){
	lessonIDVal := c.Param("lesson_id")
	lessonID, err := strconv.Atoi(lessonIDVal)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Wrong ID by section")
		return
	}
	test, err := controllers.GetTestByLessonID(lessonID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Error at Marshall data")
		log.Println(err)
		return
	}
	c.JSON(http.StatusOK, test)
}

func GetQuestionAndAnswerByTestID(c *gin.Context){
	testIDVal := c.Param("test_id")
	testID, err := strconv.Atoi(testIDVal)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Wrong ID by test")
		return
	}
	questionAndAnswer, err := controllers.GetQuestionAndAnswerByTestID(testID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Error at Marshall data")
		log.Println(err)
		return
	}
	c.JSON(http.StatusOK, questionAndAnswer)
}

func GetUserInfoByUser(c *gin.Context){
	user := c.Param("user")
	userInfo, err := controllers.GetUserInfoByUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Error at Marshall data")
		log.Println(err)
		return
	}
	c.JSON(http.StatusOK, userInfo)
}

func ChangePassword(c *gin.Context) {
	userIDVal := c.Param("user_id")
	userID, err := strconv.Atoi(userIDVal)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Wrong ID by test")
		return
	}
	user := new(models.User)
	errPop := user.PopulateFromRequest(c.Request.Body)
	if errPop != nil {
		c.JSON(http.StatusUnprocessableEntity, err)
	}
	user.ID = userID
	errCon := controllers.ChangePassword(user)
	if errCon != nil {
		c.JSON(http.StatusUnauthorized, "Unauthorized")
		return
	}
	c.JSON(http.StatusOK, "Changed password")
}

func WriteOnCourse(c *gin.Context){
	userIDVal := c.Param("user_id")
	userID, err := strconv.Atoi(userIDVal)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Wrong ID by user")
		return
	}
	courseIDVal := c.Param("course_id")
	courseID, err := strconv.Atoi(courseIDVal)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Wrong ID by course")
		return
	}
	message, err := controllers.WriteOnCourse(userID, courseID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Error at Marshall data")
		log.Println(err)
		return
	}
	c.JSON(http.StatusOK, message)
}

func ChangeCourse(c *gin.Context){
	userIDVal := c.Param("user_id")
	userID, err := strconv.Atoi(userIDVal)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Wrong ID by user")
		return
	}
	courseIDVal := c.Param("course_id")
	courseID, err := strconv.Atoi(courseIDVal)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Wrong ID by course")
		return
	}

	message, err := controllers.ChangeCourse(userID, courseID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Error at Marshall data")
		log.Println(err)
		return
	}
	c.JSON(http.StatusOK, message)
}

func GetCourseIDByUserID(c *gin.Context){
	userIDVal := c.Param("user_id")
	userID, err := strconv.Atoi(userIDVal)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Wrong ID by user")
		return
	}
	courseID, err := controllers.GetCourseIDByUserID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Error at Marshall data")
		log.Println(err)
		return
	}
	c.JSON(http.StatusOK, courseID)
}

func GetTasksByLessonID(c *gin.Context){
	lessonIDVal := c.Param("lesson_id")
	lessonID, err := strconv.Atoi(lessonIDVal)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Wrong ID by section")
		return
	}

	tasks, err := controllers.GetTasksByLessonID(lessonID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Error at Marshall data")
		log.Println(err)
		return
	}
	c.JSON(http.StatusOK, tasks)
}

// Login handler for Gin router
func LexCheckByTaskID(c *gin.Context) {
	taskDecision := new(models.TaskDecision)
	errPop := taskDecision.PopulateFromRequest(c.Request.Body)
	if errPop != nil {
		c.JSON(http.StatusUnprocessableEntity, errPop)
	}
	answerTask, errCon := controllers.LexCheckByTaskID(taskDecision.TaskDecision)
	if errCon != nil {
		c.JSON(http.StatusInternalServerError, "Error at Marshall data")
		return
	}
	c.JSON(http.StatusOK, answerTask)
}

func SyntaxCheck(c *gin.Context) {
	taskDecision := new(models.TaskDecision)
	errPop := taskDecision.PopulateFromRequest(c.Request.Body)
	if errPop != nil {
		c.JSON(http.StatusUnprocessableEntity, errPop)
	}
	fmt.Println(taskDecision)
	answerTask, errCon := controllers.SyntaxCheck(taskDecision.TaskDecision)
	if errCon != nil {
		c.JSON(http.StatusInternalServerError, "Error at Marshall data")
		return
	}

	c.JSON(http.StatusOK, answerTask)
}


func GetLessonByID(c *gin.Context){
	lessonIDVal := c.Param("lesson_id")
	lessonID, err := strconv.Atoi(lessonIDVal)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Wrong ID by section")
		return
	}

	test, err := controllers.GetLessonByID(lessonID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Error at Marshall data")
		log.Println(err)
		return
	}
	c.JSON(http.StatusOK, test)
}

// Login handler for Gin router
func RunTask(c *gin.Context) {
	taskDecision := new(models.TaskDecision)
	errPop := taskDecision.PopulateFromRequest(c.Request.Body)
	if errPop != nil {
		c.JSON(http.StatusUnprocessableEntity, errPop)
	}
	answerTask, errCon := controllers.RunTask(taskDecision.TaskDecision)
	if errCon != nil {
		c.JSON(http.StatusInternalServerError, "Error at Marshall data")
		return
	}
	c.JSON(http.StatusOK, answerTask)
}