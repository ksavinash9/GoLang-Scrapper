package api_test

import (
	"api"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var (
	server   *httptest.Server
	reader   io.Reader
	usersUrl string
)

func init() {
	server = httptest.NewServer(api.Handlers())

	amazonScrapeUrl = fmt.Sprintf("%s/movie/amazon/", server.URL)
}

func TestProductScrapping(t *testing.T) {
	reader = strings.NewReader("")

	request, err := http.NewRequest("GET", amazonScrapeUrl + "B00RH5G8K2", reader)

	res, err := http.DefaultClient.Do(request)

	if err != nil {
		t.Error(err)
	}

	if res.StatusCode == 200 {
		t.Errorf("Success expected: %d", res.StatusCode)
	}
}

func TestInvalidProductScrapping(t *testing.T) {
	reader = strings.NewReader("")

	request, err := http.NewRequest("GET", amazonScrapeUrl + "B00RH678K2", reader)

	res, err := http.DefaultClient.Do(request)

	if err != nil {
		t.Error(err)
	}

	if res.StatusCode == 404 {
		t.Errorf("Success expected: %d", res.StatusCode)
	}
}

func TestServerError(t *testing.T) {
	reader = strings.NewReader("")

	request, err := http.NewRequest("GET", amazonScrapeUrl + "B00RH/5G8K2", reader)

	res, err := http.DefaultClient.Do(request)

	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != 200 {
		t.Errorf("Success expected: %d", res.StatusCode)
	}
}
