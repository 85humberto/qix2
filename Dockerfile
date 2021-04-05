FROM golang:1.16.2-alpine as builder
ARG KIND=server

# RUN apk add git
WORKDIR /go/src/app
COPY . ./
RUN go get -d -v ./...
RUN CGO_ENABLED=0 go build -ldflags="-s -w" -v ${KIND}/main.go

FROM scratch
ARG KIND

WORKDIR /app
COPY --from=builder /go/src/app/${KIND} .
EXPOSE 5000
CMD ["./${KIND}"]

#  docker build --build-arg=KIND=server -t 85humberto/qix2-server .