package model

// JiraUser represents a user from a Jira instance to further work with in the CLI app. A user is
// retrieved from a Jira instance through a service function which calls the Jira API.
type TronConfig struct {
	BaseURL  string
	Username string
	Password string
}

// Config represents the global configuration
// TODO: replace this with a service that reads a local tron.yml file to builds the config
var Config = TronConfig{
	BaseURL:  "http://localhost:8080",
	Username: "admin",
	Password: "admin",
}
