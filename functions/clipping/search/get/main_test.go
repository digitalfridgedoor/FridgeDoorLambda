package main

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/digitalfridgedoor/fridgedoorapi/clippingapi"
	"github.com/digitalfridgedoor/fridgedoorapi/dfdmodels"
	"github.com/digitalfridgedoor/fridgedoorapi/dfdtesting"

	"github.com/stretchr/testify/assert"
)

func TestHandler(t *testing.T) {

	// Arrange
	dfdtesting.SetTestCollectionOverride()
	dfdtesting.SetUserViewFindByUsernamePredicate()
	dfdtesting.SetClippingByNamePredicate()

	user, apirequest := dfdtesting.CreateTestAuthenticatedUserAndRequest("TestUser")

	apirequest.QueryStringParameters = make(map[string]string)
	apirequest.QueryStringParameters["q"] = "fi"

	clippingapi.Create(context.TODO(), user, "fi_clipping")

	// Act
	response, err := Handler(*apirequest)

	// Assert
	assert.Nil(t, err)
	clippings := []*dfdmodels.Clipping{}

	err = json.Unmarshal([]byte(response.Body), &clippings)
	assert.Nil(t, err)
	assert.NotNil(t, clippings)
	assert.Equal(t, 1, len(clippings))
}
