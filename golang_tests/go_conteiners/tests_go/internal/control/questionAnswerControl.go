package entity

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
)

type DQuestionA struct {
	Id int `json:"id"`
}

type QuestionA struct {
	IdQa     int `json:"idQa"`
	Answer   int `json:"Answer"`
	Question int `json:"Question"`
}

func GetQuestionA(c *gin.Context) {
	var response QuestionA
	var err_r error_res
	requestBody := c.Param("id")
	log.Print(requestBody)
	err := db.QueryRow(context.Background(), `select idQa, Answer, Question from QuestionsAnswers where idQa=$1`, requestBody).Scan(&response.IdQa, &response.Answer, &response.Question)
	if err != nil {
		log.Print("db error")
		err_r.Err_s = "failed data from db"
		c.JSON(400, err_r)
		return
	}
	c.JSON(200, response)
}

func AddQuestionA(c *gin.Context) {
	var response QuestionA
	var err_r error_res
	if err := c.BindJSON(&response); err != nil {
		err_r.Err_s = "failed request"
		c.JSON(400, err_r)
		return
	}
	_, err := db.Exec(context.Background(), "INSERT INTO QuestionsAnswers (Answer, Question) values ($1, $2)", response.Answer, response.Question)
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

func DelQuestionA(c *gin.Context) {
	var response DQuestionA
	var err_r error_res
	if err := c.BindJSON(&response); err != nil {
		err_r.Err_s = "failed request"
		c.JSON(400, err_r)
		return
	}
	_, err := db.Exec(context.Background(), "delete from QuestionsAnswers where idQa=$1", response.Id)
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
