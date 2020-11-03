package main

import (
	"learning-golang/api.example.com/src/system/app"
	DB "learning-golang/api.example.com/src/system/db"
	"os"

	"github.com/joho/godotenv"
)

var port string

func init() {

	// Accessing Port via command line
	/* flag.StringVar(&port, "port", "8000", "assigning the port that server should listen on.")

	flag.Parse() */

	// Accessing Port using env file (config.ini)
	if err := godotenv.Load("config.ini"); err != nil {
		panic(err)
	}

	envPort := os.Getenv("PORT")
	if len(envPort) > 0 {
		port = envPort
	}
}

func main() {
	db, err := DB.Connect()
	if err != nil {
		panic(err)
	}

	s := app.NewServer()

	s.Init(port, db)
	s.Start()
}
