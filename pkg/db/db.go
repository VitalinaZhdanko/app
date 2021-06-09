// Package db implements mock db functions
package db

import (
	"app/diplom/pkg/models"
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"time"
)

// dbC struct for connecting to the db
type dbC struct {
	pool *pgxpool.Pool
	conn *pgxpool.Conn
}

var db = dbC{}

// Connect function gets a db connection
func Connect(connectionString string) (*pgxpool.Conn, *pgxpool.Pool, error) {
	pool, err := pgxpool.Connect(context.Background(), connectionString)
	if err != nil {
		return nil, nil, err
	}
	log.Println("Connected")

	conn, err := pool.Acquire(context.Background())
	if err != nil {
		return nil, nil, err
	}
	log.Println("Get connection")

	db.pool = pool
	db.conn = conn
	return conn, pool, nil
}

// AddUser function adds into db
func AddUser(user *models.User) (int, error) {
	err := db.conn.QueryRow(context.Background(), "insert into student(login, password, fio, role_id) values($1, $2, $3, $4) returning id", user.Username, user.Password, user.FIO, 1).Scan(&user.ID)
	return user.ID, err
}

// GetUserByUsername returns User if it was found in db, and error if not
func GetUserByUsername(username string) (*models.User, error) {
	fmt.Println("username")
	fmt.Println(username)
	fmt.Println("db 1")
	row := db.conn.QueryRow(context.Background(), "select * from student where login = $1", username)
	fmt.Println("db 2")

	user := models.User{}
	fmt.Println("db 3")

	err := row.Scan(&user.ID, &user.Username, &user.Password, &user.FIO, &user.RoleID)
	fmt.Println("db 4")
	fmt.Println(user.Username + user.Password)
	fmt.Println("db 5")


	if err != nil {
		fmt.Println(err)

		return nil, err
	}
	return &user, nil
}

func GetAllLanguage()([]models.Language, error){
	rows, err := db.conn.Query(context.Background(), "select * from language")
	if err != nil {
		return nil, err
	}
	var languages []models.Language

	for rows.Next() {
		t := models.Language{}
		err = rows.Scan(&t.ID, &t.Name)
		if err != nil {
			return nil, err
		}
		languages= append(languages, t)
	}
	return languages, err
}

func GetAllTask()([]models.Task, error){
	rows, err := db.conn.Query(context.Background(), "select * from task")
	if err != nil {
		return nil, err
	}
	var tasks []models.Task

	for rows.Next() {
		t := models.Task{}
		err = rows.Scan(&t.ID, &t.Name, &t.Description, &t.Score, &t.LessonID)
		if err != nil {
			return nil, err
		}
		tasks= append(tasks, t)
	}
	return tasks, err
}

func GetCourseByLanguage(languageID int) (*models.Course, error) {
	row := db.conn.QueryRow(context.Background(), "select * from course where language_id = $1", languageID)
	course := models.Course{}

	err := row.Scan(&course.ID, &course.Name, &course.Description, &course.LanguageID)
	if err != nil {
		return nil, err
	}

	return &course, nil
}

func GetSectionsByCourseID(courseID int) ([]models.Section, error) {
	rows, err := db.conn.Query(context.Background(), "select * from section where course_id = $1", courseID)
	if err != nil {
		return nil, err
	}
	var sections []models.Section

	for rows.Next() {
		t := models.Section{}
		err = rows.Scan(&t.ID, &t.Name, &t.Description, &t.CourseID)
		if err != nil {
			return nil, err
		}
		sections = append(sections, t)
	}
	return sections, err
}

func GetLessonsBySectionID(sectionID int)([]models.Lesson, error){
	rows, err := db.conn.Query(context.Background(), "select * from lesson where section_id = $1 order by id asc", sectionID)
	if err != nil {
		return nil, err
	}
	var lessons []models.Lesson

	for rows.Next() {
		t := models.Lesson{}
		err = rows.Scan(&t.ID, &t.Name, &t.Description, &t.SectionID, &t.Youtube)
		if err != nil {
			return nil, err
		}
		lessons = append(lessons, t)
	}
	return lessons, err
}

func GetTestByLessonID(lessonID int) (*models.Test, error) {
	row := db.conn.QueryRow(context.Background(), "select * from test where lesson_id = $1", lessonID)
	test := models.Test{}

	err := row.Scan(&test.ID, &test.Name, &test.Description, &test.LessonID)
	if err != nil {
		return nil, err
	}

	return &test, nil
}

