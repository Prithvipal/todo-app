package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"strings"

	"github.com/prithvipal/todo-app/models"
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

func validateParansAndExtract(values url.Values) (map[string]any, error) {
	var err error
	title := values.Get("title")
	var status *models.StatusType
	if values.Get("status") != "" {
		status, err = models.NewStatusType(values.Get("status"))
		if err != nil {
			return nil, err
		}

	}

	sort := values.Get("sort")
	if sort != "" && sort != "title" && sort != "status" && sort != "created_at" && sort != "updated_at" {
		return nil, fmt.Errorf("invalid status key. valid sort keys are title, status, created_at and updated_at")
	}

	m := map[string]any{
		"title":  title,
		"status": status,
		"sort":   sort,
	}
	return m, nil
}
