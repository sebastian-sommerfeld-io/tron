package commands

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func Test_ShouldCreateCmdConfig(t *testing.T) {
	expectedCmd := &cobra.Command{
		Use:  "config",
		Args: cobra.ExactArgs(0),
	}

	got := NewCmdConfig()
	assert.NotNil(t, got)
	assert.Equal(t, expectedCmd.Use, got.Use)
	assert.NotEmpty(t, got.Short)
	assert.NotEmpty(t, got.Long)
	assert.False(t, got.Runnable(), "Command should NOT be runnable")
}

func Test_ShouldCreateCmdConfigView(t *testing.T) {
	expectedCmd := &cobra.Command{
		Use:  "view",
		Args: cobra.ExactArgs(0),
	}

	got := NewCmdConfigView()
	assert.NotNil(t, got)
	assert.Equal(t, expectedCmd.Use, got.Use)
	assert.NotEmpty(t, got.Short)
	assert.NotEmpty(t, got.Long)
	assert.True(t, got.Runnable(), "Command should be runnable")
}

func Test_ShouldCreateCmdConfigInit(t *testing.T) {
	expectedCmd := &cobra.Command{
		Use:  "init",
		Args: cobra.ExactArgs(0),
	}

	got := NewCmdConfigInit()
	assert.NotNil(t, got)
	assert.Equal(t, expectedCmd.Use, got.Use)
	assert.NotEmpty(t, got.Short)
	assert.NotEmpty(t, got.Long)
	assert.True(t, got.Runnable(), "Command should be runnable")
}
