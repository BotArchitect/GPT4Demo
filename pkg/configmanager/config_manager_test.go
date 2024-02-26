// config_manager_test.go
package configmanager

import (
	"os"
	"testing"
)

// createTestConfigFile creates a temporary configuration file for testing.
func createTestConfigFile() (string, error) {
	content := []byte(`{
		"githubAPIKey": "test_api_key",
		"dbConnectionString": "test_db_connection_string",
		"logLevel": "test_log_level"
	}`)
	tmpfile, err := os.CreateTemp("", "config")
	if err != nil {
		return "", err
	}
	if _, err := tmpfile.Write(content); err != nil {
		return "", err
	}
	if err := tmpfile.Close(); err != nil {
		return "", err
	}
	return tmpfile.Name(), nil
}

func TestNewConfigManager(t *testing.T) {
	configFilePath, err := createTestConfigFile()
	if err != nil {
		t.Fatalf("Failed to create test config file: %v", err)
	}
	defer os.Remove(configFilePath) // clean up

	cm, err := NewConfigManager(configFilePath)
	if err != nil {
		t.Fatalf("NewConfigManager() error = %v, wantErr = false", err)
	}

	if cm.GetGithubAPIKey() != "test_api_key" {
		t.Errorf("GetGithubAPIKey() = %v, want %v", cm.GetGithubAPIKey(), "test_api_key")
	}
	if cm.GetDBConnectionString() != "test_db_connection_string" {
		t.Errorf("GetDBConnectionString() = %v, want %v", cm.GetDBConnectionString(), "test_db_connection_string")
	}
	if cm.GetLogLevel() != "test_log_level" {
		t.Errorf("GetLogLevel() = %v, want %v", cm.GetLogLevel(), "test_log_level")
	}
}
