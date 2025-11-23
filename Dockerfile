FROM golang:alpine as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main .

FROM scratch

COPY --from=builder /app/main /

EXPOSE 3000

CMD ["/main"]