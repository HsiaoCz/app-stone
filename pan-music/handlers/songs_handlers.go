package handlers

import "net/http"

type SongsHandlers struct{}

func (s *SongsHandlers) HandleCreateSong(w http.ResponseWriter, r *http.Request) error {
	return nil
}
