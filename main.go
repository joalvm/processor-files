package main

import (
	"github.com/joalvm/processor-medias/cmd"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	cmd.Execute()
}
