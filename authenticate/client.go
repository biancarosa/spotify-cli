package authenticate

import (
	"fmt"
	"net/http"

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

	http.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Recebi o seu request")
		token, _ := auth.Token(state, r)
		fmt.Println(token)
		client := auth.NewClient(token)
		ch <- &client
	})

	go http.ListenAndServe(":8080", nil)

	url := auth.AuthURL(state)
	fmt.Printf("%s %s\n", "Acesse a URL em :: ", url)

	client := <-ch

	return client
}
