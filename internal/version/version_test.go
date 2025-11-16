package version

import (
	"testing"
)

func TestGetVersion(t *testing.T) {
	v := GetVersion()
	if v == "" {
		t.Error("GetVersion() should not return empty string")
	}
}

func TestGetBuildInfo(t *testing.T) {
	info := GetBuildInfo()

	requiredKeys := []string{"version", "commit", "buildDate"}
	for _, key := range requiredKeys {
		if _, ok := info[key]; !ok {
			t.Errorf("GetBuildInfo() missing required key: %s", key)
		}
	}
}

func TestVersionNotEmpty(t *testing.T) {
	if Version == "" {
		t.Error("Version should not be empty")
	}
}

func TestBuildInfoStructure(t *testing.T) {
	info := GetBuildInfo()

	tests := []struct {
		name string
		key  string
		want string
	}{
		{"version matches", "version", Version},
		{"commit matches", "commit", GitCommit},
		{"buildDate matches", "buildDate", BuildDate},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := info[tt.key]; got != tt.want {
				t.Errorf("GetBuildInfo()[%s] = %v, want %v", tt.key, got, tt.want)
			}
		})
	}
}
