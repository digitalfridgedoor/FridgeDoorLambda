mkdir bin

go build -o bin/ingredients/get ./functions/ingredients/get
go build -o bin/recipe/get ./functions/recipe/get
go build -o bin/recipe/post ./functions/recipe/post
go build -o bin/recipe/put ./functions/recipe/put
go build -o bin/recipe/{id}/get ./functions/recipe/{id}/get