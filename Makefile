build:
	go build -o bin/main
run: build
	bin/main
test: 
	go test -v
cover:
	go test -coverprofile cover.out
	go tool cover -html=cover.out -o cover.html
	open cover.html
clean:
	rm -rf bin/*
	rm -f cover.html cover.out
