package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/nyudlts/go-medialog/models"
	yaml "gopkg.in/yaml.v2"
)

var (
	config      string
	environment string
)

func init() {
	flag.StringVar(&config, "config", "", "")
	flag.StringVar(&environment, "environment", "", "")
}

type Creds struct {
	URL      string `yaml:"url"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type MedialogClient struct {
	Token   string
	BaseURL string
	Client  *http.Client
}

const API_ROOT = "/api/v0"

func NewClient(creds Creds, timeout int) (MedialogClient, error) {
	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    time.Duration(timeout) * time.Second,
		DisableCompression: true,
	}

	mlClient := MedialogClient{}
	mlClient.Client = &http.Client{Transport: tr}
	mlClient.BaseURL = creds.URL + API_ROOT
	mlClient.getToken(creds)

	return mlClient, nil
}

func getCreds(config string, environment string) (Creds, error) {
	credsMap := map[string]Creds{}
	configBytes, err := os.ReadFile(config)
	if err != nil {
		return Creds{}, err
	}

	if err = yaml.Unmarshal(configBytes, &credsMap); err != nil {
		return Creds{}, err
	}

	for k, v := range credsMap {
		if environment == k {
			return v, nil
		}
	}

	return Creds{}, fmt.Errorf("Credentials file did not contain %s\n", environment)
}

func (mlClient *MedialogClient) getToken(creds Creds) {
	url := fmt.Sprintf("%s/users/%s/login?password=%s", mlClient.BaseURL, creds.Username, creds.Password)
	response, err := mlClient.Client.Post(url, "", nil)
	if err != nil {
		panic(err.Error())
	}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	token := models.Token{}
	if err := json.Unmarshal(body, &token); err != nil {
		panic(err)
	}

	mlClient.Token = token.Token
}

func (mlc MedialogClient) GetRoot() (models.MedialogInfo, error) {
	mlInfo := models.MedialogInfo{}
	req, err := http.NewRequest("GET", mlc.BaseURL, nil)
	if err != nil {
		return mlInfo, err
	}

	resp, err := mlc.Client.Do(req)
	if err != nil {
		return mlInfo, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return mlInfo, err
	}

	if err := json.Unmarshal(body, &mlInfo); err != nil {
		return mlInfo, err
	}

	return mlInfo, nil

}
