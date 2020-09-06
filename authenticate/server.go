package authenticate

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/zmb3/spotify"
)

type Server struct {
	srv *http.Server
}

func (s *Server) Start(ch chan *spotify.Client, state string, auth *spotify.Authenticator) {
	s.srv = &http.Server{
		Addr: ":8080",
	}

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

	go s.srv.ListenAndServe()
}

func (s *Server) Stop() {
	s.srv.Close()
}
