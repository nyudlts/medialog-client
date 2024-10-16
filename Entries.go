package medialogclient

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/google/uuid"
	"github.com/nyudlts/go-medialog/models"
)

const XMEDIALOGTOKEN = "X-Medialog-Token"

func (mlc MedialogClient) GetEntryUUID(uuid uuid.UUID) (models.Entry, error) {
	entry := models.Entry{}
	url := fmt.Sprintf("%s/entries/%s", mlc.BaseURL, uuid)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return entry, err
	}
	req.Header.Add("X-Medialog-Token", mlc.Token)
	resp, err := mlc.Client.Do(req)
	if err != nil {
		return entry, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return entry, err
	}

	if err := json.Unmarshal(body, &entry); err != nil {
		return entry, err
	}

	return entry, nil
}

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

func (mlc MedialogClient) GetEntryUUIDsForResource(resourceID int) ([]string, error) {

	reqURL := fmt.Sprintf("%s/resources/%d/entries?all_ids=true", mlc.BaseURL, resourceID)
	req, err := http.NewRequest("GET", reqURL, nil)

	if err != nil {
		return []string{}, err
	}
	req.Header.Add("X-Medialog-Token", mlc.Token)

	resp, err := mlc.Client.Do(req)
	if err != nil {
		return []string{}, err
	}

	if resp.StatusCode != 200 {
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
