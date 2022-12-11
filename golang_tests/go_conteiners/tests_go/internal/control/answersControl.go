package entity

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
)

type DAnswer struct {
	Id int `json:"id"`
}

type Answer struct {
	IdAnswer int    `json:"idAnswer"`
	Variant  string `json:"variant"`
	IsTrue   bool   `json:"isTrue"`
}

func GetAnswer(c *gin.Context) {
	var response Answer
	var err_r error_res
	requestBody := c.Param("id")
	log.Print(requestBody)
	err := db.QueryRow(context.Background(), `select idAnswer, variants, isTrue from Answers where idAnswer=$1`, requestBody).Scan(&response.IdAnswer, &response.Variant, &response.IsTrue)
	if err != nil {
		log.Print("db error")
		err_r.Err_s = "failed data from db"
		c.JSON(400, err_r)
		return
	}
	c.JSON(200, response)
}

func AddAnswer(c *gin.Context) {
	var response Answer
	var err_r error_res
	if err := c.BindJSON(&response); err != nil {
		err_r.Err_s = "failed request"
		c.JSON(400, err_r)
		return
	}
	_, err := db.Exec(context.Background(), "INSERT INTO Answers (variants, isTrue) values ($1, $2)", response.Variant, response.IsTrue)
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

func DelAnswer(c *gin.Context) {
	var response DAnswer
	var err_r error_res
	if err := c.BindJSON(&response); err != nil {
		err_r.Err_s = "failed request"
		c.JSON(400, err_r)
		return
	}
	_, err := db.Exec(context.Background(), "delete from Answers where idanswer=$1", response.Id)
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
