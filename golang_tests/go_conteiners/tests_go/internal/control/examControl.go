package entity

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
)

type DExam struct {
	Id int `json:"id"`
}

type Exam struct {
	IdExam  int    `json:"IdExam"`
	Title   string `json:"Title"`
	Timemin int    `json:"Timemin"`
}

func GetExam(c *gin.Context) {
	var response Exam
	var err_r error_res
	requestBody := c.Param("id")
	log.Print(requestBody)
	err := db.QueryRow(context.Background(), `select idExam, title, timemin from Exams where idExam=$1`, requestBody).Scan(&response.IdExam, &response.Title, &response.Timemin)
	if err != nil {
		log.Print("db error")
		err_r.Err_s = "failed data from db"
		c.JSON(400, err_r)
		return
	}
	c.JSON(200, response)
}

func AddExam(c *gin.Context) {
	var response Exam
	var err_r error_res
	if err := c.BindJSON(&response); err != nil {
		err_r.Err_s = "failed request"
		c.JSON(400, err_r)
		return
	}
	_, err := db.Exec(context.Background(), "INSERT INTO Exams (title, timemin) values ($1, $2)", response.Title, response.Timemin)
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

func DelExam(c *gin.Context) {
	var response DExam
	var err_r error_res
	if err := c.BindJSON(&response); err != nil {
		err_r.Err_s = "failed request"
		c.JSON(400, err_r)
		return
	}
	_, err := db.Exec(context.Background(), "delete from Exams where idExam=$1", response.Id)
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
