package routers

import (
	"tests/internal/entity"

	"github.com/gin-gonic/gin"
)

func Routes(route *gin.Engine) {
	route.GET("/api/student/:id", entity.GetStudent)
	route.POST("/api/student", entity.AddStudent)
	route.DELETE("/api/student", entity.DelStudent)

	route.GET("/api/answers/:id", entity.GetAnswer)
	route.POST("/api/answers", entity.AddAnswer)
	route.DELETE("/api/answers", entity.DelAnswer)

	route.GET("/api/exam/:id", entity.GetExam)
	route.POST("/api/exam", entity.AddExam)
	route.DELETE("/api/exam", entity.DelExam)

	route.GET("/api/question/:id", entity.GetQuestion)
	route.POST("/api/question", entity.AddQuestion)
	route.DELETE("/api/question", entity.DelQuestion)

	route.GET("/api/question_an/:id", entity.GetQuestionA)
	route.POST("/api/question_an", entity.AddQuestionA)
	route.DELETE("/api/question_an", entity.DelQuestionA)

	route.GET("/api/question_ex/:id", entity.GetQuestionE)
	route.POST("/api/question_ex", entity.AddQuestionE)
	route.DELETE("/api/question_ex", entity.DelQuestionE)

	route.GET("/api/result/:id", entity.GetResult)
	route.POST("/api/result", entity.AddResult)
	route.DELETE("/api/result", entity.DelResult)

	route.GET("/api/result/student/:id", entity.GetResultsByStudent)
	route.POST("/api/result/start", entity.PostStartTest)
	route.GET("/api/result/exam/:id", entity.GetResultsBySubject)
	route.GET("/api/result/status/:idStudent/:idSubject", entity.GetStatusTest)
	route.POST("/api/result/finish", entity.PostAnswerTest)
}
