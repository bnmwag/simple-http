FROM golang:1.22-alpine

WORKDIR /app

COPY . .

RUN go build -o /app/server .

EXPOSE 8080

CMD ["./server"]