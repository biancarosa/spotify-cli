package main

import (
	"fmt"
	"os"

	"github.com/zmb3/spotify"
)

var (
	redirectURL = "http://localhost:8080/callback"
	state       = "login" // TODO: Gerar randomicamente
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Envie pelo menos um argumento")
		return
	}
	command := os.Args[1]
	auth := spotify.NewAuthenticator(redirectURL, spotify.ScopeUserReadPrivate,
		spotify.ScopeUserLibraryRead)

	// fmt.Println(os.Getenv("SPOTIFY_ID"))

	url := auth.AuthURL(state)
	fmt.Printf("%s %s\n", "Acesse a URL em :: ", url)
	switch command {
	case "now":
		fmt.Println("O que tá tocando agora?")
	default:
		fmt.Println("Comando não implementado")
	}
}
