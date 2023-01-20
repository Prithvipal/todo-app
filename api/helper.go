package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"
)

func writeJSON(w http.ResponseWriter, records any) {
	w.Header().Set("Content-Type", "application/json")
	data, err := json.Marshal(records)
	if err != nil {
		log.Println("Error while getting TODO List", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Write(data)
}

func validateUrlAndExtractParam(endpoint string) (string, error) {
	r := `^(?:\/api/v1/todo\b)(?:\/[\w-]+)$`
	match, err := regexp.MatchString(r, endpoint)
	if err != nil {
		return "", err
	}
	if !match {
		return "", fmt.Errorf("endpoint does not match")
	}
	id := strings.TrimPrefix(endpoint, "/api/v1/todo/")
	return id, nil
}
