build:
	go build -o bin/main
run: build
	bin/main
test: 
	go test -v ./...
cover:
	go test -coverprofile cover.out
	go tool cover -html=cover.out -o cover.html
	open cover.html

migrateup:
	migrate -path db/migration -database "postgresql://hello:Hello123@@localhost:5432/hello?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://hello:Hello123@@localhost:5432/hello?sslmode=disable" -verbose down

docker-build:
	docker build -t arts-nthu-backend .

docker-run:
	docker run arts-nthu-backend
	
clean:
	rm -rf bin/*
	rm -f cover.html cover.out
