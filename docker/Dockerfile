FROM golang:1.17-alpine as builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY ./ ./

RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o server /app/cmd/api/main.go

# 'scratch' image not working
FROM golang:1.17-alpine

# ENV config=docker

COPY --from=builder /app/config /config
COPY --from=builder /app/server /server

EXPOSE 5000
EXPOSE 5555
EXPOSE 7070

ENTRYPOINT [ "/server" ]
