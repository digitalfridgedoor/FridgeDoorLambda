mkdir bin

go build -o bin/image/get ./functions/image/get
go build -o bin/ingredient/get ./functions/ingredient/get
go build -o bin/ingredient/put ./functions/ingredient/put
go build -o bin/recipe/get ./functions/recipe/get
go build -o bin/recipe/p_id/delete ./functions/recipe/p_id/delete
go build -o bin/recipe/p_id/get ./functions/recipe/p_id/get
go build -o bin/recipe/post ./functions/recipe/post
go build -o bin/recipe/put ./functions/recipe/put
go build -o bin/recipe/search/get ./functions/recipe/search/get
go build -o bin/userview/get ./functions/userview/get
go build -o bin/userview/p_id/get ./functions/userview/p_id/get