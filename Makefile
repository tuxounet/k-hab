test: 
	go test -v -coverpkg=./... -coverprofile=profile.coverage ./...

build:
	mkdir -p ./out
	go build -o ./out/k-hab main.go


release:
	mkdir -p ./out
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./out/k-hab-linux-amd64 ./main.go

provision:
	go run ./main.go provision 

unprovision:
	go run ./main.go unprovision
up:
	go run ./main.go up
shell:
	go run ./main.go shell

down:
	go run ./main.go down
rm:
	go run ./main.go rm
nuke:
	go run ./main.go nuke