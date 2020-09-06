package authenticate

import (
	"encoding/json"
	"fmt"
	"github.com/biancarosa/spotify-cli/random"
	"golang.org/x/oauth2"
	"os"

	"github.com/zmb3/spotify"
)

var (
	ch          = make(chan *spotify.Client)
	redirectURL = "http://localhost:8080/callback"
	state       = random.GenerateRandomString(10)
)

// GetClient é um método que autentica via oauth2 no spotify e retorna um spotify client
func GetClient() *spotify.Client {
	auth := spotify.NewAuthenticator(redirectURL, spotify.ScopeUserReadPrivate,
		spotify.ScopeUserLibraryRead, spotify.ScopeUserReadCurrentlyPlaying,
		spotify.ScopeUserModifyPlaybackState)

	f, err := os.Open("token.json")
	if err == nil {
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

	s := new(Server)
	s.Start(ch, state, &auth)
	url := auth.AuthURL(state)
	fmt.Printf("%s %s\n", "Acesse a URL em :: ", url)
	client := <-ch
	s.Stop()
	return client
}
