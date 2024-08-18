test: 
	go test ./... --cover

build:
	mkdir -p ./out
	go build -o ./out/k-hab main.go


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