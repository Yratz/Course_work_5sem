package entity

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
)

type DStudent struct {
	Id int `json:"id"`
}

type GetStudentStruct struct {
	Name       string `json:"name"`
	SecondName string `json:"secondname"`
}
type Student struct {
	Name       string `json:"name"`
	SecondName string `json:"secondname"`
}

func GetStudent(c *gin.Context) {
	var response GetStudentStruct
	var err_r error_res
	requestBody := c.Param("id")
	log.Print(requestBody)
	err := db.QueryRow(context.Background(), `select name, secondName from Students where idStudent=$1`, requestBody).Scan(&response.Name, &response.SecondName)
	if err != nil {
		log.Print("db error")
		err_r.Err_s = "failed data from db"
		c.JSON(400, err_r)
		return
	}
	c.JSON(200, response)
}

func AddStudent(c *gin.Context) {
	var response Student
	var err_r error_res
	if err := c.BindJSON(&response); err != nil {
		err_r.Err_s = "failed request"
		c.JSON(400, err_r)
		return
	}
	_, err := db.Exec(context.Background(), "INSERT INTO Students (name, secondName) values ($1, $2)", response.Name, response.SecondName)
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

func DelStudent(c *gin.Context) {
	var response DStudent
	var err_r error_res
	if err := c.BindJSON(&response); err != nil {
		err_r.Err_s = "failed request"
		c.JSON(400, err_r)
		return
	}
	_, err := db.Exec(context.Background(), "delete from Students where idStudent=$1", response.Id)
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
