package main

import (
	todoList "github.com/Ilhomiddin123"
	"github.com/Ilhomiddin123/pkg/handler"
	"log"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(todoList.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
