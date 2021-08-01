cd ./functions_js/code

echo "Updating "$LAMBDA_FUNCTION_NAME

zip -r lambda-function-code.zip .

aws lambda update-function-code \
    --function-name $LAMBDA_FUNCTION_NAME \
    --zip-file fileb://$(pwd)/lambda-function-code.zip
