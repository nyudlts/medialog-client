package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (mlc MedialogClient) GetEntryUUIDs() ([]string, error) {
	url := fmt.Sprintf("%s/entries?all_ids=true", mlc.BaseURL)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return []string{}, err
	}
	req.Header.Add("X-Medialog-Token", mlc.Token)
	resp, err := mlc.Client.Do(req)
	if err != nil {
		return []string{}, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return []string{}, err
	}

	uuids := []string{}
	if err := json.Unmarshal(body, &uuids); err != nil {
		return uuids, err
	}

	return uuids, nil

}
