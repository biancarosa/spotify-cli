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
		spotify.ScopeUserLibraryRead, spotify.ScopeUserReadCurrentlyPlaying)

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
		cp, _ := client.PlayerCurrentlyPlaying()
		fmt.Printf("Tocando agora: %s by %s", cp.Item.Name, cp.Item.Artists[0].Name)
	default:
		fmt.Println("Comando não implementado")
	}
}
