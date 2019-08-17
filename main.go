package main

import (
	"fmt"
	"github.com/fhs/gompd/mpd"
	"mpd-bob/handler"
	"net/http"
)

func main() {

	client, err := mpd.Dial("tcp", "192.168.10.16:6600")
	if err != nil {
		panic(err)
	}

	mpdHandler := handler.NewMpdHandler(client)

	http.HandleFunc("/play", mpdHandler.HandlePlay)
	http.HandleFunc("/pause", mpdHandler.HandlePause)
	http.HandleFunc("/next", mpdHandler.HandleNext)
	http.HandleFunc("/previous", mpdHandler.HandlePrevious)
	http.HandleFunc("/search", mpdHandler.HandleSearch)

	fmt.Println(http.ListenAndServe(":5000", nil))
}
