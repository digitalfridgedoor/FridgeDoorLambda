package main

import (
	"encoding/json"
	"testing"

	"github.com/digitalfridgedoor/fridgedoordatabase/dfdmodels"

	"github.com/digitalfridgedoor/fridgedoorapi"
	"github.com/digitalfridgedoor/fridgedoorapi/fridgedoorgatewaytesting"

	"github.com/stretchr/testify/assert"
)

func TestHandler(t *testing.T) {

	// Arrange
	apirequest := fridgedoorgatewaytesting.CreateTestAuthorizedRequest("7a401f20-86ca-442f-acf3-20d396c06d33")

	apirequest.QueryStringParameters = make(map[string]string)
	apirequest.QueryStringParameters["q"] = "fi"

	// Act
	fridgedoorapi.ConnectOrSkip(t)

	response, err := Handler(*apirequest)

	// Assert
	assert.Nil(t, err)
	recipes := []*dfdmodels.Recipe{}

	err = json.Unmarshal([]byte(response.Body), &recipes)
	assert.Nil(t, err)
	assert.NotNil(t, recipes)
	assert.Equal(t, 1, len(recipes))
}
