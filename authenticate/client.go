package authenticate

import (
	"encoding/json"
	"fmt"
	"golang.org/x/oauth2"
	"net/http"
	"os"

	"github.com/zmb3/spotify"
)

var (
	ch          = make(chan *spotify.Client)
	redirectURL = "http://localhost:8080/callback"
	state       = "login" // TODO: Gerar randomicamente
)

// GetClient é um método que autentica via oauth2 no spotify e retorna um spotify client
func GetClient() *spotify.Client {
	auth := spotify.NewAuthenticator(redirectURL, spotify.ScopeUserReadPrivate,
		spotify.ScopeUserLibraryRead, spotify.ScopeUserReadCurrentlyPlaying,
		spotify.ScopeUserModifyPlaybackState)

	f, err := os.Open("token.json")
	if err != nil {
		fmt.Println(err.Error())

		http.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Recebi o seu request")
			token, _ := auth.Token(state, r)

			f, err := os.Create("token.json")
			if err != nil {
				fmt.Println(err.Error())
			}
			// io.Writer -> interface
			// File é uma estrutura que implementa os métodos de um io.Writer
			enc := json.NewEncoder(f)
			err = enc.Encode(token)
			if err != nil {
				fmt.Println(err.Error())
			}
			f.Close()

			client := auth.NewClient(token)
			ch <- &client
		})

		go http.ListenAndServe(":8080", nil)

		url := auth.AuthURL(state)
		fmt.Printf("%s %s\n", "Acesse a URL em :: ", url)
	} else {
		enc := json.NewDecoder(f)
		var token *oauth2.Token
		err = enc.Decode(&token)
		if err != nil {
			fmt.Println(err.Error())
		}
		f.Close()
		client := auth.NewClient(token)
		return &client
	}

	client := <-ch
	return client
}
