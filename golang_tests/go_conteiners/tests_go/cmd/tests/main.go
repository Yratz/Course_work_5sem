package main

import (
	"context"
	"log"
	"tests/internal/entity"
	"tests/internal/routers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

func main() {
	router := gin.Default()
	routers.Routes(router)
	info_db := "postgres://mirea:mirea@database:5432/mirea"
	conn, err := pgx.Connect(context.Background(), info_db)
	if err != nil {
		log.Print("lol")
		//os.Exit(1)
	}
	entity.SetDb(conn)
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true

	router.Use(cors.New(config))
	defer conn.Close(context.Background())
	router.Run(":8080")
}
