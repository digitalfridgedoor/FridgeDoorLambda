mkdir bin

go build -o bin/createrecipe ./functions/createrecipe
go build -o bin/searchingredients ./functions/searchingredients
go build -o bin/updaterecipe ./functions/updaterecipe
go build -o bin/viewrecipe ./functions/viewrecipe
go build -o bin/viewrecipes ./functions/viewrecipes