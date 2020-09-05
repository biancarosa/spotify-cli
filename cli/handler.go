package cli

import (
	"fmt"
	"github.com/biancarosa/spotify-cli/authenticate"
)

//HandleCommandLineInput é um metodo que cebe um spotify client e um comando para executar com ele
func HandleCommandLineInput(command string) {
	client := authenticate.GetClient()

	sh := SpotifyHandler{
		client: client,
	}
	switch command {
	case "now":
		sh.Now()
	case "play":
		sh.Play()
	case "pause":
		sh.Pause()
	case "next":
		sh.Next()
	default:
		fmt.Println("Comando não implementado")
	}
}
