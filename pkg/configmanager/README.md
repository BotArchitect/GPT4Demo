# ConfigManager Interface

The `ConfigManager` interface abstracts the methods for accessing configuration settings. Implementations of this interface must provide functionality to load the configuration and retrieve specific settings.

## Interface

```go
type ConfigManagerInterface interface {
    GetGithubAPIKey() string
    GetDBConnectionString() string
    GetLogLevel() string
}
