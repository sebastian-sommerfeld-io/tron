package main

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ShouldBeInitialized(t *testing.T) {
	expectedLogPrefix := "[tron] "
	assert.Equal(t, log.Prefix(), expectedLogPrefix)
}
