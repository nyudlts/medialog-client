package medialogclient

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/nyudlts/go-medialog/models"
)

func (mlc MedialogClient) GetResources() ([]models.Resource, error) {
	resources := []models.Resource{}
	endpoint := fmt.Sprintf("%s/resources", mlc.BaseURL)
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return resources, err
	}

	req.Header.Add(XMEDIALOGTOKEN, mlc.Token)
	resp, err := mlc.Client.Do(req)
	if err != nil {
		return resources, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return resources, err
	}

	if err := json.Unmarshal(body, &resources); err != nil {
		return resources, err
	}

	return resources, nil
}
