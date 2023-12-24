# Generate proto # 
https://docs.buf.build/installation
1. $ buf mod update
2. $ buf generate

# Run Application #

### use envoy proxy to translate HTTP/1 to HTTP/2 for now grpc-web is required
`envoy -c envoy.yaml`
### use go
`go run server.go --logtostderr=1`

### use docker-compose
`docker-compose up`

### use nodemon
`nodemon --exec "go run" ./server.go --signal SIGTERM`

# Run Test
go test ./service/users

# User Creds # 

`{
  "username": "admin",
  "password": "admin"
}`

`{
  "username": "user",
  "password": "user"
}`

# Kubernetes Note #
kubecl
https://kubernetes.io/docs/tasks/configure-pod-container/translate-compose-kubernetes/
1. this app use `kompose convert`
2. `kubectl apply -f server-tcp-service.yaml,mongo-service.yaml,server-service.yaml,mongo-claim0-persistentvolumeclaim.yaml,server-deployment.yaml,mongo-deployment.yaml`
