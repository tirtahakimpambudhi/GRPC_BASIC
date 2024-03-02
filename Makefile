
gen:
	protoc --proto_path=proto proto/*.proto --go_out=. --go-grpc_out=.
server:
	go run ./cmd/server/main.go -port 5600
client:
	go run ./cmd/client/main.go -addr "0.0.0.0:5600" -meth "upload"
clean:
	rm -f ./pb/*.go
testing:
	go test -cover -race ./...

create_laptop:
	for i in $$(seq 0 9); do \
    	    go run ./cmd/client/main.go -addr "0.0.0.0:5600" -meth "create"; \
    done