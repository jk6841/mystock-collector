FROM golang:1.20 as builder

WORKDIR /builder

COPY . .
RUN go mod download

RUN CGO_ENABLED=0 go build -o app

FROM gcr.io/distroless/static-debian11

COPY --from=builder /builder/app /

CMD ["/app"]