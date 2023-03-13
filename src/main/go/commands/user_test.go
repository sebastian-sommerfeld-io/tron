package commands

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func Test_ShouldCreateCmdUser(t *testing.T) {
	expectedCmd := &cobra.Command{
		Use:  "user",
		Args: cobra.ExactArgs(0),
	}

	got := NewCmdUser()
	assert.NotNil(t, got)
	assert.Equal(t, expectedCmd.Use, got.Use)
	assert.NotEmpty(t, got.Short)
	assert.NotEmpty(t, got.Long)
	assert.False(t, got.Runnable(), "Command should NOT be runnable")
}

func Test_ShouldCreateCmdUserRead(t *testing.T) {
	expectedCmd := &cobra.Command{
		Use:  "read",
		Args: cobra.ExactArgs(0),
	}

	got := NewCmdUserRead()
	assert.NotNil(t, got)
	assert.Equal(t, expectedCmd.Use, got.Use)
	assert.NotEmpty(t, got.Short)
	assert.NotEmpty(t, got.Long)
	assert.True(t, got.Runnable(), "Command should be runnable")
}

func Test_ShouldCreateCmdUserExists(t *testing.T) {
	expectedCmd := &cobra.Command{
		Use:  "exists",
		Args: cobra.ExactArgs(0),
	}

	got := NewCmdUserExists()
	assert.NotNil(t, got)
	assert.Equal(t, expectedCmd.Use, got.Use)
	assert.NotEmpty(t, got.Short)
	assert.NotEmpty(t, got.Long)
	assert.True(t, got.Runnable(), "Command should be runnable")
}

func Test_ShouldCreateCmdUserListProjects(t *testing.T) {
	expectedCmd := &cobra.Command{
		Use:  "list-projects",
		Args: cobra.ExactArgs(0),
	}

	got := NewCmdUserListProjects()
	assert.NotNil(t, got)
	assert.Equal(t, expectedCmd.Use, got.Use)
	assert.NotEmpty(t, got.Short)
	assert.NotEmpty(t, got.Long)
	assert.True(t, got.Runnable(), "Command should be runnable")
}
