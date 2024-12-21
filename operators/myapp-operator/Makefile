build:
    go build -o bin/manager main.go

run:
    go run ./main.go

deploy:
    kubectl apply -f config/manifests/

test:
    go test ./...