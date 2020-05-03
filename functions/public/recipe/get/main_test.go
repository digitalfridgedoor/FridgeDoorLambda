package main

import (
	"encoding/json"
	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/digitalfridgedoor/fridgedoorapi/fridgedoorgatewaytesting"
	"github.com/digitalfridgedoor/fridgedoorapi/recipeapi"
	"github.com/digitalfridgedoor/fridgedoordatabase/dfdmodels"
	"github.com/digitalfridgedoor/fridgedoordatabase/dfdtesting"

	"github.com/stretchr/testify/assert"
)

func TestHandler(t *testing.T) {

	// Arrange
	dfdtesting.SetTestCollectionOverride()
	dfdtesting.SetUserViewFindPredicate(func(uv *dfdmodels.UserView, m primitive.M) bool {
		return m["username"] == uv.Username
	})

	apirequest := fridgedoorgatewaytesting.CreateTestAuthorizedRequest("TestUser")

	// Act
	response, err := Handler(*apirequest)

	// Assert
	assert.Nil(t, err)
	recipeCollection := []*recipeapi.Recipe{}

	err = json.Unmarshal([]byte(response.Body), &recipeCollection)
	assert.Nil(t, err)
	assert.NotNil(t, recipeCollection)
	assert.Equal(t, 0, len(recipeCollection))

	fridgedoorgatewaytesting.DeleteUserForRequest(apirequest)
}
