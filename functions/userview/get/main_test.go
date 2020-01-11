package main

import (
	"encoding/json"
	"testing"

	"github.com/digitalfridgedoor/fridgedoorapi/linkeduserapi"

	"github.com/digitalfridgedoor/fridgedoorapi"
	"github.com/digitalfridgedoor/fridgedoorapi/dfdtesting"

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
	var linkedusers []*linkeduserapi.LinkedUser

	err = json.Unmarshal([]byte(response.Body), &linkedusers)
	assert.Nil(t, err)
	assert.NotNil(t, linkedusers)
	assert.Greater(t, len(linkedusers), 0)

	dfdtesting.DeleteUserForRequest(apirequest)
}
