package main

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/digitalfridgedoor/fridgedoorapi/clippingapi"
	"github.com/digitalfridgedoor/fridgedoorapi/dfdmodels"
	"github.com/digitalfridgedoor/fridgedoorapi/dfdtesting"
	"github.com/digitalfridgedoor/fridgedoorapi/dfdtestingapi"

	"github.com/stretchr/testify/assert"
)

func TestHandler(t *testing.T) {

	// Arrange
	dfdtesting.SetTestCollectionOverride()

	ctx := context.TODO()

	user, apirequest := dfdtestingapi.CreateTestAuthenticatedUserAndRequest("TestUser")

	clippingName := "clipping"
	id, err := clippingapi.Create(ctx, user, clippingName)
	assert.Nil(t, err)

	pathParameters := make(map[string]string)
	pathParameters["id"] = id.Hex()
	apirequest.PathParameters = pathParameters

	// Act
	response, err := Handler(*apirequest)

	// Assert
	assert.Nil(t, err)
	clipping := &dfdmodels.Clipping{}

	err = json.Unmarshal([]byte(response.Body), clipping)
	assert.Nil(t, err)
	assert.NotNil(t, clipping)
	assert.Equal(t, *id, *clipping.ID)
	assert.Equal(t, clippingName, clipping.Name)

	dfdtestingapi.DeleteUserForRequest(ctx, apirequest)
}
