use go-migrate
https://github.com/golang-migrate/migrate/tree/master/cmd/migrate

`migrate -path migration -database "postgresql://postgres@localhost:5432/travel?sslmode=disable" -verbose up`

Generate proto
1. $ buf mod update
2. $ buf generate

Run Application
go run server.go --logtostderr=1

Run Test
go test ./service/movie

Log is at utils.go