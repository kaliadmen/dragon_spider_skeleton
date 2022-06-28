package handlers

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHome(t *testing.T) {
	routes := getRoutes()
	testServer := httptest.NewTLSServer(routes)

	defer testServer.Close()

	res, err := testServer.Client().Get(testServer.URL + "/")
	if err != nil {
		t.Log(err)
	}

	if res.StatusCode != 200 {
		t.Errorf("home page expected status %d, got %d", http.StatusOK, res.StatusCode)
	}

	bodyText, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	if !strings.Contains(string(bodyText), "awesome") {
		ds.TakeScreenshot(testServer.URL+"/", "HomeTest", 1500, 1000)
		t.Error("not found")
	}

}
