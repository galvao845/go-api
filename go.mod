module main

go 1.17

replace getMethods => ./src/getMethods

replace dataBaseInstance => ./src/dataBaseInstance

replace postMethods => ./src/postMethods

require (
	getMethods v0.0.0-00010101000000-000000000000
	postMethods v0.0.0-00010101000000-000000000000
)

require (
	dataBaseInstance v0.0.0-00010101000000-000000000000 // indirect
	github.com/lib/pq v1.10.4 // indirect
)
