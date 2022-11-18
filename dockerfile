FROM golang:alpine

WORKDIR /build

ADD . .

RUN go mod download
RUN GOOS=linux
RUN go build -a -installsuffix 'api' -o app cmd/main.go

# FROM alpine:latest
# COPY --from=builder /build/app .
EXPOSE 8080 
CMD ["./app"]