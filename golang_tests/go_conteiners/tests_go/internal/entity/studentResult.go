package entity

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
)

type DResult struct {
	Id int `json:"id"`
}

type Result struct {
	IdSr    int     `json:"IdSr"`
	Exam    int     `json:"Exam"`
	Student int     `json:"Student"`
	Result  float32 `json:"Result"`
	Maximum float32 `json:"Maximum"`
	State   int     `json:"State"`
}

func GetResult(c *gin.Context) {
	var response Result
	var err_r error_res
	requestBody := c.Param("id")
	log.Print(requestBody)
	err := db.QueryRow(context.Background(), `select idSr, Exam, Student, result, maximum, state from StudentsResult where idSr=$1`, requestBody).Scan(&response.IdSr, &response.Exam, &response.Student, &response.Result, &response.Maximum, &response.State)
	if err != nil {
		log.Print("db error")
		err_r.Err_s = "failed data from db"
		c.JSON(400, err_r)
		return
	}
	c.JSON(200, response)
}

func AddResult(c *gin.Context) {
	var response Result
	var err_r error_res
	if err := c.BindJSON(&response); err != nil {
		err_r.Err_s = "failed request"
		c.JSON(400, err_r)
		return
	}
	_, err := db.Exec(context.Background(), "INSERT INTO StudentsResult (Exam, Student, result, maximum, state) values ($1, $2, $3, $4, $5)", response.Exam, response.Student, response.Result, response.Maximum, response.State)
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

func DelResult(c *gin.Context) {
	var response DResult
	var err_r error_res
	if err := c.BindJSON(&response); err != nil {
		err_r.Err_s = "failed request"
		c.JSON(400, err_r)
		return
	}
	_, err := db.Exec(context.Background(), "delete from StudentsResult where idSr=$1", response.Id)
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
