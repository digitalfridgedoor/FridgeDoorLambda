package main

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/digitalfridgedoor/fridgedoorapi/dfdmodels"
	"github.com/digitalfridgedoor/fridgedoorapi/dfdtesting"
	"github.com/digitalfridgedoor/fridgedoorapi/linkeduserapi"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

func TestHandler(t *testing.T) {

	// Arrange
	dfdtesting.SetTestCollectionOverride()
	dfdtesting.SetUserViewFindPredicate(func(uv *dfdmodels.UserView, m bson.M) bool {
		return true
	})

	ctx := context.TODO()

	dfdtesting.CreateTestAuthenticatedUser("linked")

	apirequest := dfdtesting.CreateTestAuthorizedRequest("TestUser")

	// Act
	response, err := Handler(*apirequest)

	// Assert
	assert.Nil(t, err)
	var linkedusers []*linkeduserapi.LinkedUser

	err = json.Unmarshal([]byte(response.Body), &linkedusers)
	assert.Nil(t, err)
	assert.NotNil(t, linkedusers)
	assert.Greater(t, len(linkedusers), 0)

	dfdtesting.DeleteUserForRequest(ctx, apirequest)
}
