package parse

import (
	"encoding/json"
	"io"
	"net/http"

	"parse/entities"
)

func requestFiles(url string) ([]entities.FileJSON, error) {
	var files []entities.FileJSON
	client := &http.Client{}
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json; charset=utf-8")
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(responseBody, &files)
	return files, nil
}
