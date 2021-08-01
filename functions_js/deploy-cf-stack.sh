#!/bin/bash

set -e

templateFile='./functions_js/template.yml'

aws cloudformation deploy \
    --stack-name js-lambda-pipeline \
    --template-file $templateFile \
    --capabilities CAPABILITY_IAM \
    --no-fail-on-empty-changeset

resource=$(aws cloudformation describe-stack-resource \
  --stack-name js-lambda-pipeline \
  --logical-resource-id ParseRecipeLambda)

physicalResource=$(echo $resource | jq -r '.StackResourceDetail.PhysicalResourceId')

echo "::set-output name=lambda_function_name::$physicalResource"

