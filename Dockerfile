FROM golang:alpine as builder

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN go build -o main .

FROM scratch

COPY --from=builder /app/main /

EXPOSE 5000

CMD ["/main"]
# ENTRYPOINT [ "/main" ]