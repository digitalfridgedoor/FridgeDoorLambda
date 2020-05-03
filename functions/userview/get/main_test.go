package main

import (
	"encoding/json"
	"testing"

	"github.com/digitalfridgedoor/fridgedoordatabase/dfdmodels"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/digitalfridgedoor/fridgedoordatabase/dfdtesting"

	"github.com/digitalfridgedoor/fridgedoorapi/linkeduserapi"

	"github.com/digitalfridgedoor/fridgedoorapi/fridgedoorgatewaytesting"

	"github.com/stretchr/testify/assert"
)

func TestHandler(t *testing.T) {

	// Arrange
	dfdtesting.SetTestCollectionOverride()
	dfdtesting.SetUserViewFindPredicate(func(uv *dfdmodels.UserView, m bson.M) bool {
		return true
	})

	fridgedoorgatewaytesting.CreateTestAuthenticatedUser("linked")

	apirequest := fridgedoorgatewaytesting.CreateTestAuthorizedRequest("TestUser")

	// Act
	response, err := Handler(*apirequest)

	// Assert
	assert.Nil(t, err)
	var linkedusers []*linkeduserapi.LinkedUser

	err = json.Unmarshal([]byte(response.Body), &linkedusers)
	assert.Nil(t, err)
	assert.NotNil(t, linkedusers)
	assert.Greater(t, len(linkedusers), 0)

	fridgedoorgatewaytesting.DeleteUserForRequest(apirequest)
}
