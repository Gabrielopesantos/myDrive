FROM golang:1.16-alpine as build

ENV config=docker

WORKDIR /app

COPY ./ /app

RUN go mod download

FROM golang:1.16-alpine as runner

COPY --from=builder ./app ./app

RUN go get github.com/githubnemo/CompileDaemon

WORKDIR /app
ENV config=docker

EXPOSE 5000
EXPOSE 5555
EXPOSE 7070

ENTRYPOINT CompileDeamon --build="go build cmd/api/main.go" --comand=./main
