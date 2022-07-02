Generate proto
https://docs.buf.build/installation
1. $ buf mod update
2. $ buf generate

Run Application
go run server.go --logtostderr=1

Run Test
go test ./service/users