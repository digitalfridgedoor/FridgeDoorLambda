package main

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/digitalfridgedoor/fridgedoorapi/planninggroupapi"

	"github.com/digitalfridgedoor/fridgedoorapi/dfdmodels"
	"github.com/digitalfridgedoor/fridgedoorapi/dfdtesting"

	"github.com/stretchr/testify/assert"
)

func TestHandlerUpdateName(t *testing.T) {

	// Arrange
	dfdtesting.SetTestCollectionOverride()
	dfdtesting.SetUserViewFindByUsernamePredicate()
	dfdtesting.SetPlanningGroupFindByUser()

	ctx := context.TODO()

	user, apirequest := dfdtesting.CreateTestAuthenticatedUserAndRequest("TestUser")

	groupName := "Group"
	_, err := planninggroupapi.Create(ctx, user, groupName)
	assert.Nil(t, err)

	// Act
	response, err := Handler(*apirequest)

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, 200, response.StatusCode)
	assert.NotNil(t, response)

	groups := &[]*dfdmodels.PlanningGroup{}
	err = json.Unmarshal([]byte(response.Body), groups)
	assert.Nil(t, err)
	assert.NotNil(t, groups)

	assert.Equal(t, 1, len(*groups))
	assert.Equal(t, groupName, (*groups)[0].Name)

	dfdtesting.DeleteUserForRequest(ctx, apirequest)
}
