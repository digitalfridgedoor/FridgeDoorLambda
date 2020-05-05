package main

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/digitalfridgedoor/fridgedoorapi/fridgedoorgatewaytesting"
	"github.com/digitalfridgedoor/fridgedoordatabase/dfdmodels"
	"github.com/digitalfridgedoor/fridgedoordatabase/dfdtesting"

	"github.com/stretchr/testify/assert"
)

func TestHandler(t *testing.T) {

	// Arrange
	dfdtesting.SetTestCollectionOverride()
	dfdtesting.SetPlanFindPredicate(dfdtesting.FindPlanByMonthAndYearTestPredicate)
	dfdtesting.SetUserViewFindByUsernamePredicate()

	ctx := context.TODO()

	user, apirequest := fridgedoorgatewaytesting.CreateTestAuthenticatedUserAndRequest("TestUser")

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
	assert.Equal(t, user.ViewID, plan.UserID)

	fridgedoorgatewaytesting.DeleteUserForRequest(ctx, apirequest)
}
