FROM golang:latest

WORKDIR /app

COPY . .

RUN go mod download

WORKDIR /app/cmd/grpc_gateway/

RUN go build -o /grpc-gateway

EXPOSE 8081

CMD [ "/grpc-gateway" ]