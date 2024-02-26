// config_manager.go
package configmanager

import (
	"encoding/json"
	"io/ioutil"
	"sync"
)

// ConfigManagerInterface outlines the methods for accessing configuration settings.
type ConfigManagerInterface interface {
	GetGithubAPIKey() string
	GetDBConnectionString() string
	GetLogLevel() string
}

// Config represents the structure of the configuration file.
type Config struct {
	GithubAPIKey       string `json:"githubAPIKey"`
	DBConnectionString string `json:"dbConnectionString"`
	LogLevel           string `json:"logLevel"`
}

// ConfigManager implements ConfigManagerInterface and holds the application configuration.
type ConfigManager struct {
	config *Config
	mu     sync.RWMutex
}

// NewConfigManager creates a new ConfigManager, loading the configuration from the given file path.
func NewConfigManager(filePath string) (ConfigManagerInterface, error) {
	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var config Config
	err = json.Unmarshal(bytes, &config)
	if err != nil {
		return nil, err
	}

	return &ConfigManager{config: &config}, nil
}

// GetGithubAPIKey returns the GitHub API Key from the configuration.
func (cm *ConfigManager) GetGithubAPIKey() string {
	cm.mu.RLock()
	defer cm.mu.RUnlock()
	return cm.config.GithubAPIKey
}

// GetDBConnectionString returns the database connection string from the configuration.
func (cm *ConfigManager) GetDBConnectionString() string {
	cm.mu.RLock()
	defer cm.mu.RUnlock()
	return cm.config.DBConnectionString
}

// GetLogLevel returns the log level from the configuration.
func (cm *ConfigManager) GetLogLevel() string {
	cm.mu.RLock()
	defer cm.mu.RUnlock()
	return cm.config.LogLevel
}
