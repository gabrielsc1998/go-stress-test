FROM golang:1.21.1 as build

WORKDIR /go/src
COPY . .

RUN cd /go/src && GOOS=linux CGO_ENABLED=0 go build -o stress-test cmd/main.go

RUN useradd -u 10001 appuser

FROM scratch

COPY --from=0 /etc/passwd /etc/passwd
USER appuser

COPY --from=build --chown=appuser:appuser /go/src/stress-test ./

ENTRYPOINT ["./stress-test", "test"]