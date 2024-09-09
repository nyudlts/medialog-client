package main

import (
	"flag"
	"testing"
)

var mlc MedialogClient

func TestClient(t *testing.T) {
	flag.Parse()

	t.Run("test create client", func(t *testing.T) {
		creds, err := getCreds(config, environment)
		if err != nil {
			t.Error(err)
		}

		mlc, err = NewClient(creds, 20)
		if err != nil {
			t.Error(err)
		}
	})

	t.Run("test get medialog info", func(t *testing.T) {

		mlInfo, err := mlc.GetRoot()
		if err != nil {
			t.Error(err)
		}

		t.Log(mlInfo)
	})
}
