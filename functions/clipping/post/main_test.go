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

func TestHandlerUpdateName(t *testing.T) {

	dfdtesting.SetTestCollectionOverride()
	dfdtesting.SetUserViewFindByUsernamePredicate()

	ctx := context.TODO()

	user := dfdtestingapi.CreateTestAuthenticatedUser("TestUser")
	clippingName := "clipping"
	updatedClippingName := "clipping_updated"

	id, err := clippingapi.Create(ctx, user, clippingName)
	assert.Nil(t, err)

	updates := make(map[string]string)
	updates["name"] = updatedClippingName

	request := &updateClippingRequest{
		ClippingID: id,
		UpdateType: "UPDATE",
		Updates:    updates,
	}

	// Arrange
	apirequest := dfdtestingapi.CreateTestAuthorizedRequest("TestUser")

	b, err := json.Marshal(request)
	assert.Nil(t, err)
	apirequest.Body = string(b)

	// Act
	response, err := Handler(*apirequest)

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, 200, response.StatusCode)
	assert.NotNil(t, response)

	clipping := &dfdmodels.Clipping{}
	err = json.Unmarshal([]byte(response.Body), clipping)
	assert.Nil(t, err)
	assert.NotNil(t, clipping)

	assert.Equal(t, *id, *clipping.ID)
	assert.Equal(t, updatedClippingName, clipping.Name)

	dfdtestingapi.DeleteUserForRequest(ctx, apirequest)
}

// func TestRequestUnmarshalsCorrectly(t *testing.T) {

// 	o := primitive.NewObjectID()
// 	body := "{\"recipeId\": \"" + o.Hex() + "\"}"

// 	// Arrange
// 	apirequest := dfdtestingapi.CreateTestAuthorizedRequest("TestUser")
// 	apirequest.Body = body

// 	// Act
// 	r, err := parseRequest(apirequest)

// 	// Assert
// 	assert.Nil(t, err)
// 	assert.NotNil(t, r)
// 	assert.Equal(t, o, *r.RecipeID)
// }
