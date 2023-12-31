FROM golang:1.18-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build ./cmd/failed_transaction
RUN go build ./cmd/transaction

EXPOSE 8080

CMD ["./transaction"]