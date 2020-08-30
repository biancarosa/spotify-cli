package main

import (
	"fmt"
	"os"

	"github.com/biancarosa/spotify-cli/authenticate"
	"github.com/biancarosa/spotify-cli/cli"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Envie pelo menos um argumento")
		return
	}
	command := os.Args[1]
	client := authenticate.GetClient()
	cli.HandleCommandLineInput(client, command)
}
