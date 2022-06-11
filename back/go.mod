module github.com/dchf12/todo

go 1.18

require (
	github.com/dchf12/todo/table v0.0.0-00010101000000-000000000000
	github.com/mattn/go-sqlite3 v1.14.13
)

replace github.com/dchf12/todo/table => ./table
