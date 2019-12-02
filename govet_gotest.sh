go vet ./functions/ingredients/get
go vet ./functions/recipe/get
go vet ./functions/recipe/post
go vet ./functions/recipe/put
go vet ./functions/recipe/{id}/get

go test ./functions/ingredients/get
go test ./functions/recipe/get
go test ./functions/recipe/post
go test ./functions/recipe/put
go test ./functions/recipe/{id}/get