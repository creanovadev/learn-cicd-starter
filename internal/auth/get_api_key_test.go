package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name        string
		headers     http.Header
		wantAPIKey  string
		wantErr     bool
	}{
		{
			name: "valid api key",
			headers: http.Header{
				"Authorization": []string{"ApiKey test-api-key"},
			},
			wantAPIKey: "test-api-key",
			wantErr:    false,
		},
		{
			name:       "missing authorization header",
			headers:    http.Header{},
			wantAPIKey: "",
			wantErr:    true,
		},
		{
			name: "malformed authorization header",
			headers: http.Header{
				"Authorization": []string{"Bearer test-api-key"},
			},
			wantAPIKey: "",
			wantErr:    true,
		},
		{
			name: "missing api key value",
			headers: http.Header{
				"Authorization": []string{"ApiKey"},
			},
			wantAPIKey: "",
			wantErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotAPIKey, err := GetAPIKey(tt.headers)

			if tt.wantErr && err == nil {
				t.Fatal("expected error, got nil")
			}

			if !tt.wantErr && err != nil {
				t.Fatalf("expected no error, got %v", err)
			}

			if gotAPIKey != tt.wantAPIKey {
				t.Errorf("expected api key %q, got %q", tt.wantAPIKey, gotAPIKey)
			}
		})
	}
}