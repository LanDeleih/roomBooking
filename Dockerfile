FROM golang:1.15.2 as builder

WORKDIR /dist

COPY . /dist

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o /roomBooking ./cmd/roomBooking/main.go

FROM ubuntu:20.04 as prod-image

RUN apt -y update && \
    apt -y install --no-recommends ca-certificates

USER nobody

COPY --chown=nobody:nogroup --from=builder /roomBooking /roomBooking

CMD ["/roomBooking"]
