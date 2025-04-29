package auth

import (
	"net/http"
	"testing"
)

func TestValidGetAPIKey(t *testing.T) {
	url := "localhost.com"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		t.Fatalf("Error encountered setting up request")
	}

	req.Header.Add("Authorization", "ApiKey 123APIKEY")
	expected := "123APIKEY"
	actual, err := GetAPIKey(req.Header)
	if err != nil {
		t.Fatalf("There was an error getting the api key: %v", err)
	}

	if actual != expected {
		t.Fatalf("Expected API key, `%v`, does not equal actual. got=%v", expected, actual)
	}
}

func TestMalformedGetAPIKey(t *testing.T) {
	url := "localhost.com"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		t.Fatalf("Error encountered setting up request")
	}

	req.Header.Add("Authorization", "ApiKey_123APIKEY")
	expected := "123APIKEY"
	actual, err := GetAPIKey(req.Header)
	if err.Error() == "malformed authorization header" {
		return
	}

	if err != nil {
		t.Fatalf("There was an error getting the api key: %v", err)
	}

	if actual != expected {
		t.Fatalf("Expected API key, `%v`, does not equal actual. got=%v", expected, actual)
	}
}
