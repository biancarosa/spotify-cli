package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/zmb3/spotify"
)

var (
	redirectURL = "http://localhost:8080/callback"
	state       = "login" // TODO: Gerar randomicamente
	ch          = make(chan *spotify.Client)
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Envie pelo menos um argumento")
		return
	}
	command := os.Args[1]
	auth := spotify.NewAuthenticator(redirectURL, spotify.ScopeUserReadPrivate,
		spotify.ScopeUserLibraryRead, spotify.ScopeUserReadCurrentlyPlaying,
		spotify.ScopeUserModifyPlaybackState)

	url := auth.AuthURL(state)
	fmt.Printf("%s %s\n", "Acesse a URL em :: ", url)

	http.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Recebi o seu request")
		token, _ := auth.Token(state, r)
		fmt.Println(token)
		client := auth.NewClient(token)
		ch <- &client
	})
	go http.ListenAndServe(":8080", nil)

	client := <-ch

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
