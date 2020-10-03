package main

import (
	"context"
	"testing"

	"github.com/digitalfridgedoor/fridgedoorapi/clippingapi"

	"github.com/digitalfridgedoor/fridgedoorapi/dfdtesting"

	"github.com/stretchr/testify/assert"
)

func TestHandler(t *testing.T) {

	// Arrange
	dfdtesting.SetTestCollectionOverride()

	ctx := context.Background()
	name := "clipping name"
	testUserName := "TestUser"
	testUser := dfdtesting.CreateTestAuthenticatedUser(testUserName)
	clippingID, err := clippingapi.Create(ctx, testUser, name)
	assert.Nil(t, err)

	r, err := clippingapi.FindOne(ctx, testUser, clippingID)
	assert.Nil(t, err)
	assert.NotNil(t, r)

	pathParameters := make(map[string]string)
	pathParameters["id"] = clippingID.Hex()
	deleterequest := dfdtesting.CreateTestAuthorizedRequest(testUserName)
	deleterequest.PathParameters = pathParameters

	// Act
	response, err := Handler(*deleterequest)

	// Assert
	assert.Equal(t, 200, response.StatusCode)
	assert.Nil(t, err)

	dfdtesting.DeleteTestUser(ctx, testUser)
}
