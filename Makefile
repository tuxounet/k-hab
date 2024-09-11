GIT_BRANCH := $(shell git rev-parse --abbrev-ref HEAD)
RUN_ARGS := --loglevel=DEBUG --setup=./config/default.setup.yaml

test: 
	go test ./... -timeout 120s -coverpkg=./... -coverprofile=profile.coverage
	go tool cover -func profile.coverage

build:
	mkdir -p ./out
	go build -ldflags="-X 'main.version=v${GIT_BRANCH}'" -o ./out/k-hab main.go

release:
	mkdir -p ./out
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./out/k-hab-linux-amd64 ./main.go

provision:
	go run ./main.go ${RUN_ARGS} provision 
unprovision:
	go run ./main.go ${RUN_ARGS} unprovision
up:
	go run ./main.go ${RUN_ARGS} up
deploy:
	go run ./main.go ${RUN_ARGS} deploy
shell:
	go run ./main.go ${RUN_ARGS} shell
undeploy:
	go run ./main.go ${RUN_ARGS} undeploy
down:
	go run ./main.go ${RUN_ARGS} down
rm:
	go run ./main.go ${RUN_ARGS} rm
nuke:
	go run ./main.go ${RUN_ARGS} nuke