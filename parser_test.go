package main

import "testing"

func TestParseConnectionDetails(t *testing.T) {
	tests := []struct {
		input    string
		wantHost string
		wantPort int
		wantErr  bool
	}{
		{"127.0.0.1:28016", "127.0.0.1", 28016, false},
		{"localhost:28016", "localhost", 28016, false},
		{"invalid:port", "", 0, true},
		{"missingport", "", 0, true},
		{"256.256.256.256:28016", "", 0, true},
		{"127.0.0.1:70000", "", 0, true},
		{"127.0.0.1:-1", "", 0, true},
		{"127.0.0.1:0", "", 0, true},
		{"127.0.0.1:65536", "", 0, true},
		{"example.com:80", "", 0, true},
	}

	for _, tt := range tests {
		gotHost, gotPort, gotErr := ParseConnectionDetails(tt.input)
		if (gotErr != nil) != tt.wantErr {
			t.Errorf("ParseConnectionDetails(%q) error = %v, wantErr %v", tt.input, gotErr, tt.wantErr)
			continue
		}
		if gotHost != tt.wantHost {
			t.Errorf("ParseConnectionDetails(%q) gotHost = %v, want %v", tt.input, gotHost, tt.wantHost)
		}
		if gotPort != tt.wantPort {
			t.Errorf("ParseConnectionDetails(%q) gotPort = %v, want %v", tt.input, gotPort, tt.wantPort)
		}
	}
}
