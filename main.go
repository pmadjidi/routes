package main

import (
	"log"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("no .env found...")
	}
}

func main() {
	sys := createSys()
	sys.initSubSystems()
	sys.createTerminationHandler()
	sys.startHttp()
}
