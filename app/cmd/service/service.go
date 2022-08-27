package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/jinzhu/configor"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	dblogger "gorm.io/gorm/logger"

	"github.com/laterius/service_architecture_hw2/app/internal/domain"
	"github.com/laterius/service_architecture_hw2/app/internal/service"
	"github.com/laterius/service_architecture_hw2/app/internal/transport/client/dbrepo"
	"github.com/laterius/service_architecture_hw2/app/internal/transport/server/http"
	_ "github.com/laterius/service_architecture_hw2/app/migrations"
)

func main() {
	var cfg domain.Config
	err := configor.New(&configor.Config{Silent: true}).Load(&cfg, "config/config.yaml", "./config.yaml")
	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: dbrepo.Dsn(cfg.Db),
	}), &gorm.Config{
		Logger: dblogger.Default.LogMode(dblogger.Info),
	})
	if err != nil {
		panic(err)
	}

	userRepo := dbrepo.NewUserRepo(db)
	userService := service.NewUserService(userRepo)
	getUserHandler := http.NewGetUser(userService)
	postUserHandler := http.NewPostUser(userService)
	putUserHandler := http.NewPutUser(userService)
	patchUserHandler := http.NewPatchUser(userService)
	deleteUserHandler := http.NewDeleteUser(userService)

	srv := fiber.New(fiber.Config{})
	srv.Use(logger.New())

	api := srv.Group("/api")

	api.Post("/user", postUserHandler.Handle())
	api.Post("/users", postUserHandler.Handle()) //дублирует предыдущий ?

	api.Get("/user/:id", getUserHandler.Handle())
	api.Get("/users/:id", getUserHandler.Handle()) //дублирует предыдущий ?

	api.Put("/user/:id", putUserHandler.Handle())
	api.Put("/users/:id", putUserHandler.Handle()) //дублирует предыдущий ?

	api.Patch("/user/:id", patchUserHandler.Handle())
	api.Patch("/users/:id", patchUserHandler.Handle()) //дублирует предыдущий ?

	api.Delete("/user/:id", deleteUserHandler.Handle())
	api.Delete("/users/:id", deleteUserHandler.Handle()) //дублирует предыдущий ?

	srv.Get("/probe/live", http.RespondOk)
	srv.Get("/probe/ready", http.RespondOk)

	srv.All("/*", http.DefaultResponse)

	err = srv.Listen(fmt.Sprintf(":%s", cfg.Http.Port))
	if err != nil {
		panic(err)
	}
}
