package entity

import (
	"context"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

type StartTestQuery struct {
	IdExam    int `json: "IdExam"`
	IdStudent int `json: "IdStudent"`
}
type Res struct {
	Status string `json: "status"`
}

// over
func PostStartTest(c *gin.Context) {
	var response StartTestQuery
	var err_r error_res
	var isExists bool
	if err := c.BindJSON(&response); err != nil {
		err_r.Err_s = "failed request"
		c.JSON(400, err_r)
		return
	}
	err := db.QueryRow(context.Background(), `select EXISTS (select * from StudentsResult where Exam=$1 and Student=$2)`, response.IdExam, response.IdStudent).Scan(&isExists)
	if err != nil {
		log.Print("db error")
		err_r.Err_s = "failed data from db1"
		c.JSON(400, err_r)
		return
	}
	//var statusExam int
	if isExists {
		var otvet Res
		otvet.Status = "Already started"
		c.JSON(200, otvet)
	} else {
		_, err := db.Exec(context.Background(), `insert INTO StudentsResult (Exam, Student, result, maximum, state) values ($1, $2, 0, 0, 1)`, response.IdExam, response.IdStudent)
		if err != nil {
			log.Print("db error")
			err_r.Err_s = "failed data from db2"
			c.JSON(400, err_r)
			return
		}
		var timemins int
		err2 := db.QueryRow(context.Background(), `select timemin from Exams where idExam=$1`, response.IdExam).Scan(&timemins)
		if err2 != nil {
			log.Print("db error")
			err_r.Err_s = "failed data from db12222"
			c.JSON(400, err_r)
			return
		}
		go updateState(timemins, response.IdStudent, response.IdExam)
		var otvet Res
		otvet.Status = "ok"
		c.JSON(200, otvet)
	}
}
func updateState(min int, idStudent int, idExam int) {
	time.Sleep(time.Duration(min) * time.Minute)
	_, err3 := db.Exec(context.Background(), `UPDATE StudentsResult SET state = 2 where Exam = $1 and Student = $2`, idExam, idStudent)
	if err3 != nil {
		log.Print("db error")
		return
	}
}

type selectStudent struct {
	ExemTitle    string  `json: "exemTitle"`
	ResultStruct float32 `json: "result"`
	ResultMax    float32 `json: "max"`
}
type ResponseStudent struct {
	Results []selectStudent
}

func GetResultsByStudent(c *gin.Context) {
	var err_r error_res
	requestBody := c.Param("id")
	log.Print(requestBody)
	rows, err := db.Query(context.Background(), `select Exams.title, result, maximum from StudentsResult
	LEFT JOIN Exams ON idExam = Exam where Student=$1`, requestBody)
	if err != nil {
		log.Print("db error")
		err_r.Err_s = "failed data from db3"
		c.JSON(400, err_r)
		return
	}
	var exemTitle string
	var resultStruct float32
	var resultMax float32
	var structForReturn selectStudent
	var num_rows ResponseStudent
	log.Println("going")
	for rows.Next() {
		err := rows.Scan(&exemTitle, &resultStruct, &resultMax)
		if err != nil {
			log.Print("1")
		}
		structForReturn.ExemTitle = exemTitle
		structForReturn.ResultStruct = resultStruct
		structForReturn.ResultMax = resultMax
		num_rows.Results = append(num_rows.Results, structForReturn)

	}
	rows.Close()
	c.JSON(200, num_rows)
}

type selectExam struct {
	Name         string  `json: "name"`
	SecondName   string  `json: "secondName"`
	ResultStruct float32 `json: "result"`
	ResultMax    float32 `json: "max"`
}
type ResponseExam struct {
	Results []selectExam
}

// over
func GetResultsBySubject(c *gin.Context) {
	var err_r error_res
	requestBody := c.Param("id")
	log.Print(requestBody)
	rows, err := db.Query(context.Background(), `select Students.name, Students.secondName, result, maximum from StudentsResult
	LEFT JOIN Students ON idStudent = Student where Exam=$1`, requestBody)
	if err != nil {
		log.Print("db error")
		err_r.Err_s = "failed data from db4"
		c.JSON(400, err_r)
		return
	}
	var namef string
	var nameSecf string
	var resultStruct float32
	var resultMax float32
	var structForReturn selectExam
	var num_rows ResponseExam
	for rows.Next() {
		err := rows.Scan(&namef, &nameSecf, &resultStruct, &resultMax)
		if err != nil {
			log.Print("1")
		}
		structForReturn.Name = namef
		structForReturn.SecondName = nameSecf
		structForReturn.ResultStruct = resultStruct
		structForReturn.ResultMax = resultMax
		num_rows.Results = append(num_rows.Results, structForReturn)

	}
	rows.Close()
	c.JSON(200, num_rows)
}

// over
func GetStatusTest(c *gin.Context) {
	var response Res
	var err_r error_res
	requestBody := c.Param("idStudent")
	requestBody2 := c.Param("idSubject")
	var state int
	log.Print(requestBody)
	err := db.QueryRow(context.Background(), `select state from StudentsResult where Exam=$1 and Student=$2`, requestBody2, requestBody).Scan(&state)
	if err != nil {
		log.Print("db error")
		err_r.Err_s = "failed data from db5"
		c.JSON(400, err_r)
		return
	}
	if state == 1 {
		response.Status = "started"
	}
	if state == 2 {
		response.Status = "finished"
	}
	c.JSON(200, response)
}

type Postquery struct {
	IdAnswer   int `json: "answer"`
	IdQuestion int `json: "question"`
}

type PostStuct struct {
	Student int         `json: "student"`
	Exam    int         `json: "exam"`
	Report  []Postquery `json: "Report"`
}

type PostOtvet struct {
	Result float32 `json: "Result"`
	Max    float32 `json: "Max"`
}

// over
func PostAnswerTest(c *gin.Context) {
	var response PostStuct
	var err_r error_res
	var isExists bool

	if err := c.BindJSON(&response); err != nil {
		err_r.Err_s = "failed request"
		c.JSON(400, err_r)
		return
	}
	err := db.QueryRow(context.Background(), `select EXISTS (select * from StudentsResult where Exam=$1 and Student=$2)`, response.Exam, response.Student).Scan(&isExists)
	if err != nil {
		log.Print("db error")
		err_r.Err_s = "failed data from db6"
		c.JSON(400, err_r)
		return
	}
	var statusExam int
	if !isExists {
		var otvet Res
		otvet.Status = "don't find exam"
		c.JSON(200, otvet)
		return
	} else {
		err := db.QueryRow(context.Background(), `select state from StudentsResult where Exam=$1 and Student=$2`, response.Exam, response.Student).Scan(&statusExam)
		if err != nil {
			log.Print("db error")
			err_r.Err_s = "failed data from db7"
			c.JSON(400, err_r)
			return
		}
		if statusExam == 2 {
			var otvet Res
			otvet.Status = "alredy finished"
			c.JSON(200, otvet)
			return
		}
	}
	resultExam := false
	var resultOtvet float32
	var resultMax float32
	var tempResult float32
	log.Println(response.Report)
	for k, value := range response.Report {
		err := db.QueryRow(context.Background(), `select istrue from Answers where idAnswer = $1`, response.Report[k].IdAnswer).Scan(&resultExam)
		if err != nil {
			log.Print("db error")
			log.Print(response.Report[k])
			log.Print(resultExam)
			err_r.Err_s = "failed data from db8"
			c.JSON(400, err_r)
			return
		}
		err2 := db.QueryRow(context.Background(), `select ball from Questions where idQuestion=$1`, value.IdQuestion).Scan(&tempResult)
		if err2 != nil {
			log.Print("db error")
			err_r.Err_s = "failed data from db9"
			c.JSON(400, err_r)
			return
		}
		resultMax += tempResult
		if resultExam {
			resultOtvet += tempResult
		}
	}
	_, err3 := db.Exec(context.Background(), `UPDATE StudentsResult SET result = $1, maximum = $2, state = 2 where Exam = $3 and Student = $4`, resultOtvet, resultMax, response.Exam, response.Student)
	if err3 != nil {
		log.Print("db error")
		err_r.Err_s = "failed data from db10"
		c.JSON(400, err_r)
		return
	}
	var finalOtvet PostOtvet
	finalOtvet.Max = resultMax
	finalOtvet.Result = resultOtvet
	c.JSON(200, finalOtvet)
	return
}
