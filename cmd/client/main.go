package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"quantum-exposer/app/controllers"
	"quantum-exposer/app/router"
	"quantum-exposer/internal/infrastructure/danbooru"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found or error loading it")
	}

	postRepo := danbooru.InitializeDanbooruService()
	postController := controllers.NewPostController(postRepo)
	router := router.SetupRouter(*postController)
	url, port := os.Getenv("APP_URL"), os.Getenv("APP_PORT")

	fmt.Printf("Server starting with Gin on %s:%s\n", url, port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
