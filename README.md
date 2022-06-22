# Golang Websockets

## Server
PORT=3000 go run main.go

docker build -t server .
docker run -p 3000:3000 server

## Client
PORT=3000 go run client/main.go
