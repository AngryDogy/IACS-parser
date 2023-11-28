package httpreq

import (
	"encoding/json"
	"io"
	"net/http"

	"parse/iternal/entities"
)

func GetFilesRequest(url string) ([]entities.File, error) {
	var files []entities.File
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
