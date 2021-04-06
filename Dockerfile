FROM golang:1.16.2-alpine as builder
ARG KIND

# RUN apk add git
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

#  docker build --build-arg=KIND=server -t 85humberto/qix2-server .