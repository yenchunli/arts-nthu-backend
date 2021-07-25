# build stage
FROM golang:1.15-alpine AS build-env
ADD . /server
RUN cd /server && go build -o bin/main cmd/server/main.go

# final stage
FROM alpine
WORKDIR /server
COPY --from=build-env /server /server
EXPOSE 8080
ENTRYPOINT bin/main