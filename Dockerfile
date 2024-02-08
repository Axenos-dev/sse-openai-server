FROM golang:latest

WORKDIR /go/src/app

COPY . .

RUN go mod download
RUN go build -o sse-openai-server .

EXPOSE 8080
EXPOSE 8080

CMD ["./sse-openai-server"]