func GetQuestionAndAnswerByTestID(testID int) ([]models.QuestionAndAnswer, error) {
	var questionAndAnswer []models.QuestionAndAnswer
	var question []models.Question
	var answer []models.Answer

	rowsQuestion, errQ := db.conn.Query(context.Background(), "select * from question where test_id = $1", testID)
	if errQ != nil {
		return nil, errQ
	}

	for rowsQuestion.Next(){
		t := models.Question{}

		errQ = rowsQuestion.Scan(&t.ID, &t.Description, &t.Cost, &t.TestID)
		if errQ != nil {
			return nil, errQ
		}

		question = append(question, t)

	}

	for i := 0; i < len(question); i++ {
		rowsAnswer, errQ := db.conn.Query(context.Background(), "select * from answer where question_id = $1", question[i].ID)
		if errQ != nil {
			return nil, errQ
		}

		for rowsAnswer.Next(){
			t := models.Answer{}

			errQ = rowsAnswer.Scan(&t.ID, &t.Description, &t.IsTrue, &t.QuestionID)
			if errQ != nil {
				return nil, errQ
			}
			answer = append(answer, t)

		}

	}

	for i := 0; i < len(question); i++{

		t := models.QuestionAndAnswer{}
		var answerTwo []models.Answer


		t.TestID = question[i].TestID
		t.QuestionID = question[i].ID
		t.QuestionDescription = question[i].Description
		t.Cost = question[i].Cost

		for j := 0; j < len(answer); j++{
			if answer[j].QuestionID == question[i].ID {
				answerTwo = append(answerTwo, answer[j])
			}
		}

		t.Answer1 = answerTwo[0].Description
		t.Answer2 = answerTwo[1].Description
		t.Answer3 = answerTwo[2].Description

		t.IsTrue1 = answerTwo[0].IsTrue
		t.IsTrue2 = answerTwo[1].IsTrue
		t.IsTrue3 = answerTwo[2].IsTrue

		questionAndAnswer = append(questionAndAnswer, t)
	}

	return questionAndAnswer, errQ
}

func GetUserInfoByUserID(user string) (*models.UserInfo, error) {
	rowStudent := db.conn.QueryRow(context.Background(), "select id, login, fio from student where login = $1", user)

	userInfo := models.UserInfo{}
	scoreTest := 0
	scoreTask := 0

	err := rowStudent.Scan(&userInfo.ID, &userInfo.Login, &userInfo.FIO)
	if err != nil {
		return nil, err
	}

	userID := userInfo.ID
	rowScoreTest := db.conn.QueryRow(context.Background(), "select sum(score) from test_result where student_id = $1", userID)
	errTest := rowScoreTest.Scan(&scoreTest)
	if errTest != nil {
		scoreTest = 0
	}

	rowScoreTask := db.conn.QueryRow(context.Background(), "select sum(score) from task_result where student_id = $1", userID)
	errTask := rowScoreTask.Scan(&scoreTask)
	if errTask != nil {
		scoreTask = 0
	}
	fmt.Println(userInfo)

	userInfo.Score = scoreTask + scoreTest
	fmt.Println(userInfo)

	var courseID int
	rowCourseID := db.conn.QueryRow(context.Background(), "select course_id from student_course where student_id = $1", userID)
	errCourse := rowCourseID.Scan(&courseID)
	if errCourse != nil {
		courseID = 0
	}
	if courseID == 0{
		userInfo.Course = "Вы еще не записаны на курс"
	} else {
		rowCourse := db.conn.QueryRow(context.Background(), "select name from course where id = $1", courseID)
		errCourse := rowCourse.Scan(&userInfo.Course)
		if errCourse != nil {
			fmt.Println(errCourse)
		}
	}

	return &userInfo, nil
}

func ChangePassword(user *models.User) (err error) {
	_, err = db.conn.Exec(context.Background(),"update student set password = $1 where id = $2", user.Password, user.ID)
	if err != nil {
		fmt.Println(err)
	}
	return err
}

func WriteOnCourse(userID int, courseID int) (message string, err error) {
	var studCourse int
	var studentCourseID int

	row := db.conn.QueryRow(context.Background(), "select count(id) from student_course where student_id = $1 and course_id = $2", userID, courseID)

	errS := row.Scan(&studCourse)
	if errS != nil {
		return "error", errS
	}

	if studCourse == 1{
		return "Вы уже добавлены на этот курс", nil
	}

	rowC := db.conn.QueryRow(context.Background(), "select * from student_course where student_id = $1", userID)

	errC := rowC.Scan(&studCourse)
	if errC != nil {
		return "Вы записаны на другой курс", errS
	}

	dt := time.Now()

	err = db.conn.QueryRow(context.Background(),"insert into student_course(start_date, end_date, student_id, course_id) values($1,'01-01-2000', $2, $3) returning id", dt.Format("01-02-2006"), userID, courseID).Scan(&studentCourseID)
	if err != nil {
		fmt.Println(err)
	}

	return "Вы успешно добавлены на курс", err
}

