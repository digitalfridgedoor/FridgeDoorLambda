package main

import (
	"encoding/json"
	"testing"

	"github.com/digitalfridgedoor/fridgedoorapi/dfdtesting"

	"github.com/digitalfridgedoor/fridgedoorapi"

	"github.com/stretchr/testify/assert"
)

func TestHandler(t *testing.T) {

	// Arrange
	apirequest := dfdtesting.CreateTestAuthorizedRequest("TestUser")

	// Act
	fridgedoorapi.ConnectOrSkip(t)

	response, err := Handler(*apirequest)

	// Assert
	assert.Nil(t, err)
	var tags []string

	err = json.Unmarshal([]byte(response.Body), &tags)
	assert.Nil(t, err)
	assert.Nil(t, tags)

	dfdtesting.DeleteUserForRequest(apirequest)
}
