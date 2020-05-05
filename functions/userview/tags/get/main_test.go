package main

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/digitalfridgedoor/fridgedoorapi/fridgedoorgatewaytesting"
	"github.com/digitalfridgedoor/fridgedoordatabase/dfdtesting"

	"github.com/stretchr/testify/assert"
)

func TestHandler(t *testing.T) {

	// Arrange
	dfdtesting.SetTestCollectionOverride()
	apirequest := fridgedoorgatewaytesting.CreateTestAuthorizedRequest("TestUser")

	ctx := context.TODO()

	// Act
	response, err := Handler(*apirequest)

	// Assert
	assert.Nil(t, err)
	var tags []string

	err = json.Unmarshal([]byte(response.Body), &tags)
	assert.Nil(t, err)
	assert.Nil(t, tags)

	fridgedoorgatewaytesting.DeleteUserForRequest(ctx, apirequest)
}