func ChangeCourse(userID int, courseID int) (message string, err error) {
	var studentCourseID int

	_, err = db.conn.Exec(context.Background(),"delete from student_course where student_id = $1", userID)
	if err != nil {
		fmt.Println(err)
	}

	dt := time.Now()

	err = db.conn.QueryRow(context.Background(),"insert into student_course(start_date, end_date, student_id, course_id) values($1,'01-01-2000', $2, $3) returning id", dt.Format("01-02-2006"), userID, courseID).Scan(&studentCourseID)
	if err != nil {
		fmt.Println(err)
	}

	return "Вы успешно добавлены на курс", err
}

func GetCourseIDByUserID(userID int) (courseID int, err error) {
	row := db.conn.QueryRow(context.Background(), "select course_id from student_course where student_id = $1", userID)

	err = row.Scan(&courseID)
	if err != nil {
		return 0, nil
	}
	return courseID, err
}

func GetTasksByLessonID(lessonID int)([]models.Task, error){
	rows, err := db.conn.Query(context.Background(), "select * from task where lesson_id = $1", lessonID)
	if err != nil {
		return nil, err
	}
	var tasks []models.Task

	for rows.Next() {
		t := models.Task{}
		err = rows.Scan(&t.ID, &t.Name, &t.Description, &t.Score, &t.LessonID)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}
	return tasks, err
}

func GetLessonByID(lessonID int) (*models.Lesson, error) {
	row := db.conn.QueryRow(context.Background(), "select * from lesson where id = $1", lessonID)
	lesson := models.Lesson{}
	err := row.Scan(&lesson.ID, &lesson.Name, &lesson.Description, &lesson.SectionID, &lesson.Youtube)
	if err != nil {
		fmt.Println(err)
	}

	return &lesson, nil
}
//  "TODO "AddTask""
//func AddTask(task *models.Task) (int, error) {
//	err := db.conn.QueryRow(context.Background(), "insert into tasks(name, description, time_limit, memory) values($1, $2, $3, $4) returning id", task.Name, task.Description, task.TimeLimit, task.Memory).Scan(&task.ID)
//	return task.ID, err
//}
//
//// GetAllTasks returns all tasks list from db
//func GetAllTasks() ([]models.Task, error) {
//	rows, err := db.conn.Query(context.Background(), "select * from tasks")
//	if err != nil {
//		return nil, err
//	}
//	var tasks []models.Task
//
//	for rows.Next() {
//		t := models.Task{}
//		err = rows.Scan(&t.ID, &t.Name, &t.Description, &t.TimeLimit, &t.Memory)
//		if err != nil {
//			return nil, err
//		}
//		tasks = append(tasks, t)
//	}
//	return tasks, err
//}
//
//// GetTaskByID gets task by ID
//func GetTaskByID(id int) (*models.Task, error) {
//	row := db.conn.QueryRow(context.Background(), "select * from tasks where id = $1", id)
//	task := models.Task{}
//	err := row.Scan(&task.ID, &task.Name, &task.Description, &task.TimeLimit, &task.Memory)
//	if err != nil {
//		return nil, err
//	}
//	return &task, nil
//}
//
//// "TODO "AddTestCase""
////func AddTestCase(testCase *models.TestCase) (int, error) {
////	err := db.conn.QueryRow(context.Background(), "insert into test_cases(task_id, test_data, answer) values($1, $2, $3) returning id", testCase.TaskID, testCase.TestData, testCase.Answer).Scan(&testCase.ID)
////	return testCase.ID, err
////}
//
//// GetTestCasesByTaskID gets test cases by task id
//func GetTestCasesByTaskID(taskID int) ([]models.TestCase, error) {
//	var err error
//	rows, err := db.conn.Query(context.Background(), "select * from test_cases where task_id = $1", taskID)
//	if err != nil {
//		return nil, err
//	}
//	var testCases []models.TestCase
//	for rows.Next() {
//		testCase := models.TestCase{}
//		err = rows.Scan(&testCase.ID, &testCase.TaskID, &testCase.TestData, &testCase.Answer)
//		if err != nil {
//			return nil, err
//		}
//		testCases = append(testCases, testCase)
//	}
//
//	return testCases, err
//}
