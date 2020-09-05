package cli

import (
	"fmt"
	"github.com/zmb3/spotify"
	"time"
)

//SpotifyHandler receives a client and handles commands on Spotify API
type SpotifyHandler struct {
	client *spotify.Client
}

//Now show the currently playing song on spotify
func (h *SpotifyHandler) Now() {
	cp, err := h.client.PlayerCurrentlyPlaying()
	if err == nil {
		if cp.Playing {
			fmt.Printf("Tocando agora: %s by %s", cp.Item.Name, cp.Item.Artists[0].Name)
		} else {
			fmt.Println("Você não está ouvindo nada agora!")
		}
	} else {
		fmt.Println(err.Error())
	}
}

//Play plays a song
func (h *SpotifyHandler) Play() {
	err := h.client.Play()
	if err == nil {
		fmt.Println("Curta sua musiquinha!")
	} else {
		fmt.Println(err.Error())
	}
}

//Pause pauses a song
func (h *SpotifyHandler) Pause() {
	err := h.client.Pause()
	if err == nil {
		fmt.Println("Música pausada")
	} else {
		fmt.Println(err.Error())
	}
}

//Next jumps to a new song
func (h *SpotifyHandler) Next() {
	err := h.client.Next()
	if err == nil {
		fmt.Println("Mudei de música!")
	} else {
		fmt.Println(err.Error())
	}
	time.Sleep(2 * time.Second)
	h.Now()
}
