FROM golang:alpine as builder

WORKDIR /build

ADD . .

RUN go mod download
ARG GOOS=linux
ARG CGO_ENABLED=0
RUN go build -a -installsuffix 'api' -o app cmd/main.go

FROM alpine:latest
COPY --from=builder /build/app .
EXPOSE 8080 8080 
ENTRYPOINT ["./app"]
