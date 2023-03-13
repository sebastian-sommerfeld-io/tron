package commands

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

const (
	// FlagUser contains the username
	FlagUsername string = "username"
)

// AddFlags adds a set of flags to the given command.
func addFlags(cmd *cobra.Command) {
	cmd.Flags().String(FlagUsername, "", "The username to search for")
	err := cmd.MarkFlagRequired(FlagUsername)
	if err != nil {
		log.Fatal(err)
	}
}

func getUsernameValue(cmd *cobra.Command) string {
	value, err := cmd.Flags().GetString(FlagUsername)
	if err != nil {
		log.Fatal(err)
	}
	return value
}

// NewCmdUser initializes the `tron user` command.
func NewCmdUser() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "user",
		Short: "Interact with a Jira instance and get Jira user information.",
		Long:  "Interact with a Jira instance and count users or access information about a Jira user.",
		Args:  cobra.ExactArgs(0),
	}

	return cmd
}

// NewCmdUser initializes the `tron user list-projects` command.
func NewCmdUserExists() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "exists",
		Short: "Check if a Jira user exists.",
		Long:  "Check if a Jira user exists.",
		Args:  cobra.ExactArgs(0),

		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(getUsernameValue(cmd))
		},
	}

	addFlags(cmd)

	return cmd
}

// NewCmdUser initializes the `tron user list-projects` command.
func NewCmdUserListProjects() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-projects",
		Short: "List all projects a user is assigned to.",
		Long:  "List all projects a user is assigned to (either directly or through group memberships).",
		Args:  cobra.ExactArgs(0),

		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(getUsernameValue(cmd) + " call real implementation here")
		},
	}

	addFlags(cmd)

	return cmd
}
