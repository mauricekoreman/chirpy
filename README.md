This was a project on Boot.dev

Make sure you have the following installed:

Golang
Postgresql (im using v14)
goose
sqlc

in the root of the dir execute sqlc generate to generate the database queries into GO code. Run the migrations with goose [database connection] up
