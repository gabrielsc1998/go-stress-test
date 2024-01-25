package stress_test

import "testing"

func TestStressTest(t *testing.T) {
	stressTest := New("http://google.com", 11, 2)
	result := stressTest.Run()
	if result.TotalRequests != 11 {
		t.Errorf("Expected 11 requests, got %v", result.TotalRequests)
	}
	if result.TotalTime == 0 {
		t.Error("Expected TotalTime to be greater than 0")
	}
	if len(result.StatusCodes) == 0 {
		t.Error("Expected StatusCodes to be greater than 0")
	}
	if result.StatusCodes[200] != 10 {
		t.Errorf("Expected 10 requests with status code 200, got %v", result.StatusCodes[200])
	}
	if len(result.Errors) != 0 {
		t.Errorf("Expected 0 errors, got %v", len(result.Errors))
	}
}
