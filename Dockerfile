FROM golang:1-alpine as builder
WORKDIR /go/src
COPY go.mod .
COPY go.sum .
RUN GO111MODULE=on GOPROXY=https://proxy.golang.org go mod download
COPY app.go .
RUN CGO_ENABLED=0 go build -a --trimpath --installsuffix cgo --ldflags="-s" -o ip

FROM scratch
COPY --from=builder /go/src/ip .
ENTRYPOINT ["/ip"]
EXPOSE 80
