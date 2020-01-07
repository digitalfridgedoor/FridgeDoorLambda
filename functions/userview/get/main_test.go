package main

import (
	"encoding/json"
	"testing"

	"github.com/digitalfridgedoor/fridgedoorapi"
	"github.com/digitalfridgedoor/fridgedoorapi/dfdtesting"
	"github.com/digitalfridgedoor/fridgedoorapi/userviewapi"

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
	var recipes []*userviewapi.View

	err = json.Unmarshal([]byte(response.Body), &recipes)
	assert.Nil(t, err)
	assert.NotNil(t, recipes)
	assert.Greater(t, len(recipes), 0)

	dfdtesting.DeleteUserForRequest(apirequest)
}
