package cli

import (
	"fmt"
	"github.com/zmb3/spotify"
	"time"

	"github.com/biancarosa/spotify-cli/authenticate"
)

//CommandHandler handles a command from the Command Line Interface
type CommandHandler interface {
	HandleCommand()
}

type PlayingNow struct {
	client *spotify.Client
}

func (h *PlayingNow) HandleCommand() {
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

type Play struct {
	client *spotify.Client
}

func (h *Play) HandleCommand() {
	err := h.client.Play()
	if err == nil {
		fmt.Println("Curta sua musiquinha!")
	} else {
		fmt.Println(err.Error())
	}
}

type Pause struct {
	client *spotify.Client
}

func (h *Pause) HandleCommand() {
	err := h.client.Pause()
	if err == nil {
		fmt.Println("Música pausada")
	} else {
		fmt.Println(err.Error())
	}
}

type Next struct {
	client *spotify.Client
}

func (h *Next) HandleCommand() {
	err := h.client.Next()
	if err == nil {
		fmt.Println("Mudei de música!")
	} else {
		fmt.Println(err.Error())
	}
	time.Sleep(2 * time.Second)
	pn := PlayingNow{
		client: h.client,
	}
	pn.HandleCommand()
}

//HandleCommandLineInput é um metodo que cebe um spotify client e um comando para executar com ele
func HandleCommandLineInput(command string) {
	client := authenticate.GetClient()

	switch command {
	case "now":
		h := PlayingNow{
			client: client,
		}
		h.HandleCommand()
	case "play":
		h := Play{
			client: client,
		}
		h.HandleCommand()
	case "pause":
		h := Pause{
			client: client,
		}
		h.HandleCommand()
	case "next":
		h := Next{
			client: client,
		}
		h.HandleCommand()
	default:
		fmt.Println("Comando não implementado")
	}
}
