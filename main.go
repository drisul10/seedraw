package main

import (
	"log"

	"github.com/drisul10/seedraw/configs"
	"github.com/drisul10/seedraw/helpers"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	server := NewServer()
	server.Run()
}

type Server struct {
	App  *fiber.App
	Port string
}

func NewServer() *Server {
	LoadEnv()

	app := fiber.New(configs.NewFiberConfig())
	port := helpers.GetEnvOrUseDefaultValue("PORT", "9000")

	return &Server{
		App:  app,
		Port: port,
	}
}

func (s *Server) Run() {
	s.App.Listen(":" + s.Port)
}

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Errror loading .env")
	}
}
