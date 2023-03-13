package commands

import (
	"fmt"

	"github.com/sebastian-sommerfeld-io/tron/model"
	"github.com/spf13/cobra"
)

// NewCmdConfig initializes the `tron config` command.
func NewCmdConfig() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "config",
		Short: "Manage tron config.",
		Long:  "Manage tron config. Config contains connection information etc.",
		Args:  cobra.ExactArgs(0),
	}

	return cmd
}

// NewCmdConfigView initializes the `tron config view` command.
func NewCmdConfigView() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "view",
		Short: "Show tron config.",
		Long:  "Show tron config from yaml file. Config contains connection information etc.",
		Args:  cobra.ExactArgs(0),

		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(model.Config)
		},
	}

	return cmd
}

// NewCmdConfigInit initializes the `tron config init` command.
func NewCmdConfigInit() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "init",
		Short: "Initialize a new tron config file.",
		Long:  "Initialize a new tron config file. Config contains connection information etc.",
		Args:  cobra.ExactArgs(0),

		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("config view")
		},
	}

	return cmd
}
