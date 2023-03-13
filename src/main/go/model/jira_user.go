package model

// JiraUser represents a user from a Jira instance to further work with in the CLI app. A user is
// retrieved from a Jira instance through a service function which calls the Jira API.
type JiraUser struct {
	Id        int
	Firstname string
	Lastname  string
	Username  string
	Email     string
}
