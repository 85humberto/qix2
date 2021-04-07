FROM golang:1.16.2-alpine as builder
ARG KIND

WORKDIR /go/src/app
COPY . ./
RUN go get -d -v ./...
RUN CGO_ENABLED=0 go build -ldflags="-s -w" -v ${KIND}/main.go

FROM scratch

WORKDIR /app
COPY --from=builder /go/src/app/main .
EXPOSE 50051
EXPOSE 50052
CMD ["./main"]
