package main

import (
	"context"
	"encoding/json"
	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/digitalfridgedoor/fridgedoorapi/dfdmodels"
	"github.com/digitalfridgedoor/fridgedoorapi/dfdtesting"

	"github.com/stretchr/testify/assert"
)

func TestHandler(t *testing.T) {

	// Arrange
	dfdtesting.SetTestCollectionOverride()
	dfdtesting.SetPlanFindPredicate(dfdtesting.FindPlanByMonthAndYearTestPredicate)
	dfdtesting.SetUserViewFindByUsernamePredicate()

	ctx := context.TODO()

	user, apirequest := dfdtesting.CreateTestAuthenticatedUserAndRequest("TestUser")

	params := make(map[string]string)
	params["month"] = "1"
	params["year"] = "2020"
	apirequest.QueryStringParameters = params

	// Act
	response, err := Handler(*apirequest)

	// Assert
	assert.Nil(t, err)
	plan := &dfdmodels.Plan{}

	err = json.Unmarshal([]byte(response.Body), &plan)
	assert.Nil(t, err)
	assert.NotNil(t, plan)
	assert.Equal(t, 1, plan.Month)
	assert.Equal(t, 2020, plan.Year)
	assert.Equal(t, user.ViewID, *plan.UserID)
	assert.Equal(t, (*primitive.ObjectID)(nil), plan.PlanningGroupID)

	dfdtesting.DeleteUserForRequest(ctx, apirequest)
}

func TestHandlerForGroup(t *testing.T) {

	// Arrange
	dfdtesting.SetTestCollectionOverride()
	dfdtesting.SetPlanFindPredicate(dfdtesting.FindPlanByMonthAndYearForGroupTestPredicate)
	dfdtesting.SetUserViewFindByUsernamePredicate()

	ctx := context.TODO()

	_, apirequest := dfdtesting.CreateTestAuthenticatedUserAndRequest("TestUser")

	groupID := primitive.NewObjectID()

	params := make(map[string]string)
	params["month"] = "1"
	params["year"] = "2020"
	params["planningGroupId"] = groupID.Hex()
	apirequest.QueryStringParameters = params

	// Act
	response, err := Handler(*apirequest)

	// Assert
	assert.Nil(t, err)
	plan := &dfdmodels.Plan{}

	err = json.Unmarshal([]byte(response.Body), &plan)
	assert.Nil(t, err)
	assert.NotNil(t, plan)
	assert.Equal(t, 1, plan.Month)
	assert.Equal(t, 2020, plan.Year)
	assert.Equal(t, (*primitive.ObjectID)(nil), plan.UserID)
	assert.Equal(t, groupID, *plan.PlanningGroupID)

	dfdtesting.DeleteUserForRequest(ctx, apirequest)
}
