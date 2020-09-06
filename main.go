package main

import (
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"

	"github.com/biancarosa/spotify-cli/cli"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(".env file not found")
	}

	if len(os.Args) < 2 {
		fmt.Println("Envie pelo menos um argumento")
		return
	}
	command := os.Args[1]
	cli.HandleCommandLineInput(command)
}
