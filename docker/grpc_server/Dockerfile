FROM golang:latest

WORKDIR /app

COPY . .

RUN go mod download

WORKDIR /app/cmd/server/

RUN go build -o /grpc-server

EXPOSE 8080

CMD [ "/grpc-server" ]