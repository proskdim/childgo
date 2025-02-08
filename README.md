### Childgo

start app
```go
 go run main.go
 # or
 go build .
 ./childgo
```

run tests

```go
go test -v
```

build image childgo
```go
docker build -t childgo .
```

run container with env options
```go
docker run -p 3005:8087 -d --env-file ./.env childgo
```

### Endpoints

```go
GET /api/v1/
POST /api/v1/signin
POST /api/v1/signup
GET /api/v1/profile # require Authorization

CRUD 
/api/v1/child  # require Authorization
```
