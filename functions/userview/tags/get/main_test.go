package main

import (
	"encoding/json"
	"testing"

	"github.com/digitalfridgedoor/fridgedoorapi/fridgedoorgatewaytesting"

	"github.com/digitalfridgedoor/fridgedoorapi"

	"github.com/stretchr/testify/assert"
)

func TestHandler(t *testing.T) {

	// Arrange
	apirequest := fridgedoorgatewaytesting.CreateTestAuthorizedRequest("TestUser")

	// Act
	fridgedoorapi.ConnectOrSkip(t)

	response, err := Handler(*apirequest)

	// Assert
	assert.Nil(t, err)
	var tags []string

	err = json.Unmarshal([]byte(response.Body), &tags)
	assert.Nil(t, err)
	assert.Nil(t, tags)

	fridgedoorgatewaytesting.DeleteUserForRequest(apirequest)
}
