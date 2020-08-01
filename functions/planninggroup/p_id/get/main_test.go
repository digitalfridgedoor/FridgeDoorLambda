package main

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/digitalfridgedoor/fridgedoorapi/dfdtesting"
	"github.com/digitalfridgedoor/fridgedoorapi/planninggroupapi"

	"github.com/stretchr/testify/assert"
)

func TestHandler(t *testing.T) {

	// Arrange
	dfdtesting.SetTestCollectionOverride()
	dfdtesting.SetUserViewFindByUsernamePredicate()

	ctx := context.TODO()

	user, apirequest := dfdtesting.CreateTestAuthenticatedUserAndRequest("TestUser")

	groupName := "Planning group"
	id, err := planninggroupapi.Create(ctx, user, groupName)
	assert.Nil(t, err)

	pathParameters := make(map[string]string)
	pathParameters["id"] = id.Hex()
	apirequest.PathParameters = pathParameters

	// Act
	response, err := Handler(*apirequest)

	// Assert
	assert.Nil(t, err)
	planningGroup := &PlanningGroupInfo{}

	err = json.Unmarshal([]byte(response.Body), planningGroup)
	assert.Nil(t, err)
	assert.NotNil(t, planningGroup)
	assert.Equal(t, *id, *planningGroup.GroupID)
	assert.Equal(t, groupName, planningGroup.Name)
	assert.Equal(t, 1, len(planningGroup.Users))
	assert.Equal(t, "", planningGroup.Users[0])

	dfdtesting.DeleteUserForRequest(ctx, apirequest)
}
