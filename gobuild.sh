mkdir bin

go build -o bin/ingredient/get ./functions/ingredient/get
go build -o bin/recipe/get ./functions/recipe/get
go build -o bin/recipe/p_id/get ./functions/recipe/p_id/get
go build -o bin/recipe/post ./functions/recipe/post
go build -o bin/recipe/put ./functions/recipe/put