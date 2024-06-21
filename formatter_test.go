package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFormatTimestamp(t *testing.T) {
	assert := assert.New(t)

	baseTimestamp := int64(1386788400) // 2013-12-11 19:00:00 +0000 UTC

	assert.Equal("2013-12-11 19:00:00", FormatTimestamp(baseTimestamp, ""), "default layout")

	assert.Equal("2013-12-11", FormatTimestamp(baseTimestamp, "2006-01-02"), "custom layout")

	assert.Equal("19:00", FormatTimestamp(baseTimestamp, "15:04"), "custom layout")
}
