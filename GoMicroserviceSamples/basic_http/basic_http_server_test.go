package basic_http

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// func TestServerStart(t *testing.T) {
// 	go server_initialize() //Run the server in the same routine.
// 	time.Sleep(1 * time.Second)
// 	response, err := http.Get("http://localhost:3000")
// 	if err != nil {
// 		t.Fatalf("Failed to make a GET request: %v", err)
// 	}
// 	defer response.Body.Close()
// 	if response.StatusCode != http.StatusOK {
// 		t.Errorf("Expected status code %v, but got %v", http.StatusOK, response.StatusCode)
// 	}
// }

func TestSampleEndpoint(t *testing.T) {
	request, err := http.NewRequest("GET", "http://localhost:3000/sampleendpoint", nil)
	if err != nil {
		t.Fatalf("Failed to create a request: %v", err)
	}

	response := httptest.NewRecorder()
	sampleEndpoint(response, request)
	if response.Code != http.StatusOK {
		t.Errorf("Expected status code %v, but got %v", http.StatusOK, response.Code)
	}

	expectedBody := "Hello Message from the app \n"
	if response.Body.String() != expectedBody {
		t.Errorf("Expected body '%v', but got '%v'", expectedBody, response.Body.String())
	}
}

func TestSendMessage(t *testing.T) {
	request, err := http.NewRequest("GET", "http://localhost:3000/sendmessage", nil)
	if err != nil {
		t.Fatalf("Failed to create a request: %v", err)
	}

	response := httptest.NewRecorder()
	sampleEndpoint(response, request)
	if response.Code != http.StatusOK {
		t.Errorf("Expected status code %v, but got %v", http.StatusOK, response.Code)

	}

	expectedBody := "Hello Message from the app \n"
	if response.Body.String() != expectedBody {
		t.Errorf("Expected body %v, but got %v", expectedBody, response.Body.String())
	}
}
