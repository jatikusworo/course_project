package main

import (
	"course_project/configs"
	"course_project/internal/database"
	"course_project/internal/router"
	"course_project/internal/user"
	"course_project/pkg/db"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// load env
	cfg := configs.Load()
	fmt.Println("Running environment:", cfg.Env)

	// inisialisasi DB
	gormDB, err := db.NewGormDB(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("db init: %v", err)
	}
	log.Println("Connection DB Success")

	// hit migrate function
	err = database.Migrate(gormDB)
	if err != nil {
		log.Fatalf("db migrate: %v", err)
		return
	}
	log.Println("Migration OK")

	//create repo , service , handler
	repo := user.NewGormUserRepository(gormDB)
	userService := user.NewService(repo)
	userHandler := user.NewHandler(userService)
	log.Println("preparing service OK")

	log.Println("Preparing registering handlers")
	r := gin.Default()

	handlers := &router.Handlers{
		User: userHandler,
		//Room: roomHandler,
	}

	router.RegisterRoutes(r, handlers)

	address := fmt.Sprintf(":%s", cfg.Port)
	if err := r.Run(address); err != nil {
		log.Fatalf("server run: %v", err)
	}

}
