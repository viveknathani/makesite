build:
	go build -o ./bin/
	sudo cp ./bin/makesite ~/bin/makesite

test:
	go test -v ./...
	