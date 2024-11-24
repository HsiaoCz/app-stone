package handlers

import "net/http"

type RecordHandlers struct{}

func (c *RecordHandlers) HandleCreateRecord(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (c *RecordHandlers) HandleGetRecordsByUserID(w http.ResponseWriter, r *http.Request) error {
	return nil
}
