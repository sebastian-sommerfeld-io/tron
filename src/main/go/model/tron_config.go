package model

// TronConfig represents the global configuration for this app. The config holds all information
// to connect to a Jira instance.
type TronConfig struct {
	BaseURL  string
	Username string
	Password string
}

// Config represents the global configuration object.
// TODO: replace this with a service that reads a local tron.yml file to builds the config
var Config = TronConfig{
	BaseURL:  "http://localhost:8080",
	Username: "admin",
	Password: "admin",
}
