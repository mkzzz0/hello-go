FROM golang:1.23-alpine
WORKDIR /app
COPY ./hello.go .
COPY ./go.mod .
RUN go build -o main .
CMD ["./main"]
EXPOSE 7777
