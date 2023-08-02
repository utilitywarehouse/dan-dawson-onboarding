package handler

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"testing"
)

func TestServeTime(t *testing.T) {
	go InitHandler()

	url := "http://localhost:8080/time"
	tt := []struct {
		name           string
		method         string
		expectedStatus int
		expectErr      bool
	}{
		{
			name:           "1. GET method used and should expect no error",
			method:         http.MethodGet,
			expectedStatus: http.StatusOK,
			expectErr:      false,
		}, {
			name:           "2. POST method used and should expect an error",
			method:         http.MethodPost,
			expectedStatus: http.StatusMethodNotAllowed,
			expectErr:      true,
		}, {
			name:           "3. PUT method used and should expect an error",
			method:         http.MethodPut,
			expectedStatus: http.StatusMethodNotAllowed,
			expectErr:      true,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {

			req, _ := http.NewRequest(tc.method, url, bytes.NewBuffer([]byte{}))
			client := &http.Client{}

			resp, _ := client.Do(req)

			if resp.StatusCode != tc.expectedStatus {
				t.Fatalf("incorrect response status code. expected: %v || got: %v\n", tc.expectedStatus, resp.StatusCode)
			}

			body, _ := io.ReadAll(resp.Body)

			if tc.expectErr {
				expectedErr := fmt.Sprintf("invalid request method used: %v\n", tc.method)

				if string(body) != expectedErr {
					t.Fatalf("incorrect error message. expected: %v || got: %v\n", expectedErr, string(body))
				}
			}
		})
	}
}
