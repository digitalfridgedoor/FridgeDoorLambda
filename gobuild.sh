mkdir bin

go build -o bin/clipping/p_id/delete ./functions/clipping/p_id/delete
go build -o bin/clipping/p_id/get ./functions/clipping/p_id/get
go build -o bin/clipping/post ./functions/clipping/post
go build -o bin/clipping/put ./functions/clipping/put
go build -o bin/clipping/search/get ./functions/clipping/search/get
go build -o bin/image/get ./functions/image/get
go build -o bin/ingredient/get ./functions/ingredient/get
go build -o bin/ingredient/put ./functions/ingredient/put
go build -o bin/public/image/get ./functions/public/image/get
go build -o bin/public/recipe/get ./functions/public/recipe/get
go build -o bin/recipe/get ./functions/recipe/get
go build -o bin/recipe/p_id/delete ./functions/recipe/p_id/delete
go build -o bin/recipe/p_id/get ./functions/recipe/p_id/get
go build -o bin/recipe/post ./functions/recipe/post
go build -o bin/recipe/put ./functions/recipe/put
go build -o bin/userview/get ./functions/userview/get
go build -o bin/userview/tags/get ./functions/userview/tags/get