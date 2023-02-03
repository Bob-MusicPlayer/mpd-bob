package handler

import (
	"encoding/json"
	"fmt"
	"github.com/Bob-MusicPlayer/mpd-bob/model"
	"github.com/fhs/gompd/mpd"
	"net/http"
	"strconv"
)

type MpdHandler struct {
	mpdClient *mpd.Client
}

func NewMpdHandler(client *mpd.Client) MpdHandler {
	return MpdHandler{
		mpdClient: client,
	}
}

func (mh MpdHandler) Handle(w http.ResponseWriter, req *http.Request) {

}

func (mh MpdHandler) HandlePlay(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost && req.Method != http.MethodOptions {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	if req.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(200)
		return
	}

	err := mh.mpdClient.Pause(false)
	if err != nil {
		w.WriteHeader(901)
		w.Write([]byte(err.Error()))
	} else {
		w.WriteHeader(200)
	}
}

func (mh MpdHandler) HandlePause(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost && req.Method != http.MethodOptions {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	if req.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(200)
		return
	}

	err := mh.mpdClient.Pause(true)
	if err != nil {
		w.WriteHeader(901)
		w.Write([]byte(err.Error()))
	} else {
		w.WriteHeader(200)
	}
}

func (mh MpdHandler) HandleNext(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost && req.Method != http.MethodOptions {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	if req.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(200)
		return
	}

	err := mh.mpdClient.Next()
	if err != nil {
		w.WriteHeader(901)
		w.Write([]byte(err.Error()))
	} else {
		w.WriteHeader(200)
	}
}

func (mh MpdHandler) HandlePrevious(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost && req.Method != http.MethodOptions {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	if req.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(200)
		return
	}

	err := mh.mpdClient.Previous()
	if err != nil {
		w.WriteHeader(901)
		w.Write([]byte(err.Error()))
	} else {
		w.WriteHeader(200)
	}
}

func (mh MpdHandler) HandleSearch(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost && req.Method != http.MethodOptions {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if req.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(200)
		return
	}

	var search model.Search

	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&search)
	if err != nil {
		w.WriteHeader(901)
		w.Write([]byte(err.Error()))
		return
	}

	songs := make([]model.Song, 0)

	attr, err := mh.mpdClient.Find(search.SearchQuery)
	if err != nil {
		w.WriteHeader(901)
		w.Write([]byte(err.Error()))
	} else {
		for _, v := range attr {
			var song model.Song
			song.Title = v["Title"]
			song.Artist = v["Artist"]
			song.Duration, err = strconv.ParseFloat(v["duration"], 32)
			if err != nil {
				fmt.Println(err)
			}
			songs = append(songs, song)
		}
		json, err := json.Marshal(songs)
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}
		w.Write(json)
		w.WriteHeader(200)
	}
}
