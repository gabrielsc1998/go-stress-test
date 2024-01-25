package http_client

import "testing"

func TestShouldReturnStatus200WhenDoValidRequest(t *testing.T) {
	resp, err := Get("http://google.com")
	if err != nil {
		t.Errorf("Error on do request: %v", err)
	}
	if resp.StatusCode != 200 {
		t.Errorf("Expected status code 200, got %v", resp.StatusCode)
	}
}

func TestShouldReturn404WhenDoRequestForInexistentPage(t *testing.T) {
	resp, err := Get("http://google.com/invalid")
	if err != nil {
		t.Error("Expected errror to be nil")
	}
	if resp.StatusCode != 404 {
		t.Errorf("Expected 404, got %v", resp.StatusCode)
	}
}
