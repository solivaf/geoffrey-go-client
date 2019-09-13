package client_test

import (
	"geoffrey-go-client/pkg/client"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestClient_GetConfig(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		p := strings.Split(r.URL.Path, "/")
		if p[1] != "geoffrey" || p[2] != "dev" {
			t.Error("path should contains app and profile parameters")
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"bar": {"foo": "testPropertiesYml"}}`))
	}))
	defer server.Close()

	c := &http.Client{}
	gc := client.NewGeoffreyClient(server.URL, c)

	config, err := gc.GetConfig("geoffrey", "dev")
	if err != nil {
		t.Errorf("error retrieve configuration %v", err)
	}

	if _, ok := config["bar"]; !ok {
		t.Error("body response invalid")
	}
}
