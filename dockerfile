FROM golang:alpine as builder

WORKDIR /api

ADD . .

RUN go mod download
ARG GOOS=linux
ARG CGO_ENABLED=0
RUN go build -a -installsuffix 'api' -o app cmd/main.go

FROM alpine:latest
COPY --from=builder /api/app .
EXPOSE 8080 8080 
ENTRYPOINT ["./app"]
