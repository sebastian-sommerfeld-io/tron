package commands

import (
	"github.com/spf13/cobra"
)

// NewCmdUser initializes the `tron user list-projects` command.
func NewCmdJira() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "jira",
		Short: "Interact with a Jira Instance.",
		Long:  "Interact with a Jira Instance.",
		Args:  cobra.ExactArgs(0),
	}

	return cmd
}
