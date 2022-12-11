package entity

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
)

type DQuestionE struct {
	Id int `json:"id"`
}

type QuestionE struct {
	IdQe     int `json:"idQe"`
	Exam     int `json:"Exam"`
	Question int `json:"Question"`
}

func GetQuestionE(c *gin.Context) {
	var response QuestionE
	var err_r error_res
	requestBody := c.Param("id")
	log.Print(requestBody)
	err := db.QueryRow(context.Background(), `select idQe, Exam, Question from QuestionsExam where idQe=$1`, requestBody).Scan(&response.IdQe, &response.Exam, &response.Question)
	if err != nil {
		log.Print("db error")
		err_r.Err_s = "failed data from db"
		c.JSON(400, err_r)
		return
	}
	c.JSON(200, response)
}

func AddQuestionE(c *gin.Context) {
	var response QuestionE
	var err_r error_res
	if err := c.BindJSON(&response); err != nil {
		err_r.Err_s = "failed request"
		c.JSON(400, err_r)
		return
	}
	_, err := db.Exec(context.Background(), "INSERT INTO QuestionsExam (Exam, Question) values ($1, $2)", response.Exam, response.Question)
	if err != nil {
		log.Print("db error")
		err_r.Err_s = "failed data from db"
		c.JSON(400, err_r)
		return
	}
	var res Status
	res.Status = "ok"
	c.JSON(200, res)
}

func DelQuestionE(c *gin.Context) {
	var response DQuestionE
	var err_r error_res
	if err := c.BindJSON(&response); err != nil {
		err_r.Err_s = "failed request"
		c.JSON(400, err_r)
		return
	}
	_, err := db.Exec(context.Background(), "delete from QuestionsExam where idQe=$1", response.Id)
	if err != nil {
		log.Print("db error")
		err_r.Err_s = "failed data from db"
		c.JSON(400, err_r)
		return
	}
	var res Status
	res.Status = "ok"
	c.JSON(200, res)
}
