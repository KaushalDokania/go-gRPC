clean:
	find ./helloworld/ -name "*.go" -type f -delete
	find ./cert/ -name "*.srl" -type f -delete
	find ./cert/ -name "*.pem" -type f -delete

gen:
	export PATH="$PATH:$(go env GOPATH)/bin"
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative helloworld/helloworld.proto

cert:
	cd cert; ./gen-certs.sh; cd ..

.PHONY: gen server client cert 
