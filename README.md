# go-gRPC

```
make clean
make gen (optional)
```
## Commands to Start the server and client

### 1. Insecure
```
go run cmd/server/insecure/main.go
go run cmd/client/insecure/main.go
```

### 2. Server-side TLS
```
# Generate the certificates
make cert

go run cmd/server/tls/main.go
go run cmd/client/tls/main.go
```

###  3. Mutual TLS
```
# Generate the certificates
make cert

go run cmd/server/mutual-tls/main.go
go run cmd/client/mutual-tls/main.go
```