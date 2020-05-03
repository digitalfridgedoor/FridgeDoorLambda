package main

import (
	"encoding/json"
	"testing"

	"github.com/digitalfridgedoor/fridgedoordatabase/dfdtesting"

	"github.com/digitalfridgedoor/fridgedoorapi/fridgedoorgatewaytesting"

	"github.com/stretchr/testify/assert"
)

func TestHandler(t *testing.T) {

	// Arrange
	dfdtesting.SetTestCollectionOverride()
	apirequest := fridgedoorgatewaytesting.CreateTestAuthorizedRequest("TestUser")

	// Act
	response, err := Handler(*apirequest)

	// Assert
	assert.Nil(t, err)
	var tags []string

	err = json.Unmarshal([]byte(response.Body), &tags)
	assert.Nil(t, err)
	assert.Nil(t, tags)

	fridgedoorgatewaytesting.DeleteUserForRequest(apirequest)
}
