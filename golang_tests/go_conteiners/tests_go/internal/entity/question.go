package entity

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
)

type DQuestion struct {
	Id int `json:"id"`
}

type Question struct {
	IdQuestion int     `json:"IdQuestion"`
	Body       string  `json:"Body"`
	Ball       float32 `json:"Ball"`
}

func GetQuestion(c *gin.Context) {
	var response Question
	var err_r error_res
	requestBody := c.Param("id")
	log.Print(requestBody)
	err := db.QueryRow(context.Background(), `select idQuestion, body, ball from Questions where idQuestion=$1`, requestBody).Scan(&response.IdQuestion, &response.Body, &response.Ball)
	if err != nil {
		log.Print("db error")
		err_r.Err_s = "failed data from db"
		c.JSON(400, err_r)
		return
	}
	c.JSON(200, response)
}

func AddQuestion(c *gin.Context) {
	var response Question
	var err_r error_res
	if err := c.BindJSON(&response); err != nil {
		err_r.Err_s = "failed request"
		c.JSON(400, err_r)
		return
	}
	_, err := db.Exec(context.Background(), "INSERT INTO Questions (body, ball) values ($1, $2)", response.Body, response.Ball)
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

func DelQuestion(c *gin.Context) {
	var response DQuestion
	var err_r error_res
	if err := c.BindJSON(&response); err != nil {
		err_r.Err_s = "failed request"
		c.JSON(400, err_r)
		return
	}
	_, err := db.Exec(context.Background(), "delete from Questions where idQuestion=$1", response.Id)
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
