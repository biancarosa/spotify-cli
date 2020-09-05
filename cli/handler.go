package cli

import (
	"fmt"

	"github.com/biancarosa/spotify-cli/authenticate"
)

//HandleCommandLineInput é um metodo que cebe um spotify client e um comando para executar com ele
func HandleCommandLineInput(command string) {
	client := authenticate.GetClient()

	switch command {
	case "now":
		cp, err := client.PlayerCurrentlyPlaying()
		if err == nil {
			if cp.Playing {
				fmt.Printf("Tocando agora: %s by %s", cp.Item.Name, cp.Item.Artists[0].Name)
			} else {
				fmt.Println("Você não está ouvindo nada agora!")
			}
		} else {
			fmt.Println(err.Error())
		}
	case "play":
		err := client.Play()
		if err == nil {
			fmt.Println("Curta sua musiquinha!")
		} else {
			fmt.Println(err.Error())
		}
	case "pause":
		err := client.Pause()
		if err == nil {
			fmt.Println("Música pausada")
		} else {
			fmt.Println(err.Error())
		}
	case "next":
		err := client.Next()
		if err == nil {
			fmt.Println("Mudei de música!")
		} else {
			fmt.Println(err.Error())
		}
	default:
		fmt.Println("Comando não implementado")
	}
}